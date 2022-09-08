package embed_directive

import (
	"embed"
	"fmt"
)

/**
mkdir -p embed_directive/folder.ignore
echo "hello go" > embed_directive/folder.ignore/single_file.txt
echo "123" > embed_directive/folder.ignore/file1.hash
echo "456" > embed_directive/folder.ignore/file2.hash
**/

//go:embed folder.ignore/single_file.txt
var fileString string

//go:embed folder.ignore/single_file.txt
var fileByte []byte

// embed.FS acts like virtual FileSystem
//
//go:embed folder.ignore
var folder embed.FS

func Run() {

	fmt.Println("\n----- after embed")

	fmt.Printf("fileString: %v\n", fileString)
	fmt.Printf("string(fileByte): %v\n", string(fileByte))

	c, _ := folder.ReadDir("folder.ignore")
	if len(c) == 0 {
		fmt.Println("no files found")
	}
	for index, entry := range c {
		fmt.Println(index+1, ")", entry.Name(), "entry.IsDir():", entry.IsDir())
	}

	hash1, _ := folder.ReadFile("folder.ignore/file1.hash")
	fmt.Printf("string(hash1): %v\n", string(hash1))

	hash2, _ := folder.ReadFile("folder.ignore/file2.hash")
	fmt.Printf("string(hash2): %v\n", string(hash2))

}
