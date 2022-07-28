package builtins

import (
	"microshell/shell/commands/common"
	"os"
	"sync"
	"syscall"
)

//Pwd структура реализующая утилиту pwd
type Pwd struct {
	common.Command
}

//Run запуск этой утилиты
func (p Pwd) Run(group *sync.WaitGroup) (ok error) {
	var wd string

	if wd, ok = os.Getwd(); ok != nil {
		return ok
	} else if _, ok = syscall.Write(p.Writer(), []byte(wd+"\n")); ok != nil {
		return ok
	}
	p.CloseFds()
	group.Done()
	return nil
}
