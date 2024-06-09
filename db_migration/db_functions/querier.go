// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByPhone(ctx context.Context, phonenumber string) (User, error)
}

var _ Querier = (*Queries)(nil)