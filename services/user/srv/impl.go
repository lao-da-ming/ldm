package srv

import (
	"context"
	"gorm.io/gorm"
	"ldm/common/model"
)

type UserImpl struct {
	ctx       context.Context
	userModel *model.UserModel
}

func NewUserImpl(ctx context.Context, db *gorm.DB) *UserImpl {
	return &UserImpl{
		ctx:       ctx,
		userModel: model.NewUserModel(db),
	}
}
