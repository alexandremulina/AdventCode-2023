package util

import (
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

	re := regexp.MustCompile("(?i)one|two|three|four|five|six|seven|eight|nine")
	result := re.ReplaceAllStringFunc(line, func(match string) string {
		return numbersString[strings.ToLower(match)]
	})

	return result
}
