package file

import (
	"os"
)

// Update adds content as new line to file fileName
func Update(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content + "\n")
	return nil
}

// Rewrite truncates file fileName when opened and writes new line with
// content string
func Rewrite(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content + "\n")
	return nil
}

// Clear clears file fileName
func Clear(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
	return nil
}

// Content returns content of file fileName
func Content(fileName string) string {
	result, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(result)
}
