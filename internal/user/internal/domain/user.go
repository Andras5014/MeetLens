package domain

import "time"

type User struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement"`
	UUID                 string    `gorm:"type:char(36);uniqueIndex"`                // 对外唯一ID
	Phone                string    `gorm:"type:varchar(20);uniqueIndex"`             // 手机号
	Email                string    `gorm:"type:varchar(100);uniqueIndex"`            // 邮箱
	PasswordHash         string    `gorm:"type:varchar(255)"`                        // 加密密码
	Nickname             string    `gorm:"type:varchar(50)"`                         // 昵称
	AvatarURL            string    `gorm:"type:varchar(255)"`                        // 头像链接
	Bio                  string    `gorm:"type:varchar(255)"`                        // 简介
	Role                 string    `gorm:"type:enum('user','photographer','admin')"` // 角色
	RealnameVerified     bool      `gorm:"type:boolean"`                             // 实名是否通过
	PhotographerVerified bool      `gorm:"type:boolean"`                             // 摄影师认证是否通过
	Status               string    `gorm:"type:enum('active','banned')"`             // 用户状态
	ReportCount          int       `gorm:"type:int;default:0"`                       // 被举报次数
	Blacklisted          bool      `gorm:"type:boolean;default:false"`               // 是否在黑名单
	LastLoginAt          time.Time `gorm:"type:timestamp"`                           // 上次登录时间
}

// UserFilter 用于后台查询
type UserFilter struct {
	Keyword  string
	Role     string
	Status   string
	Page     int
	PageSize int
}
