package log

import "os"

type AppendOnlyLog struct {
	file *os.File
}

func CreateAppendOnlyLogger() *AppendOnlyLog {
	return &AppendOnlyLog{
		file: &os.File{},
	}
}

func (wal *AppendOnlyLog) AppendToLog([]byte) {
	file, err := os.OpenFile("log.db", os.O_APPEND, 0664)
}
