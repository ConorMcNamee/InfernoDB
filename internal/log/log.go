package log

import (
	"encoding/binary"
	"fmt"
	"os"
)

type AppendOnlyLog struct {
	file *os.File
	path string
}

func CreateAppendOnlyLogger(path string) *AppendOnlyLog {
	return &AppendOnlyLog{
		file: &os.File{},
		path: path,
	}
}

func EncodeKeyAndValue(key string, value []byte) []byte {
	data := make([]byte, 0)

	KeyLen := make([]byte, 4)
	// Append key length and key
	binary.BigEndian.PutUint16(KeyLen, uint16(len(key)))
	data = append(data, KeyLen...)
	data = append(data, []byte(key)...)

	valueLen := make([]byte, 4)
	binary.BigEndian.PutUint16(valueLen, uint16(len(value)))

	// Append value length and value
	data = append(data, valueLen...)
	data = append(data, value...)

	return data
}

func (wal *AppendOnlyLog) SaveToFile(key string, value []byte) error {
	data := EncodeKeyAndValue(key, value)
	tmp := fmt.Sprintf("%s.tmp", wal.path)
	// Attempts to open file at that path. If none exists it will create it, traunc it and open the file for write-only
	file, err := os.OpenFile(wal.path, os.O_APPEND, 0664)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = file.Write(data); err != nil {
		return err
	}

	if err = file.Sync(); err != nil {
		return err
	}

	err = os.Rename(tmp, wal.path)
	return err
}
