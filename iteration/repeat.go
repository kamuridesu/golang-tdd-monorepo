package iteration

const defaulRepeatCounter int = 5

func Repeat(char string, times ...int) string {
	counter := defaulRepeatCounter
	if len(times) > 0 {
		counter = times[0]
	}
	var repeated string
	for i := 0; i < counter; i++ {
		repeated += char
	}
	return repeated
}
