// pkg/storage/db.go
package storage

import (
	"encoding/json"
	"fmt"
	"interview-be-earning/pkg/blockchain"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	instance *leveldb.DB
}

func OpenDB(path string) (*DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &DB{instance: db}, nil
}

func (d *DB) SaveBlock(height int, block *blockchain.Block) error {
	key := fmt.Sprintf("block_%d", height)
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return d.instance.Put([]byte(key), data, nil)
}

func (d *DB) LoadBlock(height int) (*blockchain.Block, error) {
	key := fmt.Sprintf("block_%d", height)
	data, err := d.instance.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	var block blockchain.Block
	err = json.Unmarshal(data, &block)
	return &block, err
}

func (d *DB) LatestHeight() int {
	i := 0
	for {
		key := fmt.Sprintf("block_%d", i)
		_, err := d.instance.Get([]byte(key), nil)
		if err != nil {
			return i - 1
		}
		i++
	}
}

func (d *DB) Close() {
	d.instance.Close()
}
