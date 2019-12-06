package concurrency

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteCheck func(url string) bool
type result struct{
	url string
	ret bool
}
func Checkwebsites(wc WebsiteCheck, urls []string) map[string]bool {
	c:=make(chan result)
	m := make(map[string]bool)
	for _, url := range urls {
		go func(u string) {
			c<-result{u,wc(u)}
		}(url)
	}
	for i:=0;i<len(urls);i++{
		res:=<-c
		m[res.url]=res.ret
	}
	//time.Sleep(2 * time.Second)
	return m
}
func TestWebcheck(t *testing.T) {
	t.Run("mockWebcheck", func(t *testing.T) {
		urls := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		got := Checkwebsites(mockWebcheck, urls)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
func mockWebcheck(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}
func sleepcheck(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func Benchmark(t *testing.B) {
	urls := make([]string, 100)
	for i, _ := range urls {
		urls[i] = "a url"
	}
	for i := 0; i < t.N; i++ {
		Checkwebsites(sleepcheck, urls)
	}
}
