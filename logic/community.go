package logic

import (
	"jachow/code1024/dao/mysql"
	"jachow/code1024/model"
)

func GetCommunityList() (CommunityList []*model.Community, err error) {
	return mysql.QueryCommunity()
}

func GetCommunityByID(communityID int64) (community *model.Community, err error) {
	return mysql.QueryCommunityByID(communityID)
}
