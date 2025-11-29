package model

import "time"

// CREATE TABLE `posts` (
//     `id` bigint(20) NOT NULL AUTO_INCREMENT,
//     `post_id` bigint(20) NOT NULL COMMENT '帖子id',
//     `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
//     `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
//     `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
//     `community_id` bigint(20) NOT NULL COMMENT '所属社区',
//     `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
//     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//     PRIMARY KEY (`id`),
//     UNIQUE KEY `idx_post_id` (`post_id`),
//     KEY `idx_author_id` (`author_id`),
//     KEY `idx_community_id` (`community_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

type Post struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	PostID      int64     `json:"post_id"`
	Title       string    `json:"title" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	AuthorID    int64     `json:"author_id"`
	CommunityID int64     `json:"community_id" binding:"required"`
	Status      int8      `json:"status" gorm:"default:1"`
	CreateTime  time.Time `json:"create_time" gorm:"default:current_timestamp"`
	UpdateTime  time.Time `json:"update_time" gorm:"default:current_timestamp on update current_timestamp"`
}

type ApiPostInfo struct {
	AuthorName string `json:"author_name"`
	*Post
	*Community `json:"community"`
}
