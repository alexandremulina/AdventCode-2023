package api

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileContent struct {
	Content string `form:"content" binding:"required"`
}

func DayOne(c *gin.Context) {
	file, err := c.FormFile("content")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Missing or invalid 'content' field",
		})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error opening file",
		})
		return
	}
	defer fileContent.Close()

	content, err := ioutil.ReadAll(fileContent)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error reading file content",
		})
		return
	}

	contentString := string(content)

	var lines []string
	for _, line := range strings.Split(contentString, "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	c.JSON(200, gin.H{
		"message": lines,
		"status":  200,
	})
}
