package api

import (
	"go-advent/util"
	"io"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileContent struct {
	Content string `form:"content" binding:"required"`
}

// healthCheck is a simple health check endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
		"status":  200,
	})
}

// filecheck function
func FileCheck(c *gin.Context) (io.Reader, error) {
	file, err := c.FormFile("content")
	if err != nil {
		return nil, err
	}

	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	return fileContent, nil
}

func DayOne(c *gin.Context) {
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
	re := regexp.MustCompile("[1-9]+")
	errorList := []int{}
	total := 0
	var totalError int
	for i, line := range lines {
		newLine := util.ReplaceStringNameNumber(line)

		numbers := re.FindAllString(newLine, -1)

		num, err := strconv.Atoi(util.ConcatNumbers(numbers))
		if err != nil {
			errorList = append(errorList, i)
			totalError++
			continue

		}
		total += num
		// for _, char := range line {
		// 	fmt.Printf("%c", char) // Convert rune to string using %c format specifier
		// }
	}

	// fmt.Println("Total:", total)

	c.JSON(200, gin.H{
		"result":     lines,
		"total":      total,
		"errorCount": errorList,
		"status":     200,
	})
}
