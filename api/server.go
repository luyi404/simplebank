package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/luyi404/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	s := &Server{
		store:  store,
		router: gin.Default(),
	}
	s.router.POST("/accounts", s.createAccount)
	s.router.GET("/accounts/:id", s.getAccount)
	s.router.GET("/accounts", s.listAccounts)
	s.router.POST("/accounts/delete/:id", s.deleteAccount)
	return s
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
