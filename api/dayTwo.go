package api

import (
	"fmt"
	"go-advent/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// type MapColorValue struct {
// 	color string
// 	value int
// }

func DayTwo(c *gin.Context) {
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

	results := []int{}
	totalResults := 0
	for _, line := range lines {
		line = strings.TrimSpace(line[7:])
		lineSplit := strings.Split(line, ";")

		//PART 1
		// result := parseLine(lineSplit)
		// if !result {
		// 	fmt.Println("Entrou No IF:", i+1)
		// 	results = append(results, i+1)
		// }

		//PART 2

		results = append(results, parseLine(lineSplit))

	}
	fmt.Println("results:", results)
	// fmt.Println("x:", x)
	for _, result := range results {
		totalResults += result
	}
	c.JSON(200, gin.H{
		"result": totalResults,
	})
}

func parseLine(arrLine []string) int {
	var arrayObj []map[string]int

	re := regexp.MustCompile(`(?i)(blue|green|red)|(\d+)`)

	for _, line := range arrLine {
		lineInArr := re.FindAllString(line, -1)
		obj := make(map[string]int)
		obj["red"] = 0
		obj["green"] = 0
		obj["blue"] = 0

		for i := 0; i < len(lineInArr); i += 2 {
			number, err := strconv.Atoi(lineInArr[i])
			if err != nil {
				fmt.Println("Error converting string to int")
				continue
			}

			color := lineInArr[i+1]
			obj[color] = number
		}

		arrayObj = append(arrayObj, obj)
	}

	//PART 1
	//Finding values on arrayObj
	// for i := range arrayObj {
	// 	if arrayObj[i]["red"] > 12 || arrayObj[i]["green"] > 13 || arrayObj[i]["blue"] > 14 {
	// 		fmt.Println("arrayObj:", arrayObj)
	// 		return true

	// 	}
	// }
	// return false

	//Part 2
	fmt.Println("arrayObj:", arrayObj)
	var fValues = make(map[string]int)
	for i := range arrayObj {
		fmt.Println("arrayObj[i]:", arrayObj[i])
		for color, value := range arrayObj[i] {
			if fValues[color] == 0 || value > fValues[color] {
				fValues[color] = value
			}
		}
	}

	multPlies := fValues["red"] * fValues["green"] * fValues["blue"]
	return multPlies

}
