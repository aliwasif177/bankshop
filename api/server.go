package api

import (
	"fmt"

	db "github.com/aliwasif177/bankshop/db/sqlc"
	"github.com/aliwasif177/bankshop/token"
	"github.com/aliwasif177/bankshop/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

// create new http server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}
	server.setUpRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_transfer", ValidTransferParams)
		v.RegisterValidation("is_valid_currency", ValidCurrency)

	}

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/")
	authRoutes.Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.PUT("/accounts/:id", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	//create transfer
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

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
