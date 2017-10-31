package dh

import "sync"

//* 保存最新的cid,新增的cid为cid+1
type cidSeq struct {
	cid int64
	sync.Mutex
}

//* 获取下一个cid
func (c *cidSeq) next() int64 {
	c.Lock()
	c.cid++
	c.Unlock()
	return c.cid
}

var defaultCidSeq *cidSeq

func getCid() int64 {
	return defaultCidSeq.next()
}
