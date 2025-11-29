package logic

import (
	"jachow/code1024/dao/mysql"
	"jachow/code1024/model"
	"jachow/code1024/pkg"
)

func CreatePost(post *model.Post) (err error) {
	post.PostID = pkg.GetID()
	return mysql.CreatePost(post)
}

func GetPost(postID int64) (apiPostInfo *model.ApiPostInfo, err error) {
	// TODO: 调用dao层获取帖子
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	community, err := mysql.QueryCommunityByID(post.CommunityID)
	if err != nil {
		return nil, err
	}
	author, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return nil, err
	}

	apiPostInfo = &model.ApiPostInfo{
		AuthorName: author.Username,
		Post:       post,
		Community:  community,
	}

	return apiPostInfo, nil
}

func GetPostList(offset, limit int64) (ApiPostList []*model.ApiPostInfo, err error) {
	// TODO: 调用dao层获取帖子列表
	postList, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	ApiPostList = make([]*model.ApiPostInfo, 0, len(postList))
	for _, post := range postList {
		community, err := mysql.QueryCommunityByID(post.CommunityID)
		if err != nil {
			return nil, err
		}
		author, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		apiPostInfo := &model.ApiPostInfo{
			AuthorName: author.Username,
			Post:       post,
			Community:  community,
		}
		ApiPostList = append(ApiPostList, apiPostInfo)
	}

	return ApiPostList, nil
}
