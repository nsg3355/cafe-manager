package common

import (
	"regexp"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nsg3355/cafe-manager/config"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(config.GetInstance().JwtKey)

type Header struct {
	Authorization string `header:"authorization"`
}

type Result struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResult(code int, message string, data interface{}) Result {
	return Result{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

type Claims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

func IsPhoneNumber(phoneNumber string) bool {
	// 정규 표현식: 한국 핸드폰 번호 형식 검증
	re := regexp.MustCompile(`^(01[016789]{1})(\d{3,4})(\d{4})$`)
	return re.MatchString(phoneNumber)
}

func HashPassword(password string) (string, error) {
	// bcrypt 해시 생성 (비용 파라미터는 14로 설정)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Contains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

func GenerateJWT(phoneNumber string) (string, error) {
	// JWT 유효기간을 10분으로 설정
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, string) {
	if tokenString == "" {
		return false, "인증토큰이 필요합니다."
	}
	tokenString = tokenString[7:] // "Bearer " 제거
	// 1. JWT 토큰 파싱 및 유효성 검사
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	// 토큰이 올바르지 않거나 서명이 틀리면 에러 처리
	if err != nil && err == jwt.ErrSignatureInvalid {
		return false, "invalid token"
	}

	// 토큰이 만료된 경우 처리
	if !token.Valid {
		return false, "token is expired"
	}

	// 2. 데이터베이스에서 토큰이 유효한지 확인 (로그아웃 시 즉시 만료)
	var userId string
	res := config.GlobalDB.Raw("SELECT user_id FROM access_control WHERE access_token = ? AND expires_at > now()", tokenString).Find(&userId)
	if res.Error != nil || userId == "" {
		return false, "token is expired or invalid"
	}

	// 토큰이 유효하고 DB에도 존재할 경우
	return true, "ok"
}

// 초성 추출 함수
func ExtractInitialConsonants(input string) string {
	var result []rune
	var choSung = []rune{
		'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
	}

	for _, r := range input {
		// 한글 유니코드 범위: AC00 ~ D7A3
		if unicode.Is(unicode.Hangul, r) {
			// 초성 추출
			if r >= 0xAC00 && r <= 0xD7A3 {
				cho := (r - 0xAC00) / (21 * 28) // 초성 계산
				result = append(result, choSung[cho])
			}
		} else {
			// 한글이 아닌 문자는 그대로 추가
			result = append(result, r)
		}
	}

	return string(result)
}

// // 200 OK Example
// {
// 	"meta":{
// 		"code": 200, // http status code와 같은 code를 응답으로 전달
// 		"message":"ok" // 에러 발생시, 필요한 에러 메시지 전달
// 	},
// 	"data":{
// 		"products":[...]
// 	},
// }

// // 400 Bad Request Example
// {
// 	"meta":{
// 		"code": 400,
// 		"message": "잘못된 상품 사이즈 입니다."
// 	},
// 	"data": null
// }
