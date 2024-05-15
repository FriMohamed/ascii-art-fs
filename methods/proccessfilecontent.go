package methods

import (
	"fmt"
	"os"
	"strings"
)

// Get the file content and split it to get each character graphic representation alone, where also each character represrntation splited to 8 lines
func ProccessFileContent(fileName string) [][]string {
	// get file content............................................................
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	thinkertoy := 0
	var splitedContent []string
	// split the content by character graphics.....................................
	if fileContent[0] != '\r' {
		splitedContent = strings.Split(string(fileContent[1:]), "\n\n")
	} else {
		splitedContent = strings.Split(string(fileContent[2:]), "\r\n\r\n")
		thinkertoy = 1
	}

	// split character graphics to lines...........................................
	graphics := splitGraphics(splitedContent, thinkertoy)

	return graphics
}

// used to help in proccessing the file content
func splitGraphics(slice []string, thinkertoy int) [][]string {

	sep := "\n"
	if thinkertoy == 1 {
		sep = "\r\n"
	}
	var ret [][]string
	for i := 0; i < len(slice); i++ {
		ret = append(ret, strings.Split(slice[i], sep))
	}

	return ret
}
