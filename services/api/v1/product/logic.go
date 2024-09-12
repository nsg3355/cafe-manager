package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsg3355/cafe-manager/common"
	"github.com/nsg3355/cafe-manager/config"
)

func getList(c *gin.Context) common.Result {
	// define
	var header common.Header
	var params ReqListItme
	var results []ListItme

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	// vaildation
	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}

	// query
	query, args := makeGetList(params)
	res := config.GlobalDB.Raw(query, args...).Find(&results)
	if res.Error != nil {
		return common.NewResult(http.StatusInternalServerError, res.Error.Error(), nil)
	}
	if len(results) == 0 {
		return common.NewResult(http.StatusBadRequest, "상품을 찾을 수 없습니다.", nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", Products{Products: results})
}

func getByid(c *gin.Context) common.Result {
	// define
	var header common.Header
	var params ReqByidItme
	var results ByidItme

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// vaildation
	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}

	// query
	res := config.GlobalDB.Raw(selProductByid, params.ProductId).Find(&results)
	if res.Error != nil {
		return common.NewResult(http.StatusInternalServerError, res.Error.Error(), nil)
	}
	if results.Id == 0 {
		return common.NewResult(http.StatusBadRequest, "상품을 찾을 수 없습니다.", nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", results)
}

func postItem(c *gin.Context) common.Result {
	// define
	var header common.Header
	var params ReqPostItme

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// vaildation
	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}
	if !common.Contains([]string{"small", "large"}, params.Size) {
		return common.NewResult(http.StatusBadRequest, "잘못된 상품 사이즈 입니다.", nil)
	}

	// query
	result := config.GlobalDB.Exec(insProduct,
		params.UserId,
		params.Category,
		params.Price,
		params.Cost,
		params.Name,
		common.ExtractInitialConsonants(params.Name),
		params.Description,
		params.Barcode,
		params.ExpirationDate,
		params.Size,
	)
	if result.Error != nil {
		return common.NewResult(http.StatusInternalServerError, "상품 등록에 실패했습니다. => "+result.Error.Error(), nil)
	}

	// return
	return common.NewResult(http.StatusOK, "ok", nil)
}

func putItem(c *gin.Context) common.Result {
	// define
	var header common.Header
	var params ReqPutItme

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// vaildation
	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}
	if !common.Contains([]string{"small", "large"}, params.Size) {
		return common.NewResult(http.StatusBadRequest, "잘못된 상품 사이즈 입니다.", nil)
	}

	// query
	query, args := makePutItme(params)
	result := config.GlobalDB.Exec(query, args...)
	if result.Error != nil {
		return common.NewResult(http.StatusInternalServerError, result.Error.Error(), nil)
	}
	if result.RowsAffected == 0 {
		return common.NewResult(http.StatusBadRequest, "상품을 찾을 수 없습니다.", nil)
	}
	// return
	return common.NewResult(http.StatusOK, "ok", result.RowsAffected)
}

func deleteItem(c *gin.Context) common.Result {
	// define
	var header common.Header
	var params ReqDeleteItme

	// bind
	if err := c.ShouldBindHeader(&header); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.ShouldBind(&params); err != nil {
		return common.NewResult(http.StatusBadRequest, err.Error(), nil)
	}

	// vaildation
	vaild, msg := common.ValidateToken(header.Authorization)
	if !vaild {
		return common.NewResult(http.StatusBadRequest, msg, nil)
	}

	// query
	result := config.GlobalDB.Exec(delProduct, params.ProductId)
	if result.Error != nil {
		return common.NewResult(http.StatusInternalServerError, result.Error.Error(), nil)
	}
	if result.RowsAffected == 0 {
		return common.NewResult(http.StatusBadRequest, "상품을 찾을 수 없습니다.", nil)
	}
	// return
	return common.NewResult(http.StatusOK, "ok", result.RowsAffected)
}
