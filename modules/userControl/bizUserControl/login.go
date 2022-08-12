package bizUserControl

import (
	"context"
	"task1/common"
	"task1/component"
	"task1/component/tokenprovider"
	usermodel "task1/modules/user/modelUser"
	"task1/modules/userControl/modelUserControl"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateSession(ctx context.Context, email string) error
}

type Hasher interface {
	Hash(data string) string
}

type setTime struct {
	timeAccess  float32
	timeRefresh float32
}

func NewSetTime(timeAccess float32, timeRefresh float32) *setTime {
	return &setTime{
		timeAccess:  timeAccess,
		timeRefresh: timeRefresh,
	}
}

type TokenConfig interface {
	GetAtExp() float32
	GetRtExp() float32
}

func (timeSet *setTime) GetAtExp() float32 {
	return timeSet.timeAccess
}
func (timeSet *setTime) GetRtExp() float32 {
	return timeSet.timeRefresh
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, tkCfg TokenConfig) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tkCfg:         tkCfg,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *modelUserControl.UserLogin) (*tokenprovider.Account, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.ErrInvalidLogin(err)
	}

	passHash := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHash {
		return nil, common.ErrInvalidLogin(err)
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetAtExp())
	if err != nil {
		return nil, common.GenerateJWTFail(err)
	}
	refreshToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetRtExp())
	if err != nil {
		return nil, common.GenerateJWTFail(err)
	}
	if err := biz.storeUser.CreateSession(ctx, user.Email); err != nil {
		return nil, common.ErrInvalidLogin(err)
	}
	account := tokenprovider.Account{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &account, nil
}
