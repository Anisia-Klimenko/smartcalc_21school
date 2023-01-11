package file

import (
	"log"
	"os"
)

func Update(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content + "\n"); err != nil {
		log.Println(err)
	}
}

func Rewrite(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content + "\n"); err != nil {
		log.Println(err)
	}
}

func Clear(fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
}

func Content(fileName string) string {
	result, _ := os.ReadFile(fileName)
	return string(result)
}
