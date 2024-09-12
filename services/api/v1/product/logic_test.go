package product_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nsg3355/cafe-manager/common"
	"github.com/nsg3355/cafe-manager/config"
	"github.com/nsg3355/cafe-manager/services/api/v1/product"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 테스트 전용 라우터 설정
func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// 테스트 DB 초기화
func initTestDB(t *testing.T) {
	// 테스트 환경에 맞는 DSN으로 초기화 (Test DB 설정)
	dsn := "root:root123@tcp(127.0.0.1:3306)/starbucks_test"
	var err error
	config.GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	log.Println("Test DB connected")
}

// DB 정리
func cleanupTestDB(t *testing.T) {
	config.GlobalDB.Exec("DELETE FROM product_info")
}

// 테스트: GetList
func TestGetList(t *testing.T) {
	initTestDB(t) // 테스트 DB 초기화
	// defer cleanupTestDB(t) // 테스트 후 DB 정리

	// 라우터 설정
	router := SetupRouter()
	router.GET("/products", product.GetList)

	// 요청 파라미터 설정
	req, _ := http.NewRequest("GET", "/products?product_id=1&keyword=ㅋ", nil)
	req.Header.Set("Authorization", "Bearer "+testJWT())

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 본문 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// 테스트: GetByid
func TestGetByid(t *testing.T) {
	initTestDB(t) // 테스트 DB 초기화
	// defer cleanupTestDB(t) // 테스트 후 DB 정리

	// 라우터 설정
	router := SetupRouter()
	router.GET("/product", product.GetByid)

	// 요청 파라미터 설정
	req, _ := http.NewRequest("GET", "/product?product_id="+strconv.Itoa(testProductId()), nil)
	req.Header.Set("Authorization", "Bearer "+testJWT())

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 본문 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// 테스트: PostItem
func TestPostItem(t *testing.T) {
	initTestDB(t) // 테스트 DB 초기화
	// defer cleanupTestDB(t) // 테스트 후 DB 정리

	// 라우터 설정
	router := SetupRouter()
	router.POST("/product", product.PostItem)

	// 요청 파라미터 설정
	params := product.ReqPostItme{
		UserId:         1,
		Category:       "음료",
		Price:          "5000",
		Cost:           "4500",
		Name:           "아메리카노",
		Description:    "맛있는 아메리카노",
		Barcode:        "1234567890123",
		ExpirationDate: "2024-12-31",
		Size:           "small",
	}

	// JSON 변환
	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+testJWT())
	req.Header.Set("Content-Type", "application/json")

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 본문 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// 테스트: PutItem
func TestPutItem(t *testing.T) {
	initTestDB(t) // 테스트 DB 초기화
	// defer cleanupTestDB(t) // 테스트 후 DB 정리

	// 라우터 설정
	router := SetupRouter()
	router.PUT("/product", product.PutItem)

	// 요청 파라미터 설정
	params := product.ReqPutItme{
		ProductId:      testProductId(),
		Category:       "음료",
		Price:          "5500",
		Name:           "카페라떼",
		Description:    "고소한 카페라떼",
		Barcode:        "1234567890123",
		ExpirationDate: "2024-12-31",
		Size:           "large",
	}

	// JSON 변환
	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("PUT", "/product", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+testJWT())
	req.Header.Set("Content-Type", "application/json")

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 본문 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

// 테스트: DeleteItem
func TestDeleteItem(t *testing.T) {
	initTestDB(t) // 테스트 DB 초기화
	// defer cleanupTestDB(t) // 테스트 후 DB 정리

	// 라우터 설정
	router := SetupRouter()
	router.DELETE("/product", product.DeleteItem)

	// 요청 파라미터 설정
	params := product.ReqDeleteItme{
		ProductId: testProductId(),
	}

	// JSON 변환
	jsonValue, _ := json.Marshal(params)
	req, _ := http.NewRequest("DELETE", "/product", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+testJWT())
	req.Header.Set("Content-Type", "application/json")

	// 응답 기록
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 본문 검증
	var response common.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Meta.Message)
}

func testJWT() string {
	token, err := common.GenerateJWT("01012345678")
	if err != nil {
		log.Println("Test Token error")
	}
	// 외래키 제약 조건 무시
	config.GlobalDB.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	query := `
	INSERT INTO access_control (user_id, access_token, expires_at) 
	VALUES (1, ?, NOW() + INTERVAL 10 MINUTE)
	ON DUPLICATE KEY UPDATE 
    access_token = VALUES(access_token),
	expires_at = NOW() + INTERVAL 10 MINUTE;
	`
	config.GlobalDB.Exec(query, token)
	return token
}

func testProductId() int {
	var productId int
	res := config.GlobalDB.Raw("SELECT MAX(id) FROM product_info;").Find(&productId)
	if res.Error != nil {
		log.Println("Test Get Id error")
	}
	return productId
}
