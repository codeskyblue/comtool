package comtool

import (
	"os"
	"path/filepath"
)

// Return program dir
func SelfDir() string {
	return filepath.Dir(os.Args[0])
}

// Check if file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
