package utils
//Здесь форматируются строки
import (
	"strings"
	"regexp"
)

func EditResult(result string) string {
	result = strings.ReplaceAll(result, `\n`, "\n")
	result = strings.ReplaceAll(result, "*", "")
	result = strings.ReplaceAll(result, "#", "")

	return result
}

func RemoveThinkBlockFromAnswer(resutl string) string {
	re := regexp.MustCompile(`(?s)<think>.*?</think>`)

	return re.ReplaceAllString(resutl, "")
}