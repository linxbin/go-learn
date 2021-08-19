package constant_test

import "testing"

const (
	Monday = 1 + iota
	TunesDay
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, TunesDay, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&^Readable == Readable, a&^Writable == Writable, a&^Executable == Executable)
}
