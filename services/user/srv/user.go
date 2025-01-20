package srv

import (
	"context"
	"ldm/common/gen_proto/user"
)

func (u UserImpl) Profile(ctx context.Context, req *user.ProfileReq, rsp *user.ProfileRsp) error {
	rsp.Message = req.Name + "hahahah"
	return nil
}
