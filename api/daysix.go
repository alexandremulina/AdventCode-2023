package api

import (
	"fmt"
	"go-advent/util"
	_ "net/http/pprof"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DaySix(c *gin.Context) {
	fileContent, err := FileCheck(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	regexTime := regexp.MustCompile(`(\d+)`)
	regexDistance := regexp.MustCompile(`(\d+)`)

	lines, err := util.ReadFileContent(fileContent)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error reading file content",
		})
		return
	}

	// fmt.Println("time:", regexTime.FindAllString(
	// 	lines[0], -1), "distance:", regexDistance.FindAllString(lines[1], -1))
	//empty array
	emptyArray := []int{}
	emptyArray2 := []int{}
	timeArray := regexTime.FindAllString(lines[0], -1)
	distanceArray := regexDistance.FindAllString(lines[1], -1)
	// fmt.Println("timeArray:", strings.Join(timeArray, ""), "distanceArray:", strings.Join(distanceArray, ""))
	// Part two
	newTime, _ := strconv.Atoi(strings.Join(timeArray, ""))
	newDistance, _ := strconv.Atoi(strings.Join(distanceArray, ""))
	emptyArray2 = append(emptyArray2, boatMilimeterPartTwo(newTime, newDistance))
	multplieCounter2 := 1
	for _, empty2 := range emptyArray2 {
		multplieCounter2 *= empty2
	}
	// fmt.Println("multplieCounter2:", multplieCounter2)
	//Part One
	for index, time := range timeArray {
		//Covert to numbers
		time, _ := strconv.Atoi(time)
		distance, _ := strconv.Atoi(distanceArray[index])
		emptyArray = append(emptyArray, boatMilimeter(time, distance))

	}
	fmt.Println("emptyArray:", emptyArray)
	multplieCounter := 1
	for _, empty := range emptyArray {
		multplieCounter *= empty
	}
	fmt.Println("multplieCounter:", multplieCounter)

	c.JSON(200, gin.H{
		"lines":   lines,
		"message": "OK",
		"status":  200,
	})

}

func boatMilimeter(time int, distance int) int {
	counter := 0
	// fmt.Println("time:", time, "distance:", distance)
	for i := 1; i <= time; i++ {
		holder := i
		distanceAfterHold := holder * (time - holder) // ---> holder*time - holder*holder = distanceAfterHold
		// fmt.Println("distanceAfterHold:", distanceAfterHold)
		if distanceAfterHold > distance {
			counter++
			// fmt.Println("counter:", counter)
		}

	}
	return counter
}

// Part 2
func boatMilimeterPartTwo(time int, distance int) int {
	// counter := 0
	var firstHolder int
	var lasttHolder int
	for i := 1; i <= time; i++ {
		holder := i
		distanceAfterHold := holder * (time - holder)
		if distanceAfterHold > distance {
			firstHolder = holder
			break
		}

	}
	for j := time; j >= 0; j-- {
		holderTwo := j
		distanceAfterHold := holderTwo * (time - holderTwo)
		if distanceAfterHold > distance {
			lasttHolder = holderTwo
			break
		}
	}
	fmt.Println("firstHolder:", firstHolder, "lasttHolder:", lasttHolder)

	fmt.Println("FinalCounter:", lasttHolder-firstHolder+1)
	finalCounter := lasttHolder - firstHolder + 1
	return finalCounter

}
