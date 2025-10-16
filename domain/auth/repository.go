package auth

import "context"

type Repository interface {
	CreateAuth(ctx context.Context, auth *Auth) error
	GetAuthByUserID(ctx context.Context, userID string) (*Auth, error)
	UpdateAuth(ctx context.Context, auth *Auth) error
	DeleteAuth(ctx context.Context, userID string) error
}
