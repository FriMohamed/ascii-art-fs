package methods

// Proccess the output
func ProccessOutput(s string, graphics [][]string) string {
	input := ProccessTheInput(s)
	
	const graphHeight int = 8
	output := ""
	for _, str := range input {
		if len(str) == 0 {
			output += "\n"
		} else {
			for j := 0; j < graphHeight; j++ {
				for k := 0; k < len(str); k++ {
					if k == len(str)-1 {
						output += graphics[str[k]-32][j] + "\n"
					} else {
						output += graphics[str[k]-32][j]
					}
				}
			}
		}
	}
	return output
}
