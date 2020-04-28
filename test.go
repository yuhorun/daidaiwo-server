package main

func main() {
	c := make(map[int]int)
	c[22] = 1
	s := c
	println(s[22])
	c[22] = 0
	println(s[22])
}
