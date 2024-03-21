package storage

import "github.com/martinconic/gnosis-tx-indexer/pkg/indexer"

type Store interface {
	Open(string) error
	Put([]indexer.Transaction) error
	Get() error
	Close() error
}
