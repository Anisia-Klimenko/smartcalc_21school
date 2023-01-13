package file

import (
	"os"
)

// Update adds new line with content to file fileName
func Update(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	f.WriteString(content + "\n")
}

// Rewrite truncates file fileName when opened and writes new line with
// content string
func Rewrite(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	f.WriteString(content + "\n")
}

// Clear clears file fileName
func Clear(fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0600)
	if err != nil {
		//log.Println(err)
		return
	}

	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
}

// Content returns content of file fileName
func Content(fileName string) string {
	result, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(result)
}
