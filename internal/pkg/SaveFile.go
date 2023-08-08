package pkg

import (
	"fmt"
	"os"
)

func SaveFile(value string) error {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %v\n", value)))
	if err != nil {
		return err
	}
	return nil
}
