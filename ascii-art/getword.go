package asciiart

import (
	"errors"
	"strings"
)

func GetWord(input string, bannerFile string) ([]string, error) {
	lines := make([]string, 8)
	content := FileOpen(bannerFile)

	for _, char := range input {
		if !(char >= 31 && char <= 126) {
			return []string{}, errors.New("error 500")
		}
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}

	return lines, nil
}
