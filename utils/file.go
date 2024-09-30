package utils

import "os"

func CreateFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
