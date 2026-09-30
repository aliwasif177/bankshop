package api

import (
	"fmt"

	db "github.com/aliwasif177/bankshop/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// create new http server and setup routing
func NewServer(store db.Store) Server {

	server := Server{store: store}

	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_transfer", ValidTransferParams)
		v.RegisterValidation("is_valid_currency", ValidCurrency)

	}
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	//create transfer
	router.POST("/transfers", server.createTransfer)

	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	if err != nil {
		fmt.Println("failed to start server")
		return err
	}
	return nil
}
