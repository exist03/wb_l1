package main

var justString string

func createHugeString(size int) string {
	return string(make([]rune, size))
}

func trim(v string, from, to int) string {
	tmp := []rune(v)
	return string(append(tmp[:from], tmp[to:]...))
}

func someFunc() {

	v := createHugeString(1 << 10)

	justString = trim(v, 0, 100)
}

func main() {

	someFunc()
}
