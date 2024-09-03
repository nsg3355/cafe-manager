package common

type Result struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
