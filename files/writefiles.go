package files

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

func readEntireFile(filename string) {
	if data, err := os.ReadFile(filename); err != nil {
		fmt.Printf("error: %v while reading file: %q\n", err, filename)
	} else {
		fmt.Printf("reading data from file: %q:\n %v\nEOF\n", filename, string(data))
	}
}

func RunWriteExample() {
	Title("**** write-files ****")

	Title("os.WriteFile")
	d1 := []byte("hello\nworld\nfrom\ngolang")
	filename := "/tmp/dat1"
	filemode := fs.FileMode(0644)
	err := os.WriteFile(filename, d1, filemode)
	check(err)

	readEntireFile(filename)

	Title("os.Create")
	filename = "/tmp/dat2"
	f, err := os.Create(filename)
	check(err)
	defer func() {
		Title("closing file")
		f.Close()
	}()

	Title("writing bytes")
	d2 := []byte{115, 111, 109, 105, 10}
	bytesWritten, err := f.Write(d2)
	check(err)
	fmt.Printf("bytesWritten: %v\n", bytesWritten)

	Title("f.WriteString")
	bytesWritten, err = f.WriteString("Some more string3\n")
	check(err)
	fmt.Printf("bytesWritten: %v\n", bytesWritten)
	f.Sync()

	Title("Bufio Writer")
	writer := bufio.NewWriter(f)
	defer func() {
		Title("flusing to file")
		defer writer.Flush()
	}()

	bytesWritten, err = writer.WriteString("from bufio\n")
	check(err)

	fmt.Printf("bytesWritten: %v\n", bytesWritten)

}
