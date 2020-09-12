package plugin

import (
	"fmt"
)

func BuildVersion(version, commit string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s+%s", result, commit)
	}
	return result
}
