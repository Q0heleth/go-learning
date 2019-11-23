package main
import "testing"
func TestHello(t *testing.T){
	got:=Hello("csh")
	want:="Hello, csh"
	if got !=want{
		t.Errorf("Got %q, want %q",got,want)
	}
}