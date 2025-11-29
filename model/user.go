package model

import "time"

// 建表语句 uers
// CREATE TABLE `users` (
// `id` bigint(20) NOT NULL AUTO_INCREMENT,
// `user_id` bigint(20) NOT NULL,
// `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
// `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
// `email` varchar(64) COLLATE utf8mb4_general_ci,
// `gender` tinyint(4) NOT NULL DEFAULT '0',
// `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
// CURRENT_TIMESTAMP,
// PRIMARY KEY (`id`),
// UNIQUE KEY `idx_username` (`username`) USING BTREE,
// UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

type User struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	UserID     int64     `json:"user_id" gorm:"uniqueIndex"`
	Username   string    `json:"username" gorm:"uniqueIndex"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Gender     int       `json:"gender" gorm:"default:0"`
	CreateTime time.Time `json:"create_time" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
