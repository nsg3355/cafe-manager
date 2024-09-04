package product

// type Test struct {
// 	Apple  string `header:"apple" `
// 	Banana string `json:"banana,omitempty" binding:"required" gorm:"column:banana" `
// }

type ReqListItme struct {
	ProductId int    `form:"product_id" `
	Keyword   string `form:"keyword" `
}

type ReqByidItme struct {
	ProductId int `form:"product_id" binding:"required"`
}

type ReqDeleteItme struct {
	ProductId int `json:"product_id" binding:"required"`
}

type Products struct {
	Products []ListItme `json:"products" `
}

type ListItme struct {
	Id       int    `json:"id" gorm:"id" `
	Category string `json:"category" gorm:"category" `
	Price    string `json:"price" gorm:"price" `
	Name     string `json:"name" gorm:"name" `
	Size     string `json:"size" gorm:"size" `
}

type ByidItme struct {
	Id             int    `json:"id" gorm:"id" `
	UserId         string `json:"user_id" gorm:"user_id" `
	Category       string `json:"category" gorm:"category" `
	Price          string `json:"price" gorm:"price" `
	Cost           string `json:"cost" gorm:"cost" `
	Name           string `json:"name" gorm:"name" `
	Description    string `json:"description" gorm:"description" `
	Barcode        string `json:"barcode" gorm:"barcode" `
	ExpirationDate string `json:"expiration_date" gorm:"expiration_date" `
	Size           string `json:"size" gorm:"size" `
	CreatedAt      string `json:"created_at" gorm:"created_at" `
	UpdatedAt      string `json:"updated_at" gorm:"updated_at" `
}

type ReqPostItme struct {
	UserId         int    `json:"user_id" binding:"required"`
	Category       string `json:"category" binding:"required"`
	Price          string `json:"price" binding:"required"`
	Cost           string `json:"cost" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description" binding:"required"`
	Barcode        string `json:"barcode" binding:"required"`
	ExpirationDate string `json:"expiration_date" binding:"required"`
	Size           string `json:"size" binding:"required"`
}

type ReqPutItme struct {
	ProductId      int    `json:"product_id" binding:"required"`
	Category       string `json:"category"`
	Price          string `json:"price"`
	Cost           string `json:"cost"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Barcode        string `json:"barcode"`
	ExpirationDate string `json:"expiration_date"`
	Size           string `json:"size"`
}
