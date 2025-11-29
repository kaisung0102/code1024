package mysql

import (
	"jachow/code1024/model"

	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	err = DB.Create(post).Error
	if err != nil {
		zap.L().Error("mysql create post failed", zap.Error(err))
		return ErrCreatePostFailed
	}
	return err
}

func GetPostByID(postID int64) (post *model.Post, err error) {
	err = DB.Model(&model.Post{}).Where("post_id = ?", postID).First(&post).Error
	if err != nil {
		zap.L().Error("mysql get post failed", zap.Error(err))
		return nil, ErrPostNotExist
	}
	return post, nil
}

func GetPostList(offset, limit int64) (postList []*model.Post, err error) {
	// TODO: 调用dao层获取帖子列表
	err = DB.Model(&model.Post{}).Offset(int(offset)).Limit(int(limit)).Find(&postList).Error
	if err != nil {
		zap.L().Error("mysql get post list failed", zap.Error(err))
		return nil, ErrPostNotExist
	}
	return postList, nil
}
