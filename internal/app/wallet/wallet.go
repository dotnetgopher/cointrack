package wallet

import "github.com/gin-gonic/gin"

func (h *handler) GetWalletTransactions(g *gin.Context) {
	adress := g.Param("address")

	//check type of wallet
	
}