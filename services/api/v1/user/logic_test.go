package user_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nsg3355/ph-cafe-manager/common"
	"github.com/nsg3355/ph-cafe-manager/config"
	"github.com/nsg3355/ph-cafe-manager/services/api/v1/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gin 라우터 설정 함수
func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// 테스트 DB 초기화
func initTestDB(t *testing.T) {
	// 테스트 환경에 맞는 DSN으로 초기화 (Test DB 설정)
	dsn := "root:root123@tcp(127.0.0.1:3306)/payhere_test?charset=utf8mb4&parseTime=True&loc=Asia%2FSeoul"
	var err error
	config.GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	log.Println("Test DB connected")
}

// 데이터베이스 초기화 (테스트 후 DB를 정리)
func cleanupTestDB(t *testing.T) {
	config.GlobalDB.Exec("DELETE FROM user_info")
	config.GlobalDB.Exec("DELETE FROM access_control")
	log.Println("Test DB cleaned")
}

// TestPostSignup 테스트
func TestPostSignup(t *testing.T) {
	initTestDB(t) // Test DB 연결
	// defer cleanupTestDB(t) // Test 후 DB 초기화

	// 라우터 설정
	router := SetupRouter()
	router.POST("/signup", user.PostSignup)

	// 요청 파라미터 설정
	params := user.ReqUserSignup{
		PhoneNumber: "01012345678",
		Password:    "strongpassword",
	}

	// 요청 본문을 JSON으로 변환
	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 결과 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// TestPostLogin 테스트
func TestPostLogin(t *testing.T) {
	initTestDB(t) // Test DB 연결
	// defer cleanupTestDB(t) // Test 후 DB 초기화

	router := SetupRouter()
	router.POST("/login", user.PostLogin)

	params := user.ReqUserSignup{
		PhoneNumber: "01012345678",
		Password:    "strongpassword",
	}

	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 결과 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// TestPostLogout 테스트
func TestPostLogout(t *testing.T) {
	initTestDB(t) // Test DB 연결
	// defer cleanupTestDB(t) // Test 후 DB 초기화

	router := SetupRouter()
	router.POST("/logout", user.PostLogout)

	params := user.ReqUserLogout{
		UserId: "1",
	}

	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/logout", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 결과 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}
