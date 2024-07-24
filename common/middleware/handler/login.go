package handler

import (
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = '2'", u.Username).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		return
	}
	err = tx.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		log.Errorf("get role error, %s", err.Error())
		return
	}
	err = u.UpdateLastLoginTime(tx, user)
	if err != nil {
		log.Errorf("update last login time error, %s", err.Error())
		return
	}
	return
}

func (u *Login) UpdateLastLoginTime(tx *gorm.DB, user SysUser) (err error) {
	err = tx.Table("sys_user").Where("user_id = ?  and status = '2'", user.UserId).Update("last_login_at", time.Now()).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	return
}
