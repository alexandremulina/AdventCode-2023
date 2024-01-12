package api

import (
	"fmt"
	"go-advent/util"
	"math"
	"regexp"
	"strconv"
	"strings"

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

	regexSeeds := regexp.MustCompile(`seeds:\s*((\d+\s*)+)`)
	regexSeedToSoil := regexp.MustCompile(`seed-to-soil map:\s*((\d+\s*)+)`)
	regexSoilToFertilizer := regexp.MustCompile(`soil-to-fertilizer map:\s*((\d+\s*)+)`)
	regexFertilizerToWater := regexp.MustCompile(`fertilizer-to-water map:\s*((\d+\s*)+)`)
	regexWaterToLight := regexp.MustCompile(`water-to-light map:\s*((\d+\s*)+)`)
	regexLightToTemperature := regexp.MustCompile(`light-to-temperature map:\s*((\d+\s*)+)`)
	regexTemperatureToHumidity := regexp.MustCompile(`temperature-to-humidity map:\s*((\d+\s*)+)`)
	regexHumidityToLocation := regexp.MustCompile(`humidity-to-location map:\s*((\d+\s*)+)`)

	lines, err := util.ReadFileContent(fileContent)
	fmt.Println("lines before join:", lines)
	joinLines := joinNonEmptyStrings(lines, " ")
	fmt.Println("lines:", joinLines)
	allSeeds := removeEmptyStrings(strings.Split(regexSeeds.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allSeeds:", allSeeds)
	//regexSeedToSoil
	allSeedToSoil := removeEmptyStrings(strings.Split(regexSeedToSoil.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allSeedToSocil:", allSeedToSoil)
	// regexSoilToFertilizer
	allSoilToFertilizer := removeEmptyStrings(strings.Split(regexSoilToFertilizer.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allSoilToFertilizer:", allSoilToFertilizer)
	// regexFertilizerToWater
	allFertilizerToWater := removeEmptyStrings(strings.Split(regexFertilizerToWater.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allFertilizerToWater:", allFertilizerToWater)
	// regexWaterToLight
	allWaterToLight := removeEmptyStrings(strings.Split(regexWaterToLight.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allWaterToLight:", allWaterToLight)
	// regexLightToTemperature
	allLightToTemperature := removeEmptyStrings(strings.Split(regexLightToTemperature.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allLightToTemperature:", allLightToTemperature)
	// regexTemperatureToHumidity
	allTemperatureToHumidity := removeEmptyStrings(strings.Split(regexTemperatureToHumidity.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allTemperatureToHumidity:", allTemperatureToHumidity)
	// regexHumidityToLocation
	allHumidityToLocation := removeEmptyStrings(strings.Split(regexHumidityToLocation.FindAllStringSubmatch(joinLines, -1)[0][1], " "))
	fmt.Println("allHumidityToLocation:", allHumidityToLocation)

	// allSeeds := re.FindAllString(lines[0], -1)[1:]
	// fmt.Println("allSeeds:", allSeeds)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error reading file content",
		})
		return
	}
	// fmt.Println("lines:", lines)

	//PartOne
	// for i, line := range lines {
	// 	if lines[i] == "" {
	// 		continue
	// 	}
	// 	fmt.Println("line:", line)
	// 	// seedsMatches := regexSeeds.FindStringSubmatch(line)
	// 	// // seedToSoilMatches := regexSeedToSoil.FindStringSubmatch(line)
	// 	// // soilToFertilizerMatches := regexSoilToFertilizer.FindStringSubmatch(line)
	// 	// // fertilizerToWaterMatches := regexFertilizerToWater.FindStringSubmatch(line)
	// 	// // waterToLightMatches := regexWaterToLight.FindStringSubmatch(line)
	// 	// // lightToTemperatureMatches := regexLightToTemperature.FindStringSubmatch(line)
	// 	// // temperatureToHumidityMatches := regexTemperatureToHumidity.FindStringSubmatch(line)
	// 	// // humidityToLocationMatches := regexHumidityToLocation.FindStringSubmatch(line)
	// 	// seeds := extractIntArray(seedsMatches[1])
	// 	// fmt.Println("seeds:", seeds)
	// 	// // seedToSoil := extractIntArray(seedToSoilMatches[1])
	// 	// // fmt.Println("seedToSoil:", seedToSoil)

	// }
	seedsArrInt, err := stringsToInts(allSeeds)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error converting strings to ints",
		})
		return
	}
	fmt.Println("allSeedToSoil:", allSeedToSoil, "seedsArrInt :", seedsArrInt)
	seedSoil := calculateRange(allSeedToSoil, seedsArrInt)
	fmt.Println("newSeedAfterSoil:", seedSoil)
	fmt.Println("allSoilToFertilizer:", allSoilToFertilizer, "seedSoil", seedSoil)
	seedToFertilizer := calculateRange(allSoilToFertilizer, seedSoil)
	fmt.Println("newSeedToFertilizer:", seedToFertilizer)
	fmt.Println("allFertilizerToWater:", allFertilizerToWater, "seedToFertilizer:", seedToFertilizer)
	seedToWater := calculateRange(allFertilizerToWater, seedToFertilizer)
	fmt.Println("newSeedToWater:", seedToWater)
	fmt.Println("allWaterToLight:", allWaterToLight, "seedToWater:", seedToWater)
	seedToLight := calculateRange(allWaterToLight, seedToWater)
	fmt.Println("newSeedToLight:", seedToLight)
	fmt.Println("allLightToTemperature:", allLightToTemperature, "seedToLight:", seedToLight)
	seedToTemperature := calculateRange(allLightToTemperature, seedToLight)
	fmt.Println("newSeedToTemperature:", seedToTemperature)
	fmt.Println("allTemperatureToHumidity:", allTemperatureToHumidity, "seedToTemperature:", seedToTemperature)
	seedToHumidity := calculateRange(allTemperatureToHumidity, seedToTemperature)
	fmt.Println("newSeedToHumidity:", seedToHumidity)
	fmt.Println("allHumidityToLocation:", allHumidityToLocation, "seedToHumidity:", seedToHumidity)
	seedToLocation := calculateRange(allHumidityToLocation, seedToHumidity)
	fmt.Println("newSeedToLocation:", seedToLocation)
	smallestNumber := findSmallestNumber(seedToLocation)
	fmt.Println("smallestNumber:", smallestNumber)
	c.JSON(200, gin.H{
		"result": gin.H{
			"allSeeds":            allSeeds,
			"allSeedToSoil":       allSeedToSoil,
			"allSoilToFertilizer": allSoilToFertilizer,
			// "allFertilizerToWater":     allFertilizerToWater,
			// "allWaterToLight":          allWaterToLight,
			// "allLightToTemperature":    allLightToTemperature,
			// "allTemperatureToHumidity": allTemperatureToHumidity,
			// "allHumidityToLocation":    allHumidityToLocation,
		},
	})
}

func removeEmptyStrings(input []string) []string {
	var result []string

	for _, item := range input {
		if item != "" {
			result = append(result, item)
		}
	}

	return result
}

func joinNonEmptyStrings(input []string, separator string) string {
	var nonEmptyStrings []string

	for _, item := range input {
		if item != "" {
			nonEmptyStrings = append(nonEmptyStrings, item)
		}
	}

	return strings.Join(nonEmptyStrings, separator)
}

func calculateRange(array []string, seeds []int) []int {
	//Create a copy because its not possible to update the original array
	newSeedUpdated := make([]int, len(seeds))
	copy(newSeedUpdated, seeds)
	// var destinationArr, sourceArr []string
	var diffBetweenDestinationAndSource int

	for i := 0; i < len(array); i += 3 {
		destination, err := strconv.Atoi(array[i])
		if err != nil {
			continue
		}
		source, err := strconv.Atoi(array[i+1])
		if err != nil {
			continue
		}
		rangeNumber, err := strconv.Atoi(array[i+2])
		if err != nil {
			continue
		}
		diffBetweenDestinationAndSource = findAbsuleDiff(destination, source)
		fmt.Println("destination:", destination, "source:", source, "range:", rangeNumber, "diffBetweenDestinationAndSource:", diffBetweenDestinationAndSource)
		// fmt.Println("destination:", destination) // Destination = Soil
		// fmt.Println("source:", source)           // Source = Seed
		// fmt.Println("range:", rangeNumber)
		// fmt.Println("diffBetweenDestinationAndSource:", findAbsuleDiff(destination, source))
		// fmt.Println("seedsParamsOfCalculateRange:", seeds)
		for index, seed := range seeds {
			// maxDestination := destination + rangeNumber - 1
			// minDestination := destination
			maxSource := source + rangeNumber - 1
			minSource := source
			fmt.Println("maxSource:", maxSource, "minSource:", minSource)
			fmt.Println("seed to Find on Source Range", seed)
			if seed >= minSource && seed <= maxSource {
				if destination > source {
					newSeedUpdated[index] = seed + diffBetweenDestinationAndSource
				} else if destination < source {
					newSeedUpdated[index] = seed - diffBetweenDestinationAndSource
				}
			}
			fmt.Println("seedsUpdated:", newSeedUpdated)
		}
	}

	// fmt.Println("newSeed:", seeds)
	return newSeedUpdated
}

func findAbsuleDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func stringsToInts(stringArray []string) ([]int, error) {
	var intArray []int

	for _, str := range stringArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray = append(intArray, num)
	}

	return intArray, nil
}

func findSmallestNumber(array []int) int {
	smallestNumber := array[0]
	for _, num := range array {
		if num < smallestNumber {
			smallestNumber = num
		}
	}
	return smallestNumber
}
