package middleware

import (
	"fmt"
	"net/http"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/global/errcode"
	"github.com/TravisRoad/gomarkit/helper"
	"github.com/TravisRoad/gomarkit/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func debugUserinfo(s sessions.Session) (model.UserInfo, error) {
	uinfo := model.UserInfo{
		ID:       1,
		Username: "admin",
		Role:     "admin",
	}
	s.Set(global.USER_INFO_KEY, uinfo)
	err := s.Save()
	return uinfo, fmt.Errorf("debugUserinfo failed to save session: %w", err)
}

func saveUserInfo(uinfo model.UserInfo, c *gin.Context) {
	c.Set("username", uinfo.Username)
	c.Set("ID", uinfo.ID)
	c.Set("role", uinfo.Role)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uinfo, ok := session.Get(global.USER_INFO_KEY).(model.UserInfo)
		mode := helper.Mode()

		// if debug mode on
		if debug := c.Query("debug"); !ok && len(debug) != 0 && (mode == global.DEV || mode == global.TEST) {
			userinfo, err := debugUserinfo(session)
			if err != nil {
				global.Logger.Error("failed to debugUserinfo", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.SessionSave,
					"msg":  err.Error(),
				})
				return
			}
			saveUserInfo(userinfo, c)
			c.Next()
			return
		}

		// or just return unauthorized
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": errcode.NotLogin,
				"msg":  "not login",
			})
			return
		}

		saveUserInfo(uinfo, c)
		c.Next()
	}
}
