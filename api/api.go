package api

import (
	"fmt"
	"loopover-solver/solver"
	"os"

	"github.com/gin-gonic/gin"
)

type Form struct {
	Board [][]int `json:"board" binding:"required"`
}

func bindEndpoints(r *gin.Engine) {
	r.Use(corsMiddleware())
	r.GET("v1/api/ping", ping)
	r.POST("v1/api/solve", getMoves)
}

func StartServer() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	bindEndpoints(r)
	
	

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func getMoves(c *gin.Context) {
	obj := Form{}
	if err := c.ShouldBindJSON(&obj); err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"moves": solver.Solve(obj.Board)})
	c.Header("Access-Control-Allow-Origin", "*")
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
