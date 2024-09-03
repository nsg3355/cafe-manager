package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 글로벌 DB 변수
var GlobalDB *gorm.DB

// DB 초기화
func InitDB(dsn string) {
	log.Println("###### Init Database ######")

	// MySQL 연결 설정
	var err error
	GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// 실제 연결 테스트
	sqlDB, err := GlobalDB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection from gorm: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("###### Connected to the database successfully! ######")
}
