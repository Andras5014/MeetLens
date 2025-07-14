package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) Create(ctx context.Context, user *User) error {
	return d.db.WithContext(ctx).Create(user).Error
}

func (d *userDao) GetByID(ctx context.Context, id int64) (*User, error) {
	var u User
	err := d.db.WithContext(ctx).First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (d *userDao) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("uuid = ?", uuid).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (d *userDao) GetByPhone(ctx context.Context, phone string) (*User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (d *userDao) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (d *userDao) UpdateProfile(ctx context.Context, userID int64, nickname, avatarURL, bio string) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"nickname":   nickname,
			"avatar_url": avatarURL,
			"bio":        bio,
		}).Error
}

func (d *userDao) UpdatePassword(ctx context.Context, userID int64, passwordHash string) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("password_hash", passwordHash).Error
}

func (d *userDao) UpdateRealnameVerified(ctx context.Context, userID int64, verified bool) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("realname_verified", verified).Error
}

func (d *userDao) UpdatePhotographerVerified(ctx context.Context, userID int64, verified bool) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("photographer_verified", verified).Error
}

func (d *userDao) UpdateBlacklistStatus(ctx context.Context, userID int64, blacklisted bool) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("blacklisted", blacklisted).Error
}

func (d *userDao) UpdateStatus(ctx context.Context, userID int64, status string) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("status", status).Error
}

func (d *userDao) IncrementReportCount(ctx context.Context, userID int64) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		UpdateColumn("report_count", gorm.Expr("report_count + ?", 1)).Error
}

func (d *userDao) UpdateLastLogin(ctx context.Context, userID int64, loginAt time.Time) error {
	return d.db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).
		Update("last_login_at", loginAt).Error
}

// ListUsers 支持分页和简单筛选（示例实现）
func (d *userDao) ListUsers(ctx context.Context, filter *UserFilter) ([]*User, int64, error) {
	var users []*User
	query := d.db.WithContext(ctx).Model(&User{})

	if filter.Keyword != "" {
		kw := "%" + filter.Keyword + "%"
		query = query.Where("nickname LIKE ? OR phone LIKE ? OR email LIKE ?", kw, kw, kw)
	}

	if filter.Role != "" {
		query = query.Where("role = ?", filter.Role)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.PageSize
	err := query.Order("id DESC").Limit(int(filter.PageSize)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
