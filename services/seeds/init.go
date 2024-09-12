package seeds

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nsg3355/cafe-manager/config"
)

// DB 초기화
func InitSQL() {
	// 현재 작업 디렉토리 확인
	cwd, err := os.Getwd() // 절대 경로로 현재 디렉토리 가져오기
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	// ddl 폴더의 절대 경로 설정
	dirPath := filepath.Join(cwd, "services", "seeds", "ddl") // 절대 경로로 경로 설정

	// 디렉토리 읽기
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// 각 파일 실행
	for _, file := range files {
		// 파일이 .sql로 끝나는지 확인
		if filepath.Ext(file.Name()) == ".sql" {
			filePath := filepath.Join(dirPath, file.Name())

			// SQL 파일 읽기
			sqlContent, err := os.ReadFile(filePath)
			if err != nil {
				log.Println("Failed to read SQL file: %v", err)
			}

			// SQL 파일 실행
			result := config.GlobalDB.Exec(string(sqlContent))
			if result.Error != nil {
				log.Println("Failed to execute SQL file (%s): %v", file.Name(), result.Error.Error())
			}

			fmt.Printf("SQL file %s executed successfully.\n", file.Name())
		}
	}

	log.Println("###### All SQL files executed and connected to the database successfully! ######")
}
