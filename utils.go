package goku

import "os"

// Exists - this function is used to check if file exists and return bool value
func Exists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
