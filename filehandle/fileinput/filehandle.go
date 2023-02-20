package fileinput

import (
	"file-spliting-maping/filehandle"
	"file-spliting-maping/filehandle/filesplit"
	"fmt"
	"os"
	"strings"
)

func getFileSize(filename string) int {
	file, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error while determining file stats")
		return 0
	} else {
		return int(file.Size())
	}
}

func composeFileInfo(filename string, filetype string) filehandle.FileInfo {
	return filehandle.FileInfo{
		FileName:      filename,
		FileType:      filetype,
		FileSize:      getFileSize(filename),
		FilePieces:    0,
		FileProcessed: false,
	}
}

func HandleFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error during tyring to open the file for processing it")
		fmt.Println("file name is :", fileName, "[Result: Unable to open]")
		return
	} else {
		fileType := strings.Split(fileName, ".")[1]
		fileInfo := composeFileInfo(fileName, fileType)
		filesplit.SplitFile(fileInfo, file)
	}

}
