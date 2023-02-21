package consoleInput

import (
	"bufio"
	"file-spliting-maping/filehandle/fileinput"
	"file-spliting-maping/filehandle/filemap"
	"fmt"
	"os"
	"runtime"
)

func HandleInput() {
	reader := bufio.NewReader(os.Stdin)
	esacapeSeqLen := 0
	if runtime.GOOS == "windows" {
		esacapeSeqLen = 2
	} else {
		esacapeSeqLen = 1
	}
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading input")
		} else {
			if input[:3] == "<s>" {
				fmt.Println("File send command")
				fileName := input[3 : len(input)-esacapeSeqLen]
				go fileinput.HandleFile(fileName)
			}
			if input[:3] == "<m>" {
				fileName := input[3 : len(input)-esacapeSeqLen]
				fmt.Println("Retrieve file peice map command for []", fileName, "]")
				go filemap.RetrieveMapList(fileName)
			}
		}
	}
}
