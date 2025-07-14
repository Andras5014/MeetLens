package dao

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
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type UserFilter struct {
	Keyword  string
	Role     string
	Status   string
	Page     int64
	PageSize int64
}

type RealnameVerification struct {
	ID             int64     `gorm:"primaryKey;autoIncrement"`
	UserID         int64     `gorm:"index"`                                      // 外键
	RealName       string    `gorm:"type:varchar(50)"`                           // 真实姓名
	IDCardNumber   string    `gorm:"type:varchar(30)"`                           // 证件号
	IDCardFrontURL string    `gorm:"type:varchar(255)"`                          // 正面照
	IDCardBackURL  string    `gorm:"type:varchar(255)"`                          // 反面照
	Status         string    `gorm:"type:enum('pending','approved','rejected')"` // 审核状态
	Remark         string    `gorm:"type:varchar(255)"`                          // 审核备注
	SubmittedAt    time.Time `gorm:"type:timestamp"`                             // 提交时间
	ReviewedAt     time.Time `gorm:"type:timestamp"`                             // 审核时间
}

type PhotographerVerification struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	UserID       int64     `gorm:"index"`
	PortfolioURL string    `gorm:"type:varchar(255)"` // 作品集链接
	Certificates string    `gorm:"type:json"`         // 证书（JSON）
	Status       string    `gorm:"type:enum('pending','approved','rejected')"`
	Remark       string    `gorm:"type:varchar(255)"`
	SubmittedAt  time.Time `gorm:"type:timestamp"`
	ReviewedAt   time.Time `gorm:"type:timestamp"`
}

type Wallet struct {
	ID             int64   `gorm:"primaryKey;autoIncrement"`
	UserID         int64   `gorm:"uniqueIndex"`        // 一对一
	Balance        float64 `gorm:"type:decimal(10,2)"` // 可用余额
	FrozenAmount   float64 `gorm:"type:decimal(10,2)"` // 冻结金额
	TotalIncome    float64 `gorm:"type:decimal(10,2)"` // 历史总收入
	TotalWithdrawn float64 `gorm:"type:decimal(10,2)"` // 历史总提现
	UpdatedAt      time.Time
}

type UserBlacklist struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	UserID        int64  `gorm:"index"`             // 被拉黑用户
	Reason        string `gorm:"type:varchar(255)"` // 拉黑原因
	BlacklistedBy int64  `gorm:"index"`             // 操作人（管理员）
	CreatedAt     time.Time
}

type UserReport struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	ReporterID     int64  `gorm:"index"`                            // 举报人
	ReportedUserID int64  `gorm:"index"`                            // 被举报人
	Reason         string `gorm:"type:varchar(255)"`                // 举报原因
	Status         string `gorm:"type:enum('pending','processed')"` // 状态
	Remark         string `gorm:"type:varchar(255)"`                // 处理意见
	CreatedAt      time.Time
	ProcessedAt    time.Time `gorm:"type:timestamp"`
}

type WalletTransaction struct {
	ID          int64   `gorm:"primaryKey;autoIncrement"`
	WalletID    int64   `gorm:"index"` // 外键
	Type        string  `gorm:"type:enum('income','expense','freeze','unfreeze','withdraw')"`
	Amount      float64 `gorm:"type:decimal(10,2)"`
	Description string  `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
}
