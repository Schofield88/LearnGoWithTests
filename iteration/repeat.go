package iteration

func Repeat(letter string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += letter
	}
	return repeated
}
