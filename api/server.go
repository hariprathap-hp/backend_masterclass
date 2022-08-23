package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/hariprathap-hp/backend_masterclass/db/sqlc"
	"github.com/hariprathap-hp/backend_masterclass/token"
	"github.com/hariprathap-hp/backend_masterclass/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create a token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker}
	router := gin.Default()

	// bind.validator.Engine
	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	authRoutes := router.Group("/", authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/delete/:id", server.deleteAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.PATCH("/accounts", server.updateAccount)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}
