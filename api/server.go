package api

import (
	db "github.com/AveryKing/tasks/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(queries *db.Queries) (*Server, error) {
	server := &Server{
		queries: queries,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/tasks", server.createTask)
	router.GET("/tasks", server.getTasks)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
