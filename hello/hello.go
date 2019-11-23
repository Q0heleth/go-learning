package main
import "fmt"
const englishHelloprefix="Hello,"
func Hello(s string) string{
	if s==""{
		s="World"
	}
	return englishHelloprefix+s
}
func main(){
	fmt.Println(Hello("world"))
}