package repo

import (
	"MeetLens/internal/user/internal/domain"
	"MeetLens/internal/user/internal/repo/dao"
	"context"
	"time"
)

type userRepo struct {
	userDao dao.UserDao
}

func NewUserRepo(userDao dao.UserDao) UserRepo {
	return &userRepo{userDao: userDao}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	return r.userDao.Create(ctx, mapDomainToDAO(user))
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	u, err := r.userDao.GetByID(ctx, id)
	if u == nil || err != nil {
		return nil, err
	}
	return mapDAOToDomain(u), nil
}

func (r *userRepo) GetByUUID(ctx context.Context, uuid string) (*domain.User, error) {
	u, err := r.userDao.GetByUUID(ctx, uuid)
	if u == nil || err != nil {
		return nil, err
	}
	return mapDAOToDomain(u), nil
}

func (r *userRepo) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	u, err := r.userDao.GetByPhone(ctx, phone)
	if u == nil || err != nil {
		return nil, err
	}
	return mapDAOToDomain(u), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.userDao.GetByEmail(ctx, email)
	if u == nil || err != nil {
		return nil, err
	}
	return mapDAOToDomain(u), nil
}

func (r *userRepo) UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error {
	return r.userDao.UpdateProfile(ctx, userID, nickname, avatarURL, bio)
}

func (r *userRepo) UpdatePassword(ctx context.Context, userID int64, passwordHash string) error {
	return r.userDao.UpdatePassword(ctx, userID, passwordHash)
}

func (r *userRepo) UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error {
	return r.userDao.UpdateRealnameVerified(ctx, userID, verified)
}

func (r *userRepo) UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error {
	return r.userDao.UpdatePhotographerVerified(ctx, userID, verified)
}

func (r *userRepo) UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error {
	return r.userDao.UpdateBlacklistStatus(ctx, userID, blacklisted)
}

func (r *userRepo) UpdateStatus(ctx context.Context, userID int64, status string) error {
	return r.userDao.UpdateStatus(ctx, userID, status)
}

func (r *userRepo) IncrementReportCount(ctx context.Context, userID int64) error {
	return r.userDao.IncrementReportCount(ctx, userID)
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error {
	return r.userDao.UpdateLastLogin(ctx, userID, loginAt)
}

func (r *userRepo) ListUsers(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int64, error) {
	daoUsers, total, err := r.userDao.ListUsers(ctx, &dao.UserFilter{
		Keyword:  filter.Keyword,
		Role:     filter.Role,
		Status:   filter.Status,
		Page:     int64(filter.Page),
		PageSize: int64(filter.PageSize),
	})
	if err != nil {
		return nil, 0, err
	}
	var result []*domain.User
	for _, u := range daoUsers {
		result = append(result, mapDAOToDomain(u))
	}
	return result, total, nil
}

func mapDAOToDomain(u *dao.User) *domain.User {
	if u == nil {
		return nil
	}
	return &domain.User{
		ID:                   u.ID,
		UUID:                 u.UUID,
		Phone:                u.Phone,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Nickname:             u.Nickname,
		AvatarURL:            u.AvatarURL,
		Bio:                  u.Bio,
		Role:                 u.Role,
		RealnameVerified:     u.RealnameVerified,
		PhotographerVerified: u.PhotographerVerified,
		Status:               u.Status,
		ReportCount:          u.ReportCount,
		Blacklisted:          u.Blacklisted,
		LastLoginAt:          u.LastLoginAt,
	}
}

func mapDomainToDAO(u *domain.User) *dao.User {
	if u == nil {
		return nil
	}
	return &dao.User{
		ID:                   u.ID,
		UUID:                 u.UUID,
		Phone:                u.Phone,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Nickname:             u.Nickname,
		AvatarURL:            u.AvatarURL,
		Bio:                  u.Bio,
		Role:                 u.Role,
		RealnameVerified:     u.RealnameVerified,
		PhotographerVerified: u.PhotographerVerified,
		Status:               u.Status,
		ReportCount:          u.ReportCount,
		Blacklisted:          u.Blacklisted,
		LastLoginAt:          u.LastLoginAt,
	}
}
