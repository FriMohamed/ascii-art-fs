package main
import (
	"ascii/methods"
	"fmt"
	"os"
)
func main() {
	// Check argumment
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Expected: go run . <input_string>\nOr: go run . <input_string> <template_name>")
		return
	}
	// Proccess file conetent
	var fileContent [][]string
	if len(args) == 2 {
		fileContent = methods.ProccessFileContent("standard.txt")
	} else if len(args) == 3 {
		filename := args[2]
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			filename += ".txt"
		}
		fileContent = methods.ProccessFileContent(filename)
	}
		
		
	// Proccess input and print output
	input := args[1]
	if len(input) == 0 {
		return 
	} else {
		fmt.Print(methods.ProccessOutput(input, fileContent))
	}
}
