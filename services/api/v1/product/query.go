package product

import (
	"strings"

	"github.com/nsg3355/ph-cafe-manager/common"
)

var selProductList = `
SELECT
	id
	, category
	, price
	, name
	, size
FROM product_info
WHERE 1=1
#product_id
#keyword
ORDER BY id ASC
LIMIT 10;
`

var selProductByid = `
SELECT
	id
	, user_id
	, category
	, price
	, cost
	, name
	, description
	, barcode
	, expiration_date
	, size
	, created_at
	, updated_at
FROM product_info
WHERE id = ?;
`

var insProduct = `
INSERT INTO product_info
(user_id, category, price, cost, name, initial, description, barcode, expiration_date, size)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
`

var updProduct = `
UPDATE product_info SET
	#category
	#price
	#cost
	#name
	#initial
	#description
	#barcode
	#expiration_date
	#size
	updated_at = NOW()
WHERE id = ?;
`

var delProduct = `
DELETE FROM product_info WHERE id = ?;
`

func makeGetList(params ReqListItme) (string, []interface{}) {
	args := []interface{}{}
	query := selProductList
	if params.ProductId != 0 {
		query = strings.ReplaceAll(query, "#product_id", " AND id > ?")
		args = append(args, params.ProductId)
	}
	if params.Keyword != "" {
		query = strings.ReplaceAll(query, "#keyword", " AND (name LIKE CONCAT('%', ?, '%') OR initial LIKE CONCAT('%', ?, '%')) ")
		args = append(args, params.Keyword)
		args = append(args, params.Keyword)
	}
	return query, args
}

func makePutItme(params ReqPutItme) (string, []interface{}) {
	args := []interface{}{}
	query := updProduct
	if params.Category != "" {
		query = strings.ReplaceAll(query, "#category", "category = ?,")
		args = append(args, params.Category)
	}
	if params.Price != "" {
		query = strings.ReplaceAll(query, "#price", "price = ?,")
		args = append(args, params.Price)
	}
	if params.Cost != "" {
		query = strings.ReplaceAll(query, "#cost", "cost = ?,")
		args = append(args, params.Cost)
	}
	if params.Name != "" {
		query = strings.ReplaceAll(query, "#name", "name = ?,")
		query = strings.ReplaceAll(query, "#initial", "initial = ?,")
		args = append(args, params.Name)
		args = append(args, common.ExtractInitialConsonants(params.Name))
	}
	if params.Description != "" {
		query = strings.ReplaceAll(query, "#description", "description = ?,")
		args = append(args, params.Description)
	}
	if params.Barcode != "" {
		query = strings.ReplaceAll(query, "#barcode", "barcode = ?,")
		args = append(args, params.Barcode)
	}
	if params.ExpirationDate != "" {
		query = strings.ReplaceAll(query, "#expiration_date", "expiration_date = ?,")
		args = append(args, params.ExpirationDate)
	}
	if params.Size != "" {
		query = strings.ReplaceAll(query, "#size", "size = ?,")
		args = append(args, params.Size)
	}
	args = append(args, params.ProductId)
	return query, args
}
