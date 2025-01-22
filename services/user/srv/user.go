package srv

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ldm/common/gen_proto/user"
	"ldm/common/model"
)

func (u UserImpl) Profile(ctx context.Context, req *user.ProfileReq, rsp *user.ProfileRsp) error {
	var userInfo model.User
	if err := u.userModel.FindOne(u.ctx, 1, &userInfo); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	rsp.Name = userInfo.Name
	return nil
}
