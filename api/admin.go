package api

import (
	"net/http"
	"strconv"

	"github.com/TravisRoad/gomarkit/errcode"
	"github.com/TravisRoad/gomarkit/model"
	"github.com/TravisRoad/gomarkit/service"
	"github.com/gin-gonic/gin"
)

type AdminApi struct{}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type GetUsersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Page  int              `json:"page"`
		Size  int              `json:"size"`
		Total int64            `json:"total"`
		Users []model.UserInfo `json:"users"`
	} `json:"data"`
}

func (aa *AdminApi) GetUsers(c *gin.Context) {
	us := new(service.UserService)

	page, size := getPageAndSize(c)

	users, total, err := us.GetUsers(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.GetUsersFailed,
			"msg":  err.Error(),
		})
		return
	}

	res := GetUsersResponse{}
	res.Code = http.StatusOK
	res.Data.Page = page
	res.Data.Size = size
	res.Data.Total = total
	var usersInfos []model.UserInfo
	for _, u := range users {
		usersInfos = append(usersInfos, model.UserInfo{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	res.Data.Users = usersInfos

	c.JSON(http.StatusOK, res)
}

func (aa *AdminApi) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.UpdateUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	u := model.User{}
	u.ID = uint(id)
	u.Password = req.Password
	u.Role = req.Role

	us := new(service.UserService)
	if err := us.UpdateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.UpdateUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AdminApi) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	us := new(service.UserService)
	if err := us.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.DeleteUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AdminApi) AddUser(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	us := new(service.UserService)
	if err := us.AddUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.AddUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
