package methods

import "strings"

// proccess the input to be ready to print
func ProccessTheInput(s string) []string {
	// Remove the characters \a,\b,\r,\f and then split input with \\n..........................................
	s = removeEscapeChars(strings.ReplaceAll(s, "\\t", "   "))
	splitedInput := strings.Split(s, "\\n")

	// if the input consists of just new lines after split it we get an extra empty string so we remove it......
	if CheckNewLine(splitedInput) { 
		return splitedInput[:len(splitedInput)-1]
	}

	return splitedInput
}

// Remove non printiple and non ascii characters from the input
func removeEscapeChars(s string) string {
	var str string
	i := 0
	for i < len(s) {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'a', 'b', 'v', 'f', 'r':
				i += 2
				continue
			}
		} else if s[i] < 32 || s[i] > 126 {
			i++
			continue 
		}
		str += string(s[i])
		i++
	}

	return str
}

// Check if the string consists of just new lines
func CheckNewLine(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) != 0 {
			return false
		} 
	}
	
	return true
}
