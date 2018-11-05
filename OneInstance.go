package kmgSyncOneInstance

import (
	"sync"
	"fmt"
)

// 只允许 一个实例运行
type OneInstance struct{
	locker sync.Mutex
	isRun bool
	Name string // 对象建立后请不要修改，否则会data race。
}

func (oi *OneInstance) RunAndLogIfNotRun(fn func()) {
	ret:=oi.Run(fn)
	if ret == false{
		fmt.Println("[kmgSync.OneInstance] too many instance2 "+oi.Name)
	}
}

func (oi *OneInstance) IsRun() (isRun bool) {
	oi.locker.Lock()
	isRun = oi.isRun
	oi.locker.Unlock()
	return isRun
}

// 返回true表示 本次运行过了，返回false 表示 本次没有运行过。
func (oi *OneInstance) Run(fn func()) bool{
	oi.locker.Lock()
	if oi.isRun{
		oi.locker.Unlock()
		return false
	}
	oi.isRun = true
	oi.locker.Unlock()
	defer func(){
		oi.locker.Lock()
		oi.isRun = false
		oi.locker.Unlock()
	}()
	fn()
	return true
}

// 多于一个并发 会 panic
func (oi *OneInstance) MustRun(fn func()){
	oi.locker.Lock()
	if oi.isRun{
		oi.locker.Unlock()
		panic("[kmgSync.OneInstance] too many instance1 "+oi.Name)
		//return
	}
	oi.isRun = true
	oi.locker.Unlock()
	defer func(){
		oi.locker.Lock()
		oi.isRun = false
		oi.locker.Unlock()
	}()
	fn()
	return
}