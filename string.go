package helper

import (
	"regexp"
	"strings"
)

//StripTags - Strip HTML and PHP tags from a string
func StripTags(s string) string {
	reg, _ := regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllStringFunc(s, strings.ToLower)
	//remove style
	reg, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	s = reg.ReplaceAllString(s, "")
	//remove script
	reg, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	s = reg.ReplaceAllString(s, "")

	reg, _ = regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllString(s, "\n")

	reg, _ = regexp.Compile(`\s{2,}`)
	s = reg.ReplaceAllString(s, "\n")

	return strings.TrimSpace(s)
}

// Substr - Return part of a string
func Substr(s string, start int, length ...int) string {
	if len(length) == 0 {
		return s[start:]
	}

	l := length[0]
	if l < 0 {
		end := len(s) + l
		if end < 0 {
			end = 0
		}
		return s[start:end]
	}
	end := start + l
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

// MbSubstr - Get part of string
func MbSubstr(s string, start int, length ...int) string {
	runes := []rune(s)
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(runes) + l
			if end < 0 {
				end = 0
			}
			return string(runes[start:end])
		}
		end := start + l
		if end > len(runes) {
			end = len(runes)
		}
		return string(runes[start:end])
	}
	return string(runes[start:])
}
