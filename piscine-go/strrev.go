package main

func StrRev(s string) string {
	var reverse string

	for i := len(s) - 1; i >= 0; i-- {
		reverse = reverse + string(s[i])
	}
	return reverse
}
