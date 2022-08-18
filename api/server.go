package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hariprathap-hp/backend_masterclass/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//bind.validator.Engine

	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/delete/:id", server.deleteAccount)
	router.GET("/accounts", server.listAccount)
	router.PATCH("/accounts", server.updateAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}