package util

import (
	"os"
)

// MissingEnvironList returns a list of missing environ variables
func MissingEnvironList(keys []string) []string {

	missing := []string{}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}
