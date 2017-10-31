package dh

import (
	"sync"
	"time"
)

//* 用户设置自定义对象时用到的锁
var userCustomOPLock *sync.Mutex

//* 用户设置自定义连接器的接口
func SetConnection(c Connection) {
	userCustomOPLock.Lock()
	conn = c
	userCustomOPLock.Unlock()
}

//* 用户设置自定义连接超时时长接口
//* 默认60秒
func SetTimeout(timeout time.Duration) {
	userCustomOPLock.Lock()
	timeoutSecond = timeout
	userCustomOPLock.Unlock()
}
