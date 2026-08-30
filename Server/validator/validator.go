package validator

import "regexp"

var url *regexp.Regexp

func IsUrl() *regexp.Regexp {

	if url != nil {
		return url
	}

	pattern := `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\.[a-zA-Z]{2,}$`
	url = regexp.MustCompile(pattern)
	return url
}
