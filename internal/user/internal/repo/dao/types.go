package dao

import (
	"context"
	"time"
)

type UserDao interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByUUID(ctx context.Context, uuid string) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)

	UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error
	UpdatePassword(ctx context.Context, userID int64, passwordHash string) error
	UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error
	UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error
	UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error
	UpdateStatus(ctx context.Context, userID int64, status string) error
	IncrementReportCount(ctx context.Context, userID int64) error
	UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error
	ListUsers(ctx context.Context, filter *UserFilter) ([]*User, int64, error)
}
