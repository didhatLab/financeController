package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/finances/entrypoints"
	"main/finances/entrypoints/middleware"
	"main/finances/repo"
	"main/finances/services/group"
	"main/finances/services/privacy"
	"main/finances/services/spend"
	"net/http"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context, pool *pgxpool.Pool) (error, App) {

	groupRepo := repo.NewPostgresGroupRepository(pool)
	finRepo := repo.NewPostgresFinanceRepository(pool)

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
		AddMemberService:    group.NewAddGroupMemberService(groupRepo, groupAccessChecker),
		DeleteMemberService: group.NewDeleteGroupMemberSrvice(groupRepo, groupAccessChecker),
	}

	commonEntry := http.NewServeMux()

	commonEntry.Handle("/", http.StripPrefix("/spending", finEntry.FinanceEntrypoint()))
	commonEntry.Handle("group/member/add", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.AddNewMember)))
	commonEntry.Handle("group/member/delete", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.DeleteMember)))
	commonEntry.Handle("group/create", middleware.AuthMiddleware(http.HandlerFunc(groupEntry.CreateNewSpendGroup)))

	return nil, App{AppMux: commonEntry}
}

func TestMux() *http.ServeMux {
	ff := http.NewServeMux()
	ff.HandleFunc("/test", check)
	ff.HandleFunc("/rrr", check)

	return ff
}

func check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	return
}
