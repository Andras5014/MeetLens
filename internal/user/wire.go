//go:build wireinject

package user

import (
	"MeetLens/internal/user/internal/repo"
	"MeetLens/internal/user/internal/repo/dao"
	"MeetLens/internal/user/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// 定义User相关依赖的wire集合
var UserSet = wire.NewSet(
	dao.NewUserDao,
	repo.NewUserRepo,
	service.NewUserService,
)

// 注入函数
func InitializeUserService(db *gorm.DB) (service.UserService, error) {
	wire.Build(UserSet)
	return nil, nil
}
