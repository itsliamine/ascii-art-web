package asciiart

import (
	"fmt"
	"strings"
)

func GetAscii(input, style string) []string {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		lines = append(lines, GetWord(word, bannerFile)...)
	}

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], " ", "&nbsp;")
	}

	return lines
}
