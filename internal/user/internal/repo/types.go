package repo

import (
	"MeetLens/internal/user/internal/domain"
	"context"
	"time"
)

type UserRepo interface {
	// Create 创建新用户
	Create(ctx context.Context, user *domain.User) error

	// GetByID 根据ID获取
	GetByID(ctx context.Context, id int64) (*domain.User, error)

	// GetByUUID 根据UUID获取
	GetByUUID(ctx context.Context, uuid string) (*domain.User, error)

	// GetByPhone 根据手机号或邮箱获取（用于登录）
	GetByPhone(ctx context.Context, phone string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)

	// UpdateProfile 更新用户基础资料
	UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error

	// UpdatePassword 更新密码
	UpdatePassword(ctx context.Context, userID int64, passwordHash string) error

	// UpdateRealnameVerified 更新实名状态
	UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error

	// UpdatePhotographerVerified 更新摄影师认证状态
	UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error

	// UpdateBlacklistStatus 更新黑名单状态
	UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error

	// UpdateStatus 更新账号状态
	UpdateStatus(ctx context.Context, userID int64, status string) error

	// IncrementReportCount 举报次数 +1
	IncrementReportCount(ctx context.Context, userID int64) error

	// UpdateLastLogin 更新最后登录时间
	UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error

	// ListUsers 分页查询
	ListUsers(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int64, error)
}
