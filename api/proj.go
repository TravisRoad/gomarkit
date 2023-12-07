package api

import (
	"net/http"

	"github.com/TravisRoad/gomarkit/errcode"
	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"github.com/TravisRoad/gomarkit/service"
	"github.com/gin-gonic/gin"
)

type ProjApi struct{}

func (pa *ProjApi) AddProj(c *gin.Context) {

}

func (pa *ProjApi) RemoveProj(c *gin.Context) {

}

func (pa *ProjApi) GetProj(c *gin.Context) {
	ps := new(service.ProjService)

	page, size := getPageAndSize(c)
	projs, total, err := ps.GetProj(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.GetProjsFailed,
			"msg":  err.Error(),
		})
		return
	}

	res := make([]model.ProjDTO, len(projs))
	for i, p := range projs {
		id, err := global.Sqids.Encode([]uint64{uint64(p.ID)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": errcode.SqidsParseFailed,
				"msg":  err.Error(),
			})
			return
		}
		res[i] = model.ProjDTO{
			ID:   id,
			Name: p.Name,
			Desc: p.Desc,
			Type: p.Type,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"page":  page,
			"size":  size,
			"total": total,
			"projs": res,
		},
	})
}

func (pa *ProjApi) UpdateProj(c *gin.Context) {

}
