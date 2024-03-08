package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ade-iqbal/fga-assignment-3/models"
)

func GetFile() string {
	file, _ := os.OpenFile("info.json", os.O_RDONLY, 0606)
	defer file.Close()

	reader := bufio.NewReader(file)
	var content string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		content += string(line)
	}

	fmt.Println(content)
	return content
}

func UpdateFile(content models.Weather) {
	weather, _ := json.MarshalIndent(content, "", "    ")
	
	file, _ := os.OpenFile("info.json", os.O_WRONLY|os.O_TRUNC, 0606)
	defer file.Close()
	
	_, err := file.WriteString(string(weather))
	if err != nil {
		file.WriteString(string(weather))
	}
}