package api

import (
	"fmt"
	"go-advent/util"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// type MapColorValue struct {
// 	color string
// 	value int
// }

func DayFour(c *gin.Context) {
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
	cardCount := make(map[int]int)

	for i, line := range lines {
		line2 := strings.TrimSpace(line[8:])
		fmt.Printf("Card: %d %s\n", i+1, line2)

		symbolAndNumbers := regexp.MustCompile(`(\d+)|\|`)
		symbolNumbers := symbolAndNumbers.FindAllString(line2, -1)

		numbersBefore := make([]string, 0)
		numbersAfter := make([]string, 0)
		foundSymbol := false

		for _, match := range symbolNumbers {
			if match == "|" {
				foundSymbol = true
				numbersAfter = symbolNumbers[len(numbersBefore)+1:]
				break
			}
			numbersBefore = append(numbersBefore, match)
		}

		if !foundSymbol {
			numbersAfter = symbolNumbers
		}

		fmt.Println("Numbers before |:", numbersBefore)
		fmt.Println("Numbers after |:", numbersAfter)
		matches := calculateMatches(numbersBefore, numbersAfter)

		cardCount[i+1]++

		fmt.Println("lineMatches:", matches)
		//Part two
		for j := 1; j <= matches; j++ {
			cardCount[i+j+1] += cardCount[i+1]
		}
		fmt.Println("cardCount:", cardCount)
	}

	totalCount := 0
	for _, count := range cardCount {
		totalCount += count
	}
	fmt.Println("totalCount:", totalCount)
	c.JSON(200, gin.H{
		"result": lines,
	})
}

// func calculatePoints(winningNumbers, myNumbers []string) int {
// 	winningNumberSet := make(map[string]struct{})
// 	for _, num := range winningNumbers {
// 		winningNumberSet[num] = struct{}{}
// 	}

// 	points := 0
// 	matches := 0

// 	for _, num := range myNumbers {
// 		if _, ok := winningNumberSet[num]; ok {
// 			matches++
// 			if matches == 1 {
// 				points++ // 1 point for the first match
// 			} else {

// 				points *= 2
// 			}
// 		}
// 	}

// 	return points
// }

// Part One
func calculateMatches(winningNumbers, myNumbers []string) int {
	winningNumberSet := make(map[string]struct{})
	for _, num := range winningNumbers {
		winningNumberSet[num] = struct{}{}
	}

	matches := 0
	for _, num := range myNumbers {
		if _, ok := winningNumberSet[num]; ok {
			matches++
		}
	}

	return matches
}
