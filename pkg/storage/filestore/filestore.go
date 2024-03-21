package filestore

import (
	"os"

	"github.com/martinconic/gnosis-tx-indexer/pkg/storage"
)

type FileStore struct {
	File *os.File
}

func NewFileStore(path string) (storage.Store, error) {
	fs := &FileStore{}
	fs.Open(path)

	return fs, nil
}

func (f *FileStore) Open(fp string) error {
	var err error
	f.File, err = os.Open(fp)

	return err
}

func (f *FileStore) Close() error {
	return f.File.Close()
}

func (f *FileStore) Put() error {
	return nil
}

func (f *FileStore) Get() error {
	return nil
}
