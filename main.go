package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
	"github.com/TryForYourLove/Usedgo/pkg/setting"
	"github.com/TryForYourLove/Usedgo/routers/router"
	"log"
	"fmt"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	/*
	r := gin.Default()
	r.GET("ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
	*/
}
