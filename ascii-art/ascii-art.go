package asciiart

import (
	"fmt"
	"strings"
)

func GetAscii(input, style string) string {
	str := ""
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")
	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
		} else {
			lines = append(lines, GetWord(word, bannerFile)...)
		}
	}
	for _, line := range lines {
		str += line + "\n"
	}
	return str
}
