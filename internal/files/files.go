package files

import (
	"log"
	"os"
)

func ReadFile(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %s", err)
		return nil
	}

	return data
}
