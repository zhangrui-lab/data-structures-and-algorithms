package to

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const saveQueueLength = 1000

type record struct {
	Key, Value string
}

type UrlStore struct {
	data map[string]string
	save chan record
	mu   sync.RWMutex
	fp   *os.File
}

// NewUrlStore  urlStore工厂
func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{
		data: make(map[string]string),
	}
	if len(filename) > 0 {
		fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln("Error opening URLStore:", err)
		}
		s.fp = fp
		s.save = make(chan record, saveQueueLength)
		if err := s.load(); err != nil {
			log.Fatalln("Error loading data in URLStore:", err)
		}
		go s.saveLoop()
	}
	return s
}

func (s *UrlStore) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.save)
	time.Sleep(5 * time.Second)
	_ = s.fp.Close()
	log.Println("UrlStore Closed...")
}

// Put 新增url
func (s *UrlStore) Put(url, key *string) error {
	*key = s.genKey(s.Count())
	if !s.set(*key, *url) {
		return errors.New("key already exists")
	}
	if s.save != nil {
		s.save <- record{Key: *key, Value: *url}
	}
	return nil
}

// Get 获取原始连接
func (s *UrlStore) Get(key, url *string) error {
	r, ex := s.get(*key)
	if !ex {
		return errors.New(fmt.Sprintf("Key %s Not Found!", *key))
	}
	*url = r
	return nil
}

// Count 已设置url
func (s *UrlStore) Count() int {
	return s.count()
}

// set 新增键值对
func (s *UrlStore) set(key, val string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ex := s.data[key]; ex {
		return false
	}
	s.data[key] = val
	return true
}

// get 获取键值
func (s *UrlStore) get(key string) (url string, ex bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ex = s.data[key]
	return
}

// count 获取键值对数量
func (s *UrlStore) count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// genKey 设置短url
func (s *UrlStore) genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	str := make([]byte, 20)
	i := len(str)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		str[i] = keyChar[j]
	}
	return string(str[i:])
}

// saveLoop key-value 持久化
func (s *UrlStore) saveLoop() {
	e := json.NewEncoder(s.fp)
	for r := range s.save {
		if err := e.Encode(r); err != nil {
			log.Fatalln(err)
		}
	}
}

// load 从持久化文件加载 short-url 数据
func (s *UrlStore) load() error {
	if _, err := s.fp.Seek(0, 0); err != nil {
		return err
	}
	s.mu.Lock()
	s.mu.Unlock()
	d := json.NewDecoder(s.fp)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.data[r.Key] = r.Value
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
