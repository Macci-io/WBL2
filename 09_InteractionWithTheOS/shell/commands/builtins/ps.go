package builtins

import (
	"errors"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"microshell/shell/commands/common"
	"sync"
	"syscall"
)

//Ps структура реализующая утилиту ps
type Ps struct {
	common.Command
}

//Run запуск этой утилиты
func (p Ps) Run(group *sync.WaitGroup) (ok error) {
	var (
		proceses []ps.Process
	)
	if proceses, ok = ps.Processes(); ok != nil {
		return errors.New("ps: " + ok.Error())
	}

	if _, ok = syscall.Write(p.Writer(), []byte(fmt.Sprintf("pid\tproc\n"))); ok != nil {
		return errors.New("ps: " + ok.Error())
	}
	for _, proc := range proceses {
		if _, ok = syscall.Write(p.Writer(), []byte(fmt.Sprintf("%d\t%s\n", proc.Pid(), proc.Executable()))); ok != nil {
			return errors.New("ps: " + ok.Error())
		}
	}
	group.Done()
	p.CloseFds()
	return nil
}
