package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/auth/models"
)

type AuthRepositoryImp struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) AuthRepositoryImp {
	return AuthRepositoryImp{
		pool: pool,
	}
}

func (ar AuthRepositoryImp) GetUserDataByUserNameAndHash(ctx context.Context, username string, passHash string) (models.AuthUser, error) {
	var user models.AuthUser

	rows, err := ar.pool.Query(ctx, "SELECT username, pass_hash, user_id FROM auth WHERE username=$1 AND pass_hash=$2",
		username, passHash)

	if err != nil {
		log.Print(err)
		return user, err
	}

	if !rows.Next() {
		return user, errors.New("not such user")
	}

	err = rows.Scan(&user.Username, &user.PasswordHash, &user.UserId)

	if err != nil {
		log.Print(err)
		return user, err
	}

	return user, nil
}

func (ar AuthRepositoryImp) SaveNewUser(ctx context.Context, username string, passwordHash string) error {
	_, err := ar.pool.Exec(ctx, "INSERT INTO auth(username, pass_hash) values ($1, $2)", username, passwordHash)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

type AuthRepository interface {
	GetUserDataByUserNameAndHash(ctx context.Context, username string, passHash string) (models.AuthUser, error)
	SaveNewUser(ctx context.Context, username string, passwordHash string) error
}
