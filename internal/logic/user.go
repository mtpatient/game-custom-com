package logic

import (
	"context"
	"fmt"
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	service.RegisterUser(&sUser{})
}

type (
	sUser struct {
	}
)

func (s *sUser) ReplacePassword(ctx context.Context, rp api.UserReplacePassword) error {
	if rp.NewPwd != rp.ConfirmPwd {
		return gerror.Newf(`The Password value "%s" must be the same as field repwd value "%s"`,
			rp.ConfirmPwd, rp.NewPwd)
	}
	db := dao.User.Ctx(ctx)

	one, err := db.Where("id", rp.Id).Where("password", rp.CurPwd).One()
	if one == nil {
		return gerror.New("当前密码错误")
	}

	if err != nil {
		return err
	}

	_, err = db.Where("id", rp.Id).OmitEmpty().Update(g.Map{
		"password": rp.NewPwd,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sUser) Update(ctx context.Context, user entity.User) error {
	db := dao.User.Ctx(ctx)

	_, err := db.OmitEmpty().Where("id", user.Id).Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *sUser) UserRole(ctx context.Context) (int, error) {
	v := service.Context().Get(ctx)
	if v == nil {
		return 0, gerror.New("用户未登录")
	}
	return v.User.Role, nil
}

func (s *sUser) IsLogin(ctx context.Context) (bool, error) {
	if v := service.Context().Get(ctx); v != nil && v.User != nil {
		return true, nil
	}
	return false, nil
}

func (s *sUser) GetById(ctx context.Context, id int) (entity.User, error) {
	db := dao.User.Ctx(ctx)

	var user entity.User

	err := db.FieldsEx("password").Where("id", id).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *sUser) NameExist(ctx context.Context, username string) (bool, error) {
	db := dao.User.Ctx(ctx)

	count, err := db.Where("username", username).Count()

	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func (s *sUser) Register(ctx context.Context, user api.User) error {
	if user.Repwd != user.Password {
		return gerror.Newf(`The Password value "%s" must be the same as field repwd value "%s"`,
			user.Password, user.Repwd)
	}
	exist, err := s.NameExist(ctx, user.Username)
	if err != nil {
		return err
	}
	if exist {
		return gerror.Newf(`Username "%s" is already exist!`, user.Username)
	}

	db := dao.User.Ctx(ctx)

	_, err = db.Insert(do.User{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sUser) Login(ctx context.Context, user api.User) (entity.User, string, error) {
	var u entity.User
	if ok, err := s.IsLogin(ctx); err != nil {
		return u, fmt.Sprintf("%s", service.Context().Get(ctx).Data["token"]), err
	} else {
		if ok {
			return u, fmt.Sprintf("%s", service.Context().Get(ctx).Data["token"]),
				gerror.New("User has Login!")
		}
	}
	db := dao.User.Ctx(ctx)

	err := db.Where(do.User{
		Username: user.Username,
		Password: user.Password,
	}).WhereOr(do.User{
		Email:    user.Username,
		Password: user.Password,
	}).FieldsEx("password").Scan(&u)

	if err != nil {
		return u, "", err
	}

	if u.Id > 0 {
		token := guid.S()
		err = g.Redis().SetEX(ctx, consts.TokenKey+token, u.Id, consts.TokenKeyTTL*60)
		if err != nil {
			return u, "", err
		}

		return u, token, nil
	} else {
		return u, "", gerror.New("Username or password Error!")
	}
}

func (s *sUser) Logout(ctx context.Context, token string) error {
	_, err := g.Redis().Del(ctx, consts.TokenKey+token)
	if err != nil {
		return err
	}
	return nil
}
