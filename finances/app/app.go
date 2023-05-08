package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log"
	"main/finances/entrypoints"
	"main/finances/entrypoints/middleware"
	"main/finances/repo"
	"main/finances/services/group"
	"main/finances/services/notify"
	"main/finances/services/privacy"
	"main/finances/services/spend"
	"main/finances/services/statistic"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context, pool *pgxpool.Pool) (error, App) {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Print(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	groupRepo := repo.NewPostgresGroupRepository(pool)
	finRepo := repo.NewPostgresFinanceRepository(pool)
	notifier := notify.NewNotifier(rdb)

	groupAccessChecker := privacy.NewGroupAccessChecker(groupRepo)

	finEntry := entrypoints.FinanceEntryPoint{
		CreateSpendService:    spend.NewCreateSpendsService(finRepo, groupAccessChecker),
		GetSpendsService:      spend.NewGetSpendsService(finRepo),
		DeleteSpendService:    spend.NewDeleteSpendsService(finRepo, groupAccessChecker),
		UpdateSpendService:    spend.NewUpdateSpendsService(finRepo, groupAccessChecker),
		GetGroupSpendsService: spend.NewGroupSpendsService(finRepo, groupAccessChecker),
		Ctx:                   ctx,
	}

	groupEntry := entrypoints.GroupEntryPoint{
		AddMemberService:        group.NewAddGroupMemberService(groupRepo, groupAccessChecker, notifier),
		DeleteMemberService:     group.NewDeleteGroupMemberSrvice(groupRepo, groupAccessChecker, notifier),
		CreateSpendGroupServe:   group.NewCreateSpendGroupService(groupRepo),
		DeleteSpendGroupService: group.NewDeleteSpendGroupService(groupRepo, groupAccessChecker, notifier),
		GetUserGroupsService:    group.NewGetGroupService(groupRepo),
	}

	statEntryPoint := entrypoints.StatisticEntryPoint{
		StatLoader: statistic.NewStatLoader(finRepo),
	}

	commonEntry := http.NewServeMux()

	commonEntry.Handle("/", http.StripPrefix("/spending", finEntry.FinanceEntrypoint()))
	commonEntry.Handle("/group/member/add", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.AddNewMember)))
	commonEntry.Handle("/group/member/delete", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.DeleteMember)))
	commonEntry.Handle("/group/create", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.CreateNewSpendGroup)))
	commonEntry.Handle("/group/delete", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.DeleteSpendGroup)))
	commonEntry.Handle("/group/get", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.GetUserGroups)))

	commonEntry.Handle("/stats/user", middleware.AuthMiddleware(http.HandlerFunc(statEntryPoint.LoadStatsForUser)))

	return nil, App{AppMux: commonEntry}
}
