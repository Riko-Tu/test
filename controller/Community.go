package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"turan.com/web_demo/logic"
	"turan.com/web_demo/models"
)

func GetCommunity(c *gin.Context) {
	comID := c.Query("id")
	community := new(models.Community)
	atoi, err := strconv.Atoi(comID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	}

	community.CommunityID = atoi

	CommunityList := logic.GetCommunityList(community.CommunityID)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "成功", "data": CommunityList})
}

func PutCommunity(c *gin.Context) {
	community := new(models.Community)
	ParamsJson(c, community)
	community.CommunityName = GetcommunityName(community.CommunityID)
	err := logic.IsertCommunity(community)
	if errors.Is(err, logic.ComInsertFailed) {
		ResWithErr(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "成功"})
		return
	}

}

func CommunityLike(c *gin.Context) {
	value, exists := c.Get(CtxAuthUserID)
	if !exists {
		ResWithMsg(c, CodeTokeninvalid)
		return
	}
	s := value.(int64)
	Cl := new(models.CommunityLike)
	Cl.UserID = s
	ParamsJson(c, Cl)
	err := logic.LikeCommunity(Cl)
	if errors.Is(err, logic.ComInsertFailed) {
		ResWithMsg(c, CodeLikeFailed)
		return
	} else if errors.Is(err, logic.ComInsertSuccess) {
		ResWithMsg(c, CodeLikeSuccess)
		return
	} else if errors.Is(err, logic.ComLikeFailed) {
		ResWithErr(c, err)
		return
	}
}
func CommunityDisLike(c *gin.Context) {
	value, exists := c.Get(CtxAuthUserID)
	if !exists {
		ResWithMsg(c, CodeTokeninvalid)
		return
	}
	s := value.(int64)
	Cl := new(models.CommunityLike)
	Cl.UserID = s
	ParamsJson(c, Cl)
	logic.DisLikeCommunity(Cl)
}
