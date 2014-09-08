package env

import "fmt"
import "os"

// Get env var `name` or return an error if it's missing.
func Get(name string) (string, error) {
	if s := os.Getenv(name); s == "" {
		return "", fmt.Errorf("environment variable %s missing", name)
	} else {
		return s, nil
	}
}
