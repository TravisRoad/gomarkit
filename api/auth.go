package api

import (
	"net/http"

	"github.com/TravisRoad/gomarkit/errcode"
	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"github.com/TravisRoad/gomarkit/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (aa *AuthApi) Login(c *gin.Context) {
	session := sessions.Default(c)
	if uinfo, ok := session.Get(global.USER_INFO_KEY).(model.UserInfo); ok {
		global.Logger.Info("already login", zap.String("username", uinfo.Username))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}

	req := LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": errcode.LoginFailed,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid request", zap.Error(err))
		return
	}

	as := new(service.AuthService)
	user, err := as.Login(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.UsernameOrPwd,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid login auth", zap.String("username", req.Username), zap.Error(err))
		return
	}

	userInfo := model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	session.Set(global.USER_INFO_KEY, userInfo)
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errcode.SessionSave,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AuthApi) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errcode.SessionSave,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AuthApi) IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	uinfo, ok := session.Get(global.USER_INFO_KEY).(model.UserInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.NotLogin,
			"msg":  "not login",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": uinfo,
	})

}
