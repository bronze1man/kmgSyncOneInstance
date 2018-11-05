package kmgSyncOneInstance

import "testing"

func TestOneInstance(ot *testing.T){
	var gOi OneInstance
	isRun:=false
	gOi.MustRun(func(){
		isRun=true
	})
	if isRun==false{
		panic("fail")
	}
}