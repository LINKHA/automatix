package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/pb"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/usercenter"
	"github.com/LINKHA/automatix/app/usercenter/model"
	"github.com/LINKHA/automatix/common/tool"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	var userId int64
	var err error
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginOrRegister(in.AuthType, in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}

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

func (l *FastloginLogic) loginOrRegister(authType, authKey, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, authKey)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", authKey, err)
	}
	// Register
	if user == nil {
		return l.register(authType, authKey, password)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.Id, nil
}

func (l *FastloginLogic) register(authType, authKey, password string) (int64, error) {
	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = authKey
		if len(user.Nickname) == 0 {
			user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		if len(password) > 0 {
			user.Password = tool.Md5ByString(password)
		}
		insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
		}
		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = authKey
		userAuth.AuthType = authType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return userId, nil
}
