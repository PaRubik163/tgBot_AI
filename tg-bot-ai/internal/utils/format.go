package utils
//Здесь форматируются строки
import (
	"strings"
)

func EditResult(result string) string {
	result = strings.ReplaceAll(result, `\n`, "\n")
	result = strings.ReplaceAll(result, "*", "")
	result = strings.ReplaceAll(result, "#", "")

	return result
}