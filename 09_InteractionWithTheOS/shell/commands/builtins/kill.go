package builtins

import (
	"fmt"
	"microshell/shell/commands/common"
	"os"
	"strconv"
	"sync"
)

//Kill структура реализующая утилиту kill
type Kill struct {
	common.Command
}

//Run запуск этой утилиты
func (k *Kill) Run(group *sync.WaitGroup) (ok error) {
	var (
		pid  int
		args []string
		proc *os.Process
	)
	k.CloseFds()
	defer group.Done()
	args = k.Args()
	if len(args) > 1 {
		for _, v := range args[1:] {
			if pid, ok = strconv.Atoi(v); ok != nil {
				fmt.Println("kill: pid:", v, "is not valid")
			} else {
				if proc, ok = os.FindProcess(pid); ok != nil {
					fmt.Println("kill:", pid, "not found")
				} else if ok = proc.Kill(); ok != nil {
					fmt.Println("kill: " + ok.Error())
				}
			}
		}
	}

	return nil
}

//func killByName(appnames, applist []string) (ok error) {
//	var pid int
//	for _, killingApp := range appnames {
//		for _, strokeApp := range applist {
//			if strings.HasSuffix(strokeApp, killingApp) {
//				strokeApp = strings.TrimLeft(strokeApp, " \t")
//				sp := strings.IndexByte(strokeApp, byte(' '))
//				if pid, ok = strconv.Atoi(strokeApp[:sp]); ok != nil {
//					log.Fatal(ok)
//				} else if ok = syscall.Kill(pid, syscall.SIGKILL); ok != nil {
//					return errors.New("kill: " + ok.Error())
//				}
//			}
//		}
//	}
//	return nil
//}
