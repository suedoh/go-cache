package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
    lock sync.RWMutex
    data map[string][]byte
}

func NewCache() *Cache {
  return &Cache{
      data: make(map[string][]byte),
  }
}

func (c *Cache) Get(key []byte) ([]byte, error)  {
   c.lock.RLock() 
   defer c.lock.RUnlock() 

   keyStr := string(key)

   val, ok := c.data[keyStr]
   if !ok {
        return nil, fmt.Errorf("key (%s) not found", keyStr)
   }

   return val, nil
}
