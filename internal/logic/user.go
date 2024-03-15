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
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	service.RegisterUser(&sUser{})
}

type (
	sUser struct {
	}
)

func (s *sUser) Follow(ctx context.Context, follow api.UserFollow) error {
	uid := service.Context().Get(ctx).User.Id
	if uid == follow.Id {
		return gerror.New("不能关注自己")
	}
	err := dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if follow.Operate == 1 {
			// 判断是否已经关注
			one, err := tx.Ctx(ctx).GetOne("Select * From `follow` where user_id=? and follow_user_id=? and delete_time is null limit 1", uid, follow.Id)

			if one != nil && err == nil {
				return gerror.New("已经关注了")
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"follow_count": gdb.Raw("follow_count + 1"),
			}, g.Map{
				"id": uid,
			})
			if err != nil {
				return err
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"fans_count": gdb.Raw("fans_count+1"),
			}, g.Map{
				"id": follow.Id,
			})
			if err != nil {
				return err
			}

			_, err = tx.Ctx(ctx).Insert("follow", g.Map{
				"user_id":        uid,
				"follow_user_id": follow.Id,
			})
			if err != nil {
				return err
			}
		} else if follow.Operate == 2 {
			one, err := tx.Ctx(ctx).GetOne("Select * From `follow` where user_id=? and follow_user_id=? and delete_time is null limit 1", uid, follow.Id)

			if one == nil || err != nil {
				return gerror.New("没有关注")
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"follow_count": gdb.Raw("follow_count - 1"),
			}, g.Map{
				"id": uid,
			})
			if err != nil {
				return err
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"fans_count": gdb.Raw("fans_count-1"),
			}, g.Map{
				"id": follow.Id,
			})
			if err != nil {
				return err
			}

			_, err = tx.Ctx(ctx).Delete("follow", g.Map{
				"user_id":        uid,
				"follow_user_id": follow.Id,
			})
			if err != nil {
				return err
			}
		} else {
			return gerror.New("参数错误")
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("关注失败")
	}

	return nil
}

func (s *sUser) ResetPwd(ctx context.Context, rs api.ResetPwd) error {
	if rs.NewPwd != rs.ConfirmPwd {
		return gerror.Newf(`The Password value "%s" must be the same as field repwd value "%s"`,
			rs.ConfirmPwd, rs.NewPwd)
	}
	code, _ := g.Redis().Get(ctx, consts.VerifyCodeKey+rs.Username)

	if code.String() != rs.Code {
		return gerror.New("验证码错误")
	}
	db := dao.User.Ctx(ctx)

	_, err := db.Where("username", rs.Username).OmitEmpty().Update(g.Map{
		"password": rs.NewPwd,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("修改密码失败")
	}

	_, err = g.Redis().Del(ctx, consts.VerifyCodeKey+rs.Username)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("删除验证码失败")
	}

	return nil
}

func (s *sUser) GetAuthCode(ctx context.Context, str string) error {
	db := dao.User.Ctx(ctx)

	one, err := db.Where("email", str).WhereOr("username", str).One()
	if err != nil {
		return err
	}
	if one == nil {
		return gerror.Newf("%s 用户不存在", str)
	}
	// 生成6位数字验证码
	code := grand.S(6)
	content := "你的验证码为：<h1>" + code + "</h1> <p>有效期为5分钟。<p>"
	var user entity.User
	err = one.Struct(&user)
	if err != nil {
		return err
	}
	err = utility.SendEmail(content, []string{user.Email})
	if err != nil {
		return gerror.New("发送验证码失败，稍后再试!")
	}
	err = g.Redis().SetEX(ctx, consts.VerifyCodeKey+str, code, 300)
	if err != nil {
		return gerror.New("发送验证码失败，稍后再试!")
	}

	return nil
}

func (s *sUser) ReplacePassword(ctx context.Context, rp api.UserReplacePassword) error {
	uid := service.Context().Get(ctx).User.Id
	if uid != rp.Id {
		return gerror.New("无权限")
	}
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
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，修改密码失败")
	}

	_, err = db.Where("id", rp.Id).OmitEmpty().Update(g.Map{
		"password": rp.NewPwd,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，修改密码失败")
	}

	return nil
}

func (s *sUser) Update(ctx context.Context, user entity.User) error {
	uid := service.Context().Get(ctx).User.Id
	if uid != user.Id {
		return gerror.New("无权限")
	}
	db := dao.User.Ctx(ctx)

	one, err := db.Where("username", user.Username).WhereNot("id", user.Id).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，更新用户失败")
	}
	if one != nil {
		return gerror.Newf(`Username "%s" is already exist!`, user.Username)
	}

	one, _ = db.Where("email", user.Email).WhereNot("id", user.Id).One()
	if one != nil {
		return gerror.Newf(`Email "%s" is already exist!`, user.Email)
	}
	_, err = db.OmitEmpty().Where("id", user.Id).Update(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，更新用户失败")
	}

	return nil
}

func (s *sUser) UserRole(ctx context.Context) (int, error) {
	user := service.Context().Get(ctx).User
	if user == nil {
		return 0, gerror.New("用户未登录")
	}
	return user.Role, nil
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
		g.Log().Error(ctx, err)
		return user, gerror.New("服务器错误，获取用户失败")
	}

	return user, nil
}

func (s *sUser) NameExist(ctx context.Context, username string) (bool, error) {
	db := dao.User.Ctx(ctx)

	count, err := db.Where("username", username).Count()

	if err != nil {
		g.Log().Error(ctx, err)
		return false, gerror.New("服务器错误，获取用户失败")
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
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，注册失败！")
	}
	if exist {
		return gerror.Newf(`Username "%s" is already exist!`, user.Username)
	}

	db := dao.User.Ctx(ctx)

	count, _ := db.Where("email", user.Email).Count()
	if count != 0 {
		return gerror.Newf(`Email "%s" is already exist!`, user.Email)
	}

	_, err = db.Insert(do.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Sex:      2,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("服务器错误，注册失败！")
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
		g.Log().Error(ctx, err)
		return u, "", gerror.New("服务器错误！")
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
		g.Log().Error(ctx, err)
		return gerror.New("Logout Error!")
	}
	return nil
}
