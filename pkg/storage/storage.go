package storage

type Store interface {
	Open(path string) error
	Put() error
	Get() error
	Close() error
}
