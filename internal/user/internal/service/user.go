package service

import (
	"MeetLens/internal/user/internal/domain"
	"MeetLens/internal/user/internal/repo"
	"MeetLens/pkg/encrypt"
	"MeetLens/pkg/snowflake"
	"MeetLens/pkg/validate"
	"context"
	"fmt"
	"strconv"
	"time"
)

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(ctx context.Context, phone, email, password string) (*domain.User, error) {
	// 检查是否已注册
	exist, _ := u.userRepo.GetByPhone(ctx, phone)
	if exist != nil {
		return nil, fmt.Errorf("phone already registered")
	}

	hash, err := encrypt.HashPassword(password)
	if err != nil {
		return nil, err
	}

	usr := &domain.User{
		UUID:         strconv.FormatInt(snowflake.GenerateID(), 10),
		Phone:        phone,
		Email:        email,
		PasswordHash: hash,
		Role:         "user",
		Status:       "active",
	}

	if err := u.userRepo.Create(ctx, usr); err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userService) Login(ctx context.Context, identifier, password string) (*domain.User, error) {
	var usr *domain.User
	var err error

	// 手机或邮箱登录
	if validate.IsEmail(identifier) {
		usr, err = u.userRepo.GetByEmail(ctx, identifier)
	} else {
		usr, err = u.userRepo.GetByPhone(ctx, identifier)
	}

	if err != nil || usr == nil {
		return nil, fmt.Errorf("user not found")
	}

	// 密码校验（示例）
	if !encrypt.CheckPassword(password, usr.PasswordHash) {
		return nil, fmt.Errorf("invalid password")
	}

	// 更新最后登录时间
	_ = u.userRepo.UpdateLastLogin(ctx, usr.ID, time.Now())

	return usr, nil
}

func (u *userService) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userService) GetByUUID(ctx context.Context, uuid string) (*domain.User, error) {
	return u.userRepo.GetByUUID(ctx, uuid)
}

func (u *userService) UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error {
	return u.userRepo.UpdateProfile(ctx, userID, nickname, avatarURL, bio)
}

func (u *userService) UpdatePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	// 先查用户
	usr, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if usr == nil {
		return fmt.Errorf("user not found")
	}

	// 验证旧密码
	if !encrypt.CheckPassword(oldPassword, usr.PasswordHash) {
		return fmt.Errorf("old password does not match")
	}

	// hash 新密码
	newHash, err := encrypt.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新
	return u.userRepo.UpdatePassword(ctx, userID, newHash)
}

func (u *userService) UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error {
	return u.userRepo.UpdateRealnameVerified(ctx, userID, verified)
}

func (u *userService) UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error {
	return u.userRepo.UpdatePhotographerVerified(ctx, userID, verified)
}

func (u *userService) UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error {
	return u.userRepo.UpdateBlacklistStatus(ctx, userID, blacklisted)
}

func (u *userService) UpdateStatus(ctx context.Context, userID int64, status string) error {
	// 可选：你可以对 status 做校验，只允许 active/banned
	if status != "active" && status != "banned" {
		return fmt.Errorf("invalid status")
	}
	return u.userRepo.UpdateStatus(ctx, userID, status)
}

func (u *userService) IncrementReportCount(ctx context.Context, userID int64) error {
	return u.userRepo.IncrementReportCount(ctx, userID)
}

func (u *userService) UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error {
	return u.userRepo.UpdateLastLogin(ctx, userID, loginAt)
}

func (u *userService) ListUsers(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int64, error) {
	return u.userRepo.ListUsers(ctx, filter)
}
