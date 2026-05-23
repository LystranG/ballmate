package config

import "time"

const (
	// JWT 签名密钥
	JWTSecret = "ballmate-secret-key-2024"

	// 数据库文件路径
	DBPath = "./data/ballmate.db"

	// Token 过期时间：7天
	TokenExpiry = 7 * 24 * time.Hour
)
