package srv

import (
	"gorm.io/gorm"
	"ldm/common/model"
)

type UserImpl struct {
	userModel *model.UserModel
}

func NewUserImpl(db *gorm.DB) *UserImpl {
	return &UserImpl{
		userModel: model.NewUserModel(db),
	}
}
