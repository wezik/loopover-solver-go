package api

import (
	"fmt"
	"loopover-solver/solver"
	"github.com/gin-gonic/gin"
)

type Form struct {
        State [][]int `json:"state" binding:"required"`
}

func bindEndpoints(r *gin.Engine) {
        r.GET("v1/api/ping", ping) 
        r.GET("v1/api/solve", getMoves) 
}

func StartServer() {
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
        c.JSON(200, gin.H{"moves": solver.Solve(obj.State)})
}

