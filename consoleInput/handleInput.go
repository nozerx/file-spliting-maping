package consoleInput

import (
	"bufio"
	"file-spliting-maping/filehandle/fileinput"
	"fmt"
	"os"
	"runtime"
)

func HandleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading input")
		} else {
			if input[:3] == "<s>" {
				fmt.Println("File send command")
				esacapeSeqLen := 0
				if runtime.GOOS == "windows" {
					esacapeSeqLen = 2
				} else {
					esacapeSeqLen = 1
				}
				fileName := input[3 : len(input)-esacapeSeqLen]
				fileinput.HandleFile(fileName)
			}
		}
	}
}
