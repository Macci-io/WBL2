package builtins

import (
	"errors"
	"microshell/shell/commands/common"
	"os"
	"sync"
)

//Cd структура реализующая утилиту cd
type Cd struct {
	common.Command
}

//Run запуск этой утилиты
func (c Cd) Run(group *sync.WaitGroup) (ok error) {
	args := c.Args()
	c.CloseFds()
	defer group.Done()
	if len(args) != 2 {
		return errors.New("cd: too many arguments")
	}
	if ok = os.Chdir(args[1]); ok != nil {
		return errors.New("cd: " + ok.Error())
	}

	return nil
}
