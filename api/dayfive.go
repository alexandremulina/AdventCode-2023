package api

import (
	"fmt"
	"go-advent/util"
	"math"
	_ "net/http/pprof"
	"regexp"
	"strconv"
	"strings"
	"sync"

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
	seedsArrInt, err := stringsToInts(allSeeds)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error converting strings to ints",
		})
		return
	}
	//Part Two
	// fmt.Println("transFormSeedsIntoRangeArray:", transFormSeedsIntoRangeArray(seedsArrInt))
	seedsArrInt2 := transFormSeedsIntoRangeArray(seedsArrInt)
	//

	fmt.Println("allSeedToSoil:", allSeedToSoil, "seedsArrInt :", seedsArrInt2)
	seedSoil := calculateRange(allSeedToSoil, seedsArrInt2)
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

// Part one
func calculateRange(array []string, seeds []int) []int {
	//Create a copy because its not possible to update the original array
	newSeedUpdated := make([]int, len(seeds))
	copy(newSeedUpdated, seeds)
	// var destinationArr, sourceArr []string
	var diffBetweenDestinationAndSource int

	for i := 0; i < len(array); i += 3 {
		destination, _ := strconv.Atoi(array[i])

		source, _ := strconv.Atoi(array[i+1])

		rangeNumber, _ := strconv.Atoi(array[i+2])

		diffBetweenDestinationAndSource = findAbsuleDiff(destination, source)
		fmt.Println("destination:", destination, "source:", source, "range:", rangeNumber, "diffBetweenDestinationAndSource:", diffBetweenDestinationAndSource)

		for index, seed := range seeds {
			// maxDestination := destination + rangeNumber - 1
			// minDestination := destination
			maxSource := source + rangeNumber - 1
			minSource := source
			// fmt.Println("maxSource:", maxSource, "minSource:", minSource)
			// fmt.Println("seed to Find on Source Range", seed)
			if seed >= minSource && seed <= maxSource {
				if destination > source {
					newSeedUpdated[index] = seed + diffBetweenDestinationAndSource
				} else if destination < source {
					newSeedUpdated[index] = seed - diffBetweenDestinationAndSource
				}
			}
			// fmt.Println("seedsUpdated:", newSeedUpdated)
		}
	}

	// fmt.Println("newSeed:", seeds)
	return newSeedUpdated
}

// func calculateRange(array []string, seeds []int) []int {
// 	newSeedUpdated := make([]int, len(seeds))
// 	copy(newSeedUpdated, seeds)

// 	for i := 0; i < len(array); i += 3 {
// 		destination, _ := strconv.Atoi(array[i])
// 		source, _ := strconv.Atoi(array[i+1])
// 		rangeNumber, _ := strconv.Atoi(array[i+2])

// 		minSource := source
// 		maxSource := source + rangeNumber - 1

// 		// Find seeds within the range
// 		indicesToUpdate := findIndicesInRange(seeds, minSource, maxSource)

// 		for _, index := range indicesToUpdate {
// 			diffBetweenDestinationAndSource := findAbsuleDiff(destination, source)
// 			if destination > source {
// 				newSeedUpdated[index] = seeds[index] + diffBetweenDestinationAndSource
// 			} else if destination < source {
// 				newSeedUpdated[index] = seeds[index] - diffBetweenDestinationAndSource
// 			}
// 		}
// 	}

// 	return newSeedUpdated
// }

// func findIndicesInRange(seeds []int, min, max int) []int {
// 	fmt.Println("seeds:", seeds, "min:", min, "max:", max)
// 	var indices []int

// 	//
// 	for i := 0; i < len(seeds); i += 2 {
// 		fmt.Println("isNumberInRange:", findMinInRange(seeds[i], seeds[i+1], min, max))
// 		partTwoSeed := findMinInRange(seeds[i], seeds[i+1], min, max)
// 		if partTwoSeed != -1 {
// 			if partTwoSeed >= min && partTwoSeed <= max {
// 				indices = append(indices, i)
// 			}
// 		}

// 	}
// 	//

// 	// for i, seed := range seeds {
// 	// 	if seed >= min && seed <= max {
// 	// 		indices = append(indices, i)
// 	// 	}
// 	// }
// 	fmt.Println("indices:", indices)
// 	return indices
// }

// func findMinInRange(aMin, aLen, bRangeMin, bRangeMax int) int {

// 	aMax := aMin + aLen - 1

// 	if aMin > aMax || bRangeMin > bRangeMax {
// 		return -1 // Use -1 to indicate an error or no overlap
// 	}

// 	if aMin > bRangeMax || aMax < bRangeMin {
// 		return -1 // Use -1 to indicate an error or no overlap
// 	}

// 	min := int(math.Max(float64(aMin), float64(bRangeMin)))
// 	return min
// }

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

// func transFormSeedsIntoRangeArray(seeds []int) []int {
// 	fmt.Println("seeds:", seeds)
// 	totalLength := 0
// 	for i := 1; i < len(seeds); i += 2 {
// 		totalLength += seeds[i]
// 	}

// 	rangeArray := make([]int, 0, totalLength)

// 	for i := 0; i < len(seeds); i += 2 {
// 		originalSeed := seeds[i]
// 		rangeSeed := seeds[i+1]
// 		for j := 0; j < rangeSeed; j++ {
// 			rangeArray = append(rangeArray, originalSeed+j)
// 		}
// 	}

//		fmt.Println("len(rangeArray):", len(rangeArray))
//		return rangeArray
//	}
type SeedIterator struct {
	start, end int
	current    int
}

func NewSeedIterator(start, end int) *SeedIterator {
	return &SeedIterator{start: start, end: end, current: start - 1}
}

func (s *SeedIterator) Next() int {
	s.current++
	if s.current > s.end {
		return -1
	}
	return s.current
}

func seedRangeGenerator(start, end int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	iterator := NewSeedIterator(start, end)
	for {
		seed := iterator.Next()
		if seed == -1 {
			break
		}
		ch <- seed
	}
}

func transFormSeedsIntoRangeArray(seeds []int) []int {
	fmt.Println("seeds:", seeds)

	var rangeArray []int
	var wg sync.WaitGroup

	for i := 0; i < len(seeds); i += 2 {
		originalSeed, rangeSeed := seeds[i], seeds[i+1]

		seedGen := make(chan int)

		wg.Add(1)
		go seedRangeGenerator(originalSeed, originalSeed+rangeSeed-1, seedGen, &wg)

		go func() {
			for seed := range seedGen {
				rangeArray = append(rangeArray, seed)
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	// fmt.Println("len(rangeArray):", len(rangeArray))
	return rangeArray
}

// func transFormSeedsIntoRangeArray(seeds []int) []int {
// 	fmt.Println("seeds:", seeds)

// 	var rangeArray []int

// 	for i := 0; i < len(seeds); i += 2 {
// 		originalSeed, rangeSeed := seeds[i], seeds[i+1]

// 		start := originalSeed
// 		end := originalSeed + rangeSeed - 1

// 		for j := start; j <= end; j++ {
// 			rangeArray = append(rangeArray, j)
// 		}
// 	}

// 	// fmt.Println("len(rangeArray):", len(rangeArray))
// 	return rangeArray
// }
