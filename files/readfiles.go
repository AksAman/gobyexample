package files

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal("error: ", e)
	}
}

type SeekWhence int

const (
	SeekFromBegin SeekWhence = iota
	SeekFromCurrent
	SeekFromEnd
)

func Title(s string) {
	fmt.Printf("\n----- %v -----\n", s)
}

func RunReadExample() {
	filePath := "./files/sample-file.txt"

	// ReadFile reads entire file into memory
	Title("os.ReadFile")
	data, err := os.ReadFile(filePath)
	check(err)
	fmt.Printf("%q\n  \n", string(data))
	fmt.Println(string(data))

	// Opening File
	Title("os.Open")
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	// try reading first 5 bytes
	Title("f.Read")
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes read, string(b1): %q\n", n1, string(b1))

	// seeking to location
	// takes two params: offset and whence
	// whence
	// 	0: from begining
	// 	1: from current offset
	//  2: from end (use negative offset)
	Title("f.Seek")
	offset, err := f.Seek(int64(n1+1), int(SeekFromBegin))
	check(err)
	fmt.Printf("ret: %v\n", offset)
	readFromFile(f, offset, 2)

	// using io package to read
	Title("io.ReadAtLeast")
	offset, err = f.Seek(6, int(SeekFromBegin))
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("using io, after seeking %d, %d bytes read, string(b1): %q\n", offset, n3, string(b3))

	Title("rewind using f.Seek")
	_, err = f.Seek(0, 0)
	check(err)
	readFromFile(f, 0, 2)

	Title("using bufio.NewReader and peeking")
	f.Seek(0, 0)
	reader := bufio.NewReader(f)
	b4, err := reader.Peek(5)
	check(err)
	fmt.Printf("string(b4): %q\n", string(b4))

}

func readFromFile(f *os.File, offset int64, nBytes int) {
	b2 := make([]byte, nBytes)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("after seeking %d, %d bytes read, string(b1): %q\n", offset, n2, string(b2))
}
