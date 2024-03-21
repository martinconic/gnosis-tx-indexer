package filestore

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/martinconic/gnosis-tx-indexer/pkg/indexer"
	"github.com/martinconic/gnosis-tx-indexer/pkg/storage"
)

type FileStore struct {
	File *os.File
}

func NewFileStore() (storage.Store, error) {
	fs := &FileStore{}
	return fs, nil
}

func (f *FileStore) Open(fp string) error {
	var err error
	f.File, err = os.Create(fp)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (f *FileStore) Close() error {
	return f.File.Close()
}

func (f *FileStore) Put(transactions []indexer.Transaction) error {
	defer f.File.Close()

	// Iterate over each transaction and write it to the file
	for _, transaction := range transactions {
		// Marshal the struct into JSON
		transactionJSON, err := json.Marshal(transaction)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return err
		}
		// Write the JSON string followed by a newline character to the file
		_, err = f.File.Write(append(transactionJSON, '\n'))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	return f.Close()
}

func (f *FileStore) Get() error {
	return nil
}
