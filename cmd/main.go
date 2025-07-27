package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bandanascripts/phantompay/pkg/core"
	"github.com/bandanascripts/phantompay/pkg/server/routes"
	"github.com/bandanascripts/phantompay/pkg/service/redis"
	"github.com/gin-gonic/gin"
)

func GetPort() string {
	return os.Getenv("PORT")
}

func init() {
	redis.Connect()
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	core.GenAndStoreKey(ctx, "RSA:PRIVATEKEY:", "RSA:PUBLICKEY:", 3600)

	var r = gin.Default()
	routes.RegisteredRoutes(r)

	var svr = &http.Server{Addr: ":" + GetPort(), Handler: r}

	fmt.Printf("starting server at port %s\n", GetPort())

	if err := svr.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server at port %s : %v", GetPort(), err)
	}
}
