package api

import (
	"fmt"

	db "example.com/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// create new http server and setup routing
func NewServer(store db.Store) Server {

	server := Server{store: store}

	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

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
