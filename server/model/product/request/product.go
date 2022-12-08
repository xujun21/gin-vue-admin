package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"time"
)

type ProductSearch struct {
	product.Product
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartExp_date  *time.Time `json:"startExp_date" form:"startExp_date"`
	EndExp_date    *time.Time `json:"endExp_date" form:"endExp_date"`
	StartPrice     *float64   `json:"startPrice" form:"startPrice"`
	EndPrice       *float64   `json:"endPrice" form:"endPrice"`
	StartVat       *float64   `json:"startVat" form:"startVat"`
	EndVat         *float64   `json:"endVat" form:"endVat"`
	StartStore     *int       `json:"startStore" form:"startStore"`
	EndStore       *int       `json:"endStore" form:"endStore"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
