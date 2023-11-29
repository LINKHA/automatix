package logic

import (
	"context"

	"automatix/app/usercenter/cmd/rpc/internal/svc"
	"automatix/app/usercenter/cmd/rpc/pb"
	"automatix/app/usercenter/cmd/rpc/usercenter"
	"automatix/app/usercenter/model"
	"automatix/common/tool"
	"automatix/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FastloginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFastloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastloginLogic {
	return &FastloginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FastloginLogic) Fastlogin(in *pb.FastLoginReq) (*pb.FastLoginResp, error) {

	newLoginLogic := NewLoginLogic(l.ctx, l.svcCtx)
	newLoginLogic.loginByMobile()

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &usercenter.FastLoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginOrRegister(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		registerLogic := NewRegisterLogic(l.ctx, l.svcCtx)
		registerLogic.Register(&usercenter.RegisterReq{
			Mobile:   mobile,
			Nickname: "",
			Password: password,
		})
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.Id, nil
}
