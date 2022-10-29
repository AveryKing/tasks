package api

import (
	"database/sql"
	db "github.com/AveryKing/tasks/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	db      *sql.DB
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(conn *sql.DB) (*Server, error) {
	server := &Server{
		queries: db.New(conn),
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
