package main

import (
	"fmt"
	"loopover-solver/solver"
	"github.com/gin-gonic/gin"
)

type Form struct {
        State [][]int `json:"state" binding:"required"`
}

func main() {
        r := gin.Default()

        r.GET("v1/api/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "pong",
                })
        })
        r.GET("v1/api/solve", func(c *gin.Context) {
                obj := Form{}
                if err := c.ShouldBindJSON(&obj); err != nil {
                        fmt.Println(err)
                }
                c.JSON(200, gin.H{
                        "moves": solver.Solve(obj.State),
                })
        })

        err := r.Run()
        if err != nil {
                panic(err)
        }
}

