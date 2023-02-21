package filemap

import (
	"encoding/json"
	"file-spliting-maping/filehandle"
	"fmt"
	"os"
	"strings"
)

type FileMapList []filehandle.PieceInfo

type FileMap struct {
	Content        FileMapList
	ParentFileName string
	ParentFileType string
}

func (flmp FileMapList) AppendFilePieceInfo(filePiece filehandle.PieceInfo) FileMapList {
	flmp = append(flmp, filePiece)
	return flmp
}

func getFileSize(filename string) int {
	file, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error while determining file stats")
		return 0
	} else {
		return int(file.Size())
	}
}

func (flmp FileMapList) Save(fileName string) {
	fmt.Println("Trying to save the maplist")
	fileNamePart := strings.Split(fileName, ".")[0]
	fileTypePart := strings.Split(fileName, ".")[1]
	file, err := os.Create("MapList/" + fileNamePart + ".txt")
	if err != nil {
		fmt.Println("Error while creating map list for filename", fileName)
	} else {
		fileMap := FileMap{
			Content:        flmp,
			ParentFileName: fileName,
			ParentFileType: fileTypePart,
		}
		mapBytes, err := json.Marshal(fileMap)
		if err != nil {
			fmt.Println("error while marshalling MapList")
		} else {
			fmt.Println("Length of maplist :", len(flmp))
			fmt.Println("Length of mapBytes to be writter:", len(mapBytes))
			file.Write(mapBytes)
		}
	}
}

func RetrieveMapList(fileName string) {
	fmt.Println("Trying to retrieve mapList for file ", fileName)
	fileNamePart := strings.Split(fileName, ".")[0]
	file, err := os.Open("MapList/" + fileNamePart + ".txt")
	if err != nil {
		fmt.Println("Error while trying to open the MapList")
		fmt.Println("Ensure this file was properly split, before trying to retreive mapList")
	} else {
		fileMap := &FileMap{}
		bufferSize := getFileSize("MapList/" + fileNamePart + ".txt")
		buffer := make([]byte, bufferSize)
		file.Read(buffer)
		err := json.Unmarshal(buffer, fileMap)
		if err != nil {
			fmt.Println("Error during umarshalling the file mapList")
		} else {
			fmt.Println("Parent file name " + fileMap.ParentFileName)
			fmt.Println("Parent file type " + fileMap.ParentFileType)
			for _, piece := range fileMap.Content {
				fmt.Println("[", piece.PieceName, "|", piece.Piecesize, "|", piece.ReplicationFactor, "]")
			}
		}
	}
}

func InitMapList() FileMapList {
	var fileMaplist FileMapList
	return fileMaplist
}
