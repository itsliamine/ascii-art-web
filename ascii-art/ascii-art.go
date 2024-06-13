package asciiart

import (
	"errors"
	"fmt"
	"strings"
)

func GetAscii(input, style string) ([]string, error) {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}, errors.New(" ")
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		getW, err := GetWord(word, bannerFile)
		if err != nil {
			return []string{}, errors.New("error 500")
		}
		lines = append(lines, getW...)
	}

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], " ", "&nbsp;")
	}

	return lines, nil
}
