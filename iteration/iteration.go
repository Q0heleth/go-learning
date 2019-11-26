package iteration

var repeatCount = 5

func Repeat(s string) string {
	var ret string
	for i := 0; i < repeatCount; i++ {
		ret += s
	}
	return ret
}
