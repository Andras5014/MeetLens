package service

import (
	"MeetLens/internal/user/internal/domain"
	"context"
	"time"
)

type UserService interface {
	// Register 注册新用户
	Register(ctx context.Context, phone, email, password string) (*domain.User, error)

	// Login 登录（根据手机号/邮箱）
	Login(ctx context.Context, identifier, password string) (*domain.User, error)

	// GetByID 根据ID获取用户
	GetByID(ctx context.Context, id int64) (*domain.User, error)

	// GetByUUID 根据UUID获取用户
	GetByUUID(ctx context.Context, uuid string) (*domain.User, error)

	// UpdateProfile 修改用户基础资料（昵称/头像/简介）
	UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error

	// UpdatePassword 修改用户密码
	UpdatePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error

	// UpdateRealnameVerified 更新实名状态
	UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error

	// UpdatePhotographerVerified 更新摄影师认证状态
	UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error

	// UpdateBlacklistStatus 拉黑或解封
	UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error

	// UpdateStatus 更新用户状态（封禁/激活）
	UpdateStatus(ctx context.Context, userID int64, status string) error

	// IncrementReportCount 增加举报次数
	IncrementReportCount(ctx context.Context, userID int64) error

	// UpdateLastLogin 记录登录行为
	UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error

	// ListUsers 分页搜索（后台用）
	ListUsers(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int64, error)
}
