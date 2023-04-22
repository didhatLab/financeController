package middleware

import (
	"context"
	"main/finances/models/user"
)

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user user.User) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (user.User, bool) {
	_user, ok := ctx.Value(key).(user.User)
	return _user, ok
}
