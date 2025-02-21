package wallet

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *handler {
	return &handler{
		db: db,
	}
}
func (h *handler) InitRoutes(route *gin.Engine) {
	route.GET("/wallet/:address/transactions")
	route.GET("/wallet/:address/balance")
}
