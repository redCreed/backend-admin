package service

import (
	"back-admin/api/models"
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/driver"
	"back-admin/internal/app/sys/dto"
	"back-admin/pkg/queue"
	"back-admin/pkg/utils"
	"back-admin/pkg/xerr"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"runtime"
	"strconv"
	"time"
)

//UserList 用户列表
func (u *SysSrv) UserList(param *dto.UserListParam) (*dto.UserListResp, error) {
	data, err := u.db.Sys().UserList(param)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *SysSrv) UserLogin(c *gin.Context, param *dto.UserLoginReq) (*dto.UserLoginResp, error) {
	log := &models.SysLoginLog{
		Username:      param.Username,
		Ipaddr:        c.RemoteIP(),
		LoginLocation: "",
		Browser:       "",
		Os:            runtime.GOOS,
		Platform:      c.Request.UserAgent(),
		LoginTime:     time.Now(),
		Remark:        "",
		Msg:           "登录成功",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	defer func() {

		if d, err := json.Marshal(log); err != nil {
			fmt.Println("log marshal:", err)
			return
		} else {
			msg := &queue.Message{}
			msg.SetId(uuid.New().String())
			msg.SetKey(common.LoginLog)
			msg.SetErrorCount(0)
			msg.SetValues(d)
			//推送内存队列
			driver.Instance.Queue.Add(msg)
		}
	}()
	resp := new(dto.UserLoginResp)
	user, err := u.db.Sys().GetUserByField("username", param.Username)
	if err != nil {
		log.Msg = err.Error()
		return resp, err
	}
	log.Status = user.Status
	//前端2次md5加密后在传递到后端
	password := utils.Md5(param.Password + user.Salt)
	if password != user.Password {
		log.Msg = xerr.ErrMapMsg(xerr.PasswordErr)
		return resp, xerr.NewErrCode(xerr.PasswordErr)
	}
	log.CreateBy = user.UserId
	//获取用户角色id
	roleId, err := u.db.Sys().GetRoleIdByUserId(user.UserId)
	if err != nil {
		log.Msg = err.Error()
		return nil, err
	}
	roleInfo := make([]string, 0)
	for _, v := range roleId {
		roleInfo = append(roleInfo, strconv.Itoa(v))
	}
	expireTime := time.Now().Add(2 * time.Hour)
	claim := utils.CustomClaims{
		UserId:   user.UserId,
		Roles:    roleInfo,
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := utils.CreateToken(claim)
	if err != nil {
		log.Msg = err.Error()
		return nil, err
	}
	resp.Token = token

	return resp, nil
}

func (u *SysSrv) Refresh(param *dto.UserToken) (*dto.UserToken, error) {
	claims, err := utils.ParseToken(param.Token)
	if err != nil {
		return nil, err
	}

	//token过期重新登录
	if claims.ExpiresAt > time.Now().Unix() {
		return nil, xerr.NewErrCode(xerr.TokenExpired)
	}
	expire, err := strconv.Atoi(driver.Conf.Jwt.Expire)
	if err != nil {
		return nil, err
	}
	claims.ExpiresAt = time.Now().Add(time.Duration(expire) * time.Second).Unix()

	newToken, err := utils.CreateToken(*claims)
	if err != nil {
		return nil, err
	}
	resp := new(dto.UserToken)
	resp.Token = newToken
	return resp, nil
}

func (u *SysSrv) UpdateUserRole(param *dto.UpdateUserRoleReq, userId int) error {
	user, err := u.db.Sys().GetUserById(userId)
	if err != nil {
		return err
	}
	if len(user.Username) == 0 {
		return xerr.NewErrCode(xerr.UserNoExist)
	}

	return u.updateUserRoles(userId, param.RoleId)
}

func (u *SysSrv) updateUserRoles(userId int, roleId []int) error {
	//判断角色信息
	if roles, err := u.db.Sys().GetRoleByRoleId(roleId); err != nil {
		return err
	} else {
		if len(roles) != len(roleId) {
			return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "角色参数错误")
		}
	}

	//删除全部角色
	if err := u.db.Sys().DeleteRoleByUserId(userId); err != nil {
		return err
	}
	//新增角色
	if err := u.db.Sys().AddUserRole(userId, roleId); err != nil {
		return err
	}

	return nil
}

func (u *SysSrv) AddUser(c *gin.Context, param *dto.AddUserReq) error {
	salt := utils.GetRandomStr(4)
	pw := utils.Md5(utils.Md5(param.Password) + salt)

	user := &models.SysUser{
		Username: param.Username,
		Password: pw,
		NickName: param.NickName,
		Phone:    param.Phone,
		Salt:     salt,
		Avatar:   param.Avatar,
		Sex:      param.Sex,
		Email:    param.Email,
		Remark:   param.Remark,
		Status:   1,
		ControlBy: models.ControlBy{
			CreateBy: c.GetInt("user_id"),
			UpdateBy: 0,
		},
		ModelTime: models.ModelTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := u.db.Sys().AddUser(user); err != nil {
		return err
	}
	//判断角色信息
	if roles, err := u.db.Sys().GetRoleByRoleId(param.RoleId); err != nil {
		return err
	} else {
		if len(roles) != len(param.RoleId) {
			return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "角色参数错误")
		}
	}

	//新增角色
	if err := u.db.Sys().AddUserRole(user.UserId, param.RoleId); err != nil {
		return err
	}
	return nil
}

func (u *SysSrv) UpdateUser(c *gin.Context, param *dto.UpdateUserReq, userId int) error {
	user := &models.SysUser{
		UserId:   userId,
		NickName: param.NickName,
		Phone:    param.Phone,
		Avatar:   param.Avatar,
		Sex:      param.Sex,
		Email:    param.Email,
		Remark:   param.Remark,
		Status:   1,
		ControlBy: models.ControlBy{
			UpdateBy: c.GetInt("user_id"),
		},
		ModelTime: models.ModelTime{
			UpdatedAt: time.Now(),
		},
	}
	if param.Password != "" {
		salt := utils.GetRandomStr(4)
		pw := utils.Md5(utils.Md5(param.Password) + salt)
		user.Password = pw
		user.Salt = salt
	}

	if err := u.db.Sys().UpdateUser(user); err != nil {
		return err
	}

	return u.updateUserRoles(userId, param.RoleId)
}

func (u *SysSrv) DeleteUser(c *gin.Context, userId int) error {
	var (
		user *models.SysUser
		err  error
	)
	if user, err = u.db.Sys().GetUserById(userId); err != nil {
		return err
	}

	if err := u.db.Sys().DeleteUser(userId); err != nil {
		return err
	}

	//删除用户角色
	if err = u.db.Sys().DeleteRoleByUserId(user.UserId); err != nil {
		return err
	}
	return nil
}
