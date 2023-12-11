package api

import (
	"fmt"
	"go-advent/util"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type MapColorValue struct {
// 	color string
// 	value int
// }

func DayThree(c *gin.Context) {
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

	re := regexp.MustCompile(`(\d+)`)
	stringArray := []string{}
	sum := 0
	for i := 0; i < len(lines); i++ {
		// fmt.Println("line Numbers:", re.FindAllString(lines[i], -1))
		// fmt.Println("line Numbers Index:", i, re.FindAllStringIndex(lines[i], -1))
		lineNumbersIndexs := re.FindAllStringIndex(lines[i], -1)
		// fmt.Println("Symbol:", i, symbolRegex.FindAllStringIndex(lines[i], -1))
		lineNumbers := re.FindAllString(lines[i], -1)
		for index, lineNumber := range lineNumbers {
			// fmt.Println("lineNumber:", lineNumber, lineNumbersIndexs[index])

			checkAdjacentSymbols(lineNumber, lineNumbersIndexs[index], i, lines)
			fmt.Println("checkAdjacentSymbols:", checkAdjacentSymbols(lineNumber, lineNumbersIndexs[index], i, lines), "number:", lineNumber)
			if checkAdjacentSymbols(lineNumber, lineNumbersIndexs[index], i, lines) {
				stringArray = append(stringArray, lineNumber)
			}

		}

	}
	fmt.Println("stringArray:", stringArray)
	for _, number := range stringArray {
		numberInt, _ := strconv.Atoi(number)
		sum += numberInt
	}
	fmt.Println("sum:", sum)

	c.JSON(200, gin.H{
		"result": lines,

		"status": 200,
	})
}

func checkAdjacentSymbols(stringNumber string, coord []int, indexLine int, lines []string) bool {
	symbolRegex := regexp.MustCompile(`[^0-9.]`)
	fmt.Println("stringNumber:", stringNumber, "coord:", coord, "indexLine:", indexLine)

	// Before Line
	if indexLine > 0 {
		beforeSymbols := symbolRegex.FindAllStringIndex(lines[indexLine-1], -1)
		// fmt.Println("Symbol (Before):", beforeSymbols)

		if rangesIntersect(coord, beforeSymbols) {
			fmt.Println(rangesIntersect(coord, beforeSymbols))
			return true
		}
	}

	//Current Line
	if indexLine >= 0 && indexLine < len(lines) {
		currentSymbols := symbolRegex.FindAllStringIndex(lines[indexLine], -1)
		// fmt.Println("Symbol (Current):", currentSymbols)

		if rangesIntersect(coord, currentSymbols) {
			fmt.Println(rangesIntersect(coord, currentSymbols))
			return true
		}
	}

	// After Line
	if indexLine < len(lines)-1 {
		afterSymbols := symbolRegex.FindAllStringIndex(lines[indexLine+1], -1)
		// fmt.Println("Symbol (After):", afterSymbols)

		if rangesIntersect(coord, afterSymbols) {
			fmt.Println(rangesIntersect(coord, afterSymbols))
			return true
		}
	}

	return false
}

func rangesIntersect(coord []int, symbolRanges [][]int) bool {
	fmt.Println("coord:", coord, "symbolRanges:", symbolRanges)
	for _, symbolRange := range symbolRanges {

		if (symbolRange[0] >= coord[0] && symbolRange[0] <= coord[1]) ||
			(symbolRange[1] >= coord[0] && symbolRange[1] <= coord[1]) {
			fmt.Println("symbolRange:", symbolRange[0], symbolRange[1])
			fmt.Println("coord:", coord[0], coord[1])
			fmt.Println("return True")
			return true
		}
	}

	return false
}
