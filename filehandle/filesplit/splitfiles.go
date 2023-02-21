package filesplit

import (
	"file-spliting-maping/filehandle"
	"file-spliting-maping/filehandle/filemap"
	"fmt"
	"os"
)

func SplitFile(fileInfo filehandle.FileInfo, file *os.File) {
	buffer := make([]byte, filehandle.BufferSize)
	iterationCount := fileInfo.FileSize / filehandle.BufferSize
	leftByteSize := filehandle.BufferSize
	if fileInfo.FileSize%filehandle.BufferSize != 0 {
		iterationCount += 1
		leftByteSize = fileInfo.FileSize % filehandle.BufferSize
	}
	lastPieceBuffer := make([]byte, leftByteSize)
	fileInfo.FilePieces = iterationCount
	fileMapList := filemap.InitMapList()
	for i := 0; i < iterationCount; i++ {
		if i == (iterationCount - 1) {
			file.Read(lastPieceBuffer)
			tempFileName := fmt.Sprintf("part_%d.%s", i, fileInfo.FileType)
			tempFile, err := os.Create("temp/" + tempFileName)
			if err != nil {
				fmt.Println("Error during creating the temp file :", tempFileName)
			} else {
				tempFile.Write(lastPieceBuffer)
				tempFile.Close()
				pieceInfo := filehandle.ComposePieceInfo(tempFileName, leftByteSize)
				fileMapList = fileMapList.AppendFilePieceInfo(pieceInfo)
				fmt.Println("Handled file ", tempFileName)
			}
		} else {
			file.Read(buffer)
			tempFileName := fmt.Sprintf("part_%d.%s", i, fileInfo.FileType)
			tempFile, err := os.Create("temp/" + tempFileName)
			if err != nil {
				fmt.Println("Error during creating the temp file :", tempFileName)
			} else {
				tempFile.Write(buffer)
				tempFile.Close()
				pieceInfo := filehandle.ComposePieceInfo(tempFileName, filehandle.BufferSize)
				fileMapList = fileMapList.AppendFilePieceInfo(pieceInfo)
				fmt.Println("Handled file ", tempFileName)
			}
		}
	}
	fileMapList.Save(fileInfo.FileName)
}
