package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramonmacias/graphql/internal/handler"
	"github.com/ramonmacias/graphql/pkg/util"
)

var host, port string

func init() {
	host = util.MustGet("GQL_SERVER_HOST")
	port = util.MustGet("GQL_SERVER_PORT")
}

func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handler.Ping())
	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
