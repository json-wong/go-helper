package helper

import (
	"net/url"
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

// ParseURL parse_url()
// Parse a URL and return its components
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment
func ParseURL(str string, component int) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var components = make(map[string]string)
	if (component & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		components["port"] = u.Port()
	}
	if (component & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		components["path"] = u.Path
	}
	if (component & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}
