/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'. 
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(bytes []byte) (int, error) {
	bytes[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

/*
OK!

Program exited.
*/
