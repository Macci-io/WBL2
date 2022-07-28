package common

import (
	"log"
	"os"
	"sync"
	"syscall"
)

// (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

//Command родитель всех builtin команд и исполнитель не builtin комманд
type Command struct {
	args   []string
	writer int
	reader int
	pid    uintptr
}

//Pid возвращает пид текущей команды
func (c *Command) Pid() uintptr {
	return c.pid
}

//SetPid установить pid
func (c *Command) SetPid(pid uintptr) {
	c.pid = pid
}

//Args геттер аргументов
func (c *Command) Args() []string {
	return c.args
}

//Writer геттер файлового дескриптора для записи
func (c *Command) Writer() int {
	return c.writer
}

//SetWriter сеттер файлового дескриптора для записи
func (c *Command) SetWriter(writer int) {
	c.writer = writer
}

//Reader геттер файлового дескриптора для чтения
func (c *Command) Reader() int {
	return c.reader
}

//SetReader сеттер файлового дескриптора для чтения
func (c *Command) SetReader(reader int) {
	c.reader = reader
}

//NewCommand конструктор команды
func NewCommand(args []string, writer, reader int) *Command {
	return &Command{args, writer, reader, 0}
}

//DupAll для подмены файловых дескрипторов ввода и вывода
func (c Command) DupAll() (ok error) {
	if ok = syscall.Dup2(c.writer, 1); ok != nil {
		return ok
	}
	if ok = syscall.Dup2(c.reader, 0); ok != nil {
		return ok
	}
	c.CloseFds()
	return nil
}

//CloseFds Закрытие файловых дескрипторов если они не дефолтные
func (c Command) CloseFds() {
	var ok error
	if c.writer != 1 {
		if ok = syscall.Close(c.writer); ok != nil {
			log.Fatal(ok)
		}
	}
	if c.reader != 0 {
		if ok = syscall.Close(c.reader); ok != nil {
			log.Fatal(ok)
		}
	}
}

//ForkMe разветление процесса
func (c Command) ForkMe() (pid uintptr) {
	pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	return pid
}

//Run Выполнение команды
func (c *Command) Run(group *sync.WaitGroup) (ok error) {
	pid := c.ForkMe()
	if pid == 0 {
		if ok = c.DupAll(); ok != nil {
			log.Fatal(ok)
		}
		if ok = syscall.Exec(c.args[0], c.args, os.Environ()); ok != nil {
			log.Fatal(ok)
		}
		os.Exit(1)
	}
	c.CloseFds()
	c.SetPid(pid)
	group.Done()
	return nil
}
