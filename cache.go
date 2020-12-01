package ises

import "time"

type CacheValue struct {
	Time time.Time
	Val interface{}
}
type ICacheKey interface {
	CacheKey() interface{}
}
type CacheKeyInt64 int64

func (k CacheKeyInt64)CacheKey() interface{} {
	return k
}

type Cache struct {
	//key-val-time
	d map[interface{}]*CacheValue
}

func NewCache() *Cache {
	return &Cache{
		d: map[interface{}]*CacheValue{},
	}
}


func (c *Cache)SetCache(k interface{}, v CacheValue) {
	c.d[k] = &v
}
func (c *Cache)GetCache(k interface{}) *CacheValue {
	if v, ok := c.d[k]; ok{
		return v
	}
	return nil
}
func (c *Cache)GetCacheTimeout(k interface{}, t time.Time) *CacheValue {
	if v, ok := c.d[k]; ok{
		if v.Time.Before(t){
			return nil
		}
		return v
	}
	return nil
}
type CacheReLoader func() (*CacheValue, error)
func (c *Cache)ReloadCache(k interface{}, l CacheReLoader) (*CacheValue, error) {
	if v, ok := c.d[k]; ok{
		return v, nil
	}
	v, err := l()
	if err != nil || v == nil{
		return v, err
	}
	c.d[k] = v
	return v, nil
}
func (c *Cache)ReloadCacheTimeout(k interface{}, l CacheReLoader, t time.Time) (*CacheValue, error) {
	if v, ok := c.d[k]; ok{
		if v.Time.After(t){
			return v, nil
		}
	}
	v, err := l()
	if err != nil || v == nil{
		return v, err
	}
	c.d[k] = v
	return v, nil
}