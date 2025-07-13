package main

func main() {
	// anonymous empty struct type
	empty := struct{}{}
	_ = empty

	// named empty struct type
	type emptyStruct struct{}
	namedEmpty := emptyStruct{}
	_ = namedEmpty
}
