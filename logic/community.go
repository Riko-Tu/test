package logic

import (
	"errors"
	"time"
	"turan.com/web_demo/dao/mysql"
	"turan.com/web_demo/models"
	"turan.com/web_demo/utils"
)

var (
	ComInsertSuccess = errors.New("插入成功")
	ComInsertFailed  = errors.New("插入失败")
	ComLikeFailed    = errors.New("每个账号只能点赞一次")
	ComDisLikeFailed = errors.New("请先点赞")
)

func GetCommunityList(communityID int) []*models.CommunityRes {
	return mysql.GetCommunity(communityID)
}
func IsertCommunity(comm *models.Community) (err error) {
	comm.UpdateTime = time.Now()
	comm.CreateTime = time.Now()
	comm.CommunityTitleID = utils.GetCommunityID().Generate().Int64()
	if !mysql.InserCommunity(comm) {
		return ComInsertFailed
	}
	return ComInsertSuccess
}
func LikeCommunity(comm *models.CommunityLike) (err error) {
	stutas := mysql.IslikeStutas(comm)
	if stutas != 0 {
		return ComLikeFailed
	}
	IsLike := mysql.LikeCommunity(comm)
	if !IsLike {
		return ComInsertFailed
	}
	return ComInsertSuccess
}
func DisLikeCommunity(com *models.CommunityLike) (err error) {
	stutas := mysql.IslikeStutas(com)
	if stutas != 1 {
		return ComDisLikeFailed
	}
	dislike := mysql.Dislikecommunity(com)
	if !dislike {
		return ComInsertFailed
	}
	return ComInsertSuccess
}
