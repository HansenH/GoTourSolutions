package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (myReader MyReader) Read(char []byte) (int, error) {
	n := len(char)
	for i := 0; i < n; i++ {
		char[i] = 'A'
	}
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
