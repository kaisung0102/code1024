package mysql

import (
	"jachow/code1024/model"

	"go.uber.org/zap"
)

func QueryCommunity() (CommunityList []*model.Community, err error) {
	err = DB.Model(&model.Community{}).Find(&CommunityList).Error
	if err != nil {
		zap.L().Error("mysql query community failed", zap.Error(err))
		return nil, ErrCommunity
	}
	return CommunityList, nil
}

func QueryCommunityByID(communityID int64) (community *model.Community, err error) {
	err = DB.Model(&model.Community{}).Where("community_id = ?", communityID).First(&community).Error
	if err != nil {
		zap.L().Error("mysql query community failed", zap.Error(err))
		return nil, ErrCommunityNotExist
	}
	return community, nil
}
