package filehandle

type FileInfo struct {
	FileName      string
	FileType      string
	FileSize      int
	FilePieces    int
	FileProcessed bool
}

const BufferSize int = 8388608 // 8 mb is the buffer size set now

type PieceInfo struct {
	PieceName         string
	Piecesize         int
	ReplicationFactor int
	// Replica1Peer      peer.ID
	// Replica2Peer      peer.ID
	// Replica3Peer      peer.ID
}

func ComposePieceInfo(pieceName string, pieceSize int) PieceInfo {
	return PieceInfo{
		PieceName:         pieceName,
		Piecesize:         pieceSize,
		ReplicationFactor: 3,
	}
}
