package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsg3355/ph-cafe-manager/common"
	"github.com/nsg3355/ph-cafe-manager/config"
)

func postSignup(c *gin.Context) common.Result {
	// define
	var params ReqUserSignup

	// bind
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// validate
	if !common.IsPhoneNumber(params.PhoneNumber) {
		return common.NewResult(http.StatusBadRequest, "유효한 핸드폰 번호가 아닙니다.", nil)
	}

	// logic
	hash, err := common.HashPassword(params.Password)
	if err != nil {
		return common.NewResult(http.StatusInternalServerError, "비밀번호 해싱 에러", nil)
	}

	// query
	_, err = config.GlobalDB.Raw(insUser, params.PhoneNumber, hash).Rows()
	if err != nil {
		return common.NewResult(http.StatusInternalServerError, err.Error(), nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", nil)
}

func postLogin(c *gin.Context) common.Result {
	// define
	var params ReqUserSignup
	var userId int
	var hash string

	// bind
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// validate
	if !common.IsPhoneNumber(params.PhoneNumber) {
		return common.NewResult(http.StatusBadRequest, "유효한 핸드폰 번호가 아닙니다.", nil)
	}

	// query
	row := config.GlobalDB.Raw(selUser, params.PhoneNumber).Row()
	row.Scan(&userId, &hash)

	// check
	if hash == "" || !common.CheckPasswordHash(params.Password, hash) {
		return common.NewResult(http.StatusBadRequest, "로그인 정보가 잘못 되었습니다.", nil)
	}

	// make token
	token, err := common.GenerateJWT(params.PhoneNumber)
	if err != nil {
		return common.NewResult(http.StatusInternalServerError, "토큰 생성 에러", nil)
	}

	// token update
	_, err = config.GlobalDB.Raw(insAccess, userId, token).Rows()
	if err != nil {
		return common.NewResult(http.StatusInternalServerError, "토큰 저장 에러", nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", "Bearer "+token)
}

func postLogout(c *gin.Context) common.Result {
	// define
	var params ReqUserLogout

	// bind
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// query
	_, err := config.GlobalDB.Raw(delAccess, params.UserId).Rows()
	if err != nil {
		return common.NewResult(http.StatusInternalServerError, "토큰 삭제 에러", nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", nil)
}

func getVerification(c *gin.Context) common.Result {
	// define
	var header common.Header

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", nil)
}
