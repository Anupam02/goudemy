package main

import "fmt"

func main() {
	s := []string{"abc", "anupam", "sahil", "malayalam", "test"}
	display(s...)
	filterpallindrome(display, s...)

}

func display(s ...string) {
	fmt.Println("The slice is :")
	for _, sv := range s {
		fmt.Printf("\t%v", sv)
	}
	fmt.Println("")
}

func isPallindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func filterpallindrome(f func(s ...string), s ...string) {
	var si []string
	for _, v := range s {
		if isPallindrom(v) {
			si = append(si, v)
		}
	}
	f(si...)
}
