package main
import "testing"
func TestHello(t *testing.T){
	assertCorrectMessage:=func(t *testing.T,got,want string){
		t.Helper()
		if got!=want{
			t.Errorf("got '%q' want '%q'",got,want)
		}
	}
	t.Run("saying hello to people",func (t *testing.T){
		got:=Hello("csh","")
	want:="Hello,csh"
	assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello world when name is empty",func(t *testing.T){
		got:=Hello("","")
		want:="Hello,World"
		assertCorrectMessage(t,got,want)
	})
	t.Run("language test of Spanish",func(t *testing.T){
		got:=Hello("Elodie","Spanish")
		want:="Hola,Elodie"
		assertCorrectMessage(t,got,want)
	})
	t.Run("language test of French",func(t *testing.T){
		got:=Hello("Mark","French")
		want:="Bonjour,Mark"
		assertCorrectMessage(t,got,want)
	})
}