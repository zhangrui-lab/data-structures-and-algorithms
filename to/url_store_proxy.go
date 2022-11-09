package to

import (
	"log"
	"net/rpc"
)

// UrlStoreProxy Slave代理节点
type UrlStoreProxy struct {
	urls   *UrlStore
	client *rpc.Client
}

// Put 新增映射
func (s *UrlStoreProxy) Put(url, key *string) error {
	if err := s.client.Call("to.Store.Put", url, key); err != nil {
		return err
	}
	return nil
}

// Get 获取短链对应的原始域名
func (s *UrlStoreProxy) Get(key, url *string) error {
	if err := s.urls.Get(key, url); err == nil {
		return nil
	}
	if err := s.client.Call("to.Store.Get", key, url); err != nil {
		return err
	}
	s.urls.set(*key, *url)
	return nil
}

// NewUrlStoreProxy 获取代理存储
func NewUrlStoreProxy(addr string) *UrlStoreProxy {
	c, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	return &UrlStoreProxy{
		urls:   NewUrlStore(""),
		client: c,
	}
}
