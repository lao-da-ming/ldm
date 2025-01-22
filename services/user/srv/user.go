package srv

import (
	"context"
	"errors"
	"ldm/common/gen_proto/user"
)

func (u UserImpl) Profile(ctx context.Context, req *user.ProfileReq, rsp *user.ProfileRsp) error {
	rsp.Message = req.Name + "hahahah"
	return errors.New("666666")
}
