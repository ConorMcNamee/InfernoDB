package main

import (
	"database/internal/store"
	"fmt"
	"os"
)

func SaveToFile(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp", path)
	// Attempts to open file at that path. If none exists it will create it, traunc it and open the file for write-only
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
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

	err = os.Rename(tmp, path)
	return err
}

func main() {

	kvStorage := store.NewKVStore()
	kvStorage.Set("name", []byte("Hello world"))

	name, ok := kvStorage.Get("name")
	if !ok {
		fmt.Println("name does not exist")
	}

	fmt.Println(string(name))
}
