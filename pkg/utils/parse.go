package utils

import (
	"regexp"
	"strings"
)

func ParseEnv(src []byte) map[string]string {
	LINE := regexp.MustCompile(`(?m)(?:^|^)\s*(?:export\s+)?([\w.-]+)(?:\s*=\s*?|:\s+?)(\s*'(?:\\'|[^'])*'|\s*"(?:\\"|[^"])*"|\s*` + "`" + `(?:\\` + "`" + `|[^` + "`" + `])*` + "`" + `|[^#\r\n]+)?\s*(?:#.*)?(?:$|$)`)

	matches := LINE.FindAllSubmatch(src, -1)

	obj := make(map[string]string, 0)

	for _, match := range matches {
		key := string(match[1])
		value := string(match[2])
		value = strings.TrimSpace(value)

		maybeQuote := value[0]

		value = strings.Trim(value, `"`)

		if maybeQuote == '"' {
			value = strings.ReplaceAll(value, "\\n", "\n")
			value = strings.ReplaceAll(value, "\\r", "\r")
		}

		obj[key] = value
	}
	return obj
}
