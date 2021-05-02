package repeats

const numRepeats = 4

func Repeat(input string) string {
	var output string
	for i := 0; i < numRepeats; i++ {
		output += input
	}
	return output
}
