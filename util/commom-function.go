package util

import (
	"fmt"
	"regexp"
	"strings"
)

func ConcatNumbers(numbers []string) string {
	var stringNumbers string
	concatString := strings.Join(numbers, "")
	if len(concatString) > 1 {
		stringNumbers = concatString[0:1] + concatString[len(concatString)-1:]
	} else if len(concatString) == 1 {
		// fmt.Println("concatString====1:", concatString)
		stringNumbers = concatString + concatString
		// fmt.Println("stringNumbers====1:", stringNumbers)
	}

	return stringNumbers
}

// replaceStringNameToNumber
func ReplaceStringNameNumber(line string) string {
	numbersString := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	// numbersSlice := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re := regexp.MustCompile("(?i)one|two|three|four|five|six|seven|eight|nine")
	// fmt.Println("FindAllString", re.FindAllString(line, -1))
	result := re.ReplaceAllStringFunc(line, func(match string) string {
		return numbersString[strings.ToLower(match)]
	})
	// fmt.Println("ReplaceAllStringFunc:", findNumberByRegex(result))
	// fmt.Println("ReverseMatchString:", re.ReplaceAllStringFunc(ReverseMatchString(line), func(match string) string {
	// 	return numbersString[strings.ToLower(match)]
	// }))
	findNumberRege := strings.Join(findNumberByRegex(result), "")
	// fmt.Println(("result:"), result)

	// fmt.Println("findNumberByRegex", findNumberByRegex(findNumberRege[0:1]))
	reverseStringMatch := re.ReplaceAllStringFunc(ReverseMatchString(line), func(match string) string {
		return numbersString[strings.ToLower(match)]
	})
	fmt.Println("FIRST NORMAL:", findNumberRege[0:1])
	fmt.Println("reverseStringMatch:", reverseStringMatch)

	return findNumberRege[0:1] + reverseStringMatch
}

func ReverseMatchString(line string) string {
	reversedLine := reverseString(line)
	re := regexp.MustCompile("(?i)eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[1-9]+")

	reversedMatches := re.FindAllString(reversedLine, -1)

	for i := range reversedMatches {
		reversedMatches[i] = reverseString(reversedMatches[i])
	}

	// firstMatch := ""
	// if len(matches) > 0 {
	// 	firstMatch = matches[0]
	// }

	firstReversedMatch := ""
	if len(reversedMatches) > 0 {
		firstReversedMatch = reversedMatches[0]
	}

	// fmt.Println("First Match (Normal):", firstMatch)
	// fmt.Println("First Match (Reversed):", firstReversedMatch)

	return firstReversedMatch
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findNumberByRegex(result string) []string {
	re := regexp.MustCompile("[1-9]+")
	return re.FindAllString(result, -1)
}
