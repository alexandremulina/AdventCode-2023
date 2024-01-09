package api

import (
	"fmt"
	"go-advent/util"

	"github.com/gin-gonic/gin"
)

// type MapColorValue struct {
// 	color string
// 	value int
// }

func DayFive(c *gin.Context) {
	fileContent, err := FileCheck(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	lines, err := util.ReadFileContent(fileContent)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error reading file content",
		})
		return
	}
	fmt.Println("lines:", lines)

	c.JSON(200, gin.H{
		"result": lines,
	})
}
