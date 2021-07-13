package record

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

// Memcached memcached client
type Memcached struct {
	client *memcache.Client
}

// NewMemcachedConnect memcached connector
func NewMemcachedConnect(servers ...string) *Memcached {
	return &Memcached{client: memcache.New(servers...)}
}

func (m *Memcached) RetrieveDocument(key DocumentKey) (*ReadDocument, error) {
	it, err := m.client.Get(string(key))
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	var rd ReadDocument
	if err := json.Unmarshal(it.Value, &rd.Data); err != nil {
		return nil, err
	}
	rd.Expire = it.Expiration
	rd.Key = DocumentKey(it.Key)
	return &rd, nil
}

func (m *Memcached) SaveDocument(document WriteDocument) error {
	bytes, err := json.Marshal(document.Data)
	if err != nil {
		return err
	}
	it := &memcache.Item{Key: string(document.Key), Value: bytes, Expiration: document.Expire}
	return m.client.Add(it)
}
