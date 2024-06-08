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

func FormatHTML(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		prev, next := "", ""
		if i == 0 {
			prev = " "
		} else {
			prev = string(s[i-1])
		}
		if i == len(s)-1 {
			next = " "
		} else {
			next = string(s[i+1])
		}

		if s[i] == ' ' {
			if prev == " " && next != " " {
				str += "&nbsp;"
			} else if prev == " " && next == " " {
				str += "&nbsp;"
			} else {
				str += " "
			}
		} else {
			str += string(s[i])
		}
	}
	return str
}
