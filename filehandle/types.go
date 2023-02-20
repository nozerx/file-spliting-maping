package filehandle

type FileInfo struct {
	FileName      string
	FileType      string
	FileSize      int
	FilePieces    int
	FileProcessed bool
}

const BufferSize int = 8388608 // 8 mb is the buffer size set now
