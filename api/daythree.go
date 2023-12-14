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
	// fmt.Println("coord:", coord, "numbersRange:", symbolRanges)
	for _, symbolRange := range symbolRanges {

		if (symbolRange[0] >= coord[0] && symbolRange[0] <= coord[1]) ||
			(symbolRange[1] >= coord[0] && symbolRange[1] <= coord[1]) {
			// fmt.Println("symbolRange:", symbolRange[0], symbolRange[1])
			// fmt.Println("coord:", coord[0], coord[1])
			// fmt.Println("return True")
			return true
		}
	}

	return false
}

//Part 2

//Check for Numbers Adjancts to Symbol *

func DayThreePart2(c *gin.Context) {
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
	totalGears := 0
	re := regexp.MustCompile(`(?i)\*`)
	// numbersRegex := regexp.MustCompile(`(\d+)`)
	for i := 0; i < len(lines); i++ {
		lineSymbolsIndexs := re.FindAllStringIndex(lines[i], -1)
		lineSymbols := re.FindAllString(lines[i], -1)
		// fmt.Println("lineSymbolsIndexs:", lineSymbolsIndexs)
		// umbersIndexes := numbersRegex.FindAllStringIndex(lines[i], -1)
		// fmt.Println("numbers:", numbersRegex.FindAllString(lines[i], -1))
		// fmt.Println("numbersIndexes:", umbersIndexes)
		for index, lineSymbol := range lineSymbols {
			// fmt.Println("lineSymbolsIndexs", lineSymbolsIndexs[index])
			numbers := checkAdjacentNumbersFromSymbols(lineSymbol, lineSymbolsIndexs[index], i, lines)
			fmt.Println("numbers:", numbers, "index:", i)
			intNumbers := convertStringArrayToInt(numbers)

			if len(intNumbers) > 1 {
				multiplied := intNumbers[0] * intNumbers[1]
				totalGears += multiplied

			}
		}

	}
	fmt.Println("totalGears:", totalGears)

	c.JSON(200, gin.H{
		"result": lines,
		"status": 200,
	})
}

func checkAdjacentNumbersFromSymbols(stringSymbol string, coord []int, indexLine int, lines []string) []string {
	arrString := []string{}
	re := regexp.MustCompile(`(\d+)`)

	// Before Line
	if indexLine > 0 {
		beforeNumbers := re.FindAllStringIndex(lines[indexLine-1], -1)
		if len(beforeNumbers) > 0 && beforeNumbers != nil {
			for _, coords := range rangeIntersect(coord, beforeNumbers) {
				arrString = append(arrString, lines[indexLine-1][coords[0]:coords[1]])
			}
		}
	}

	// Current Line
	if indexLine >= 0 && indexLine < len(lines) {
		currentNumbers := re.FindAllStringIndex(lines[indexLine], -1)
		if len(currentNumbers) > 0 && currentNumbers != nil {
			for _, coords := range rangeIntersect(coord, currentNumbers) {
				arrString = append(arrString, lines[indexLine][coords[0]:coords[1]])
			}
		}
	}

	// After Line
	if indexLine < len(lines)-1 {
		afterNumbers := re.FindAllStringIndex(lines[indexLine+1], -1)
		if len(afterNumbers) > 0 && afterNumbers != nil {
			for _, coords := range rangeIntersect(coord, afterNumbers) {
				arrString = append(arrString, lines[indexLine+1][coords[0]:coords[1]])
			}
		}
	}

	return arrString
}

func rangeIntersect(coord []int, numbersRanges [][]int) [][]int {
	var result [][]int
	for _, numbersRange := range numbersRanges {
		if (numbersRange[0] >= coord[0] && numbersRange[0] <= coord[1]) ||
			(numbersRange[1] >= coord[0] && numbersRange[1] <= coord[1]) ||
			(numbersRange[0] <= coord[0] && numbersRange[1] >= coord[1]) {
			result = append(result, numbersRange)
		}
	}
	return result
}

func convertStringArrayToInt(strNumbers []string) []int {
	intNumbers := make([]int, len(strNumbers))
	for i, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			intNumbers[i] = num
		}
	}
	return intNumbers
}
