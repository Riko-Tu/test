package mysql

import (
	"turan.com/web_demo/models"
)

func GetCommunity(Cid int) []*models.CommunityRes {
	sqlstr := "select community_id,community_name,introduction,community_title_id from community where community_id=?"
	community := make([]*models.CommunityRes, 100, 200) //里面定义了指定但是community这个变量不是指定类型
	mysql.Raw(sqlstr, Cid).Scan(&community)
	return community
}
func InserCommunity(community *models.Community) bool {
	sqlstr := "insert into community(community_title_id,community_id,community_name,introduction,create_time,update_time) value(?,?,?,?,?,?)"
	affected := mysql.Exec(sqlstr, community.CommunityTitleID, community.CommunityID, community.CommunityName, community.Introduction, community.CreateTime, community.UpdateTime).RowsAffected
	if affected == 0 {
		return false
	}
	return true
}
func LikeCommunity(commLike *models.CommunityLike) bool {
	sqlstr := "insert into community_like(user_id,community_title_id,liketo) VALUE( ?, ?, ?)"
	affected := mysql.Exec(sqlstr, commLike.UserID, commLike.CommunityTitleID, commLike.LikeTo).RowsAffected
	if affected == 0 {
		return false
	}
	return true

}

func IslikeStutas(com *models.CommunityLike) int {
	m := new(models.CommunityLike)
	sqlstr := "select liketo from community_like where user_id=? and community_title_id=? "
	mysql.Raw(sqlstr, com.UserID, com.CommunityTitleID).Scan(m)
	return m.LikeTo
}

func Dislikecommunity(comm *models.CommunityLike) bool {
	sqlstr := "update community_like set liketo=? where user_id=? and community_title_id=? "
	affected := mysql.Exec(sqlstr, comm.LikeTo, comm.UserID, comm.CommunityTitleID).RowsAffected
	if affected == 0 {
		return false
	}
	return true
}
