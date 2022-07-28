package parse

import (
	"errors"
	"fmt"
	"log"
	"microshell/shell/commands"
	"microshell/shell/commands/builtins"
	"microshell/shell/commands/common"
	"os"
	"strings"
	"syscall"
)

func customSplit(data, delim, ignore string) (result []string) {
	var ign, ign2 bool
	var pnt int
	var i int

	data = strings.Trim(data, delim)
	for ; i < len(data); i++ {
		if strings.ContainsRune(ignore, rune(data[i])) && (i == 0 || ign || data[i-1] != '\\') {
			ign = !ign
			ign2 = true
		}
		if !ign {
			if strings.ContainsRune(delim, rune(data[i])) && (i == 0 || ign2 || data[i-1] != '\\') {
				result = append(result, data[pnt:i])
				for i+1 < len(data) && strings.ContainsRune(delim, rune(data[i+1])) {
					i++
				}
				pnt = i + 1
			}
			ign2 = false
		}
	}
	if i != pnt {
		result = append(result, data[pnt:i])
	}
	return result
}

//CreateCommands парсит и создает команды
func CreateCommands(input string, paths []string) (cms []commands.ICommand, ok error) {
	const ignore = "\"'"
	var pipex = make([]int, 2)
	var out, in int
	var cmd commands.ICommand

	if ok = syscall.Pipe(pipex); ok != nil {
		return nil, ok
	}

	out = pipex[1]
	groups := customSplit(input, ";", ignore)
	for _, group := range groups {
		in = 0
		pipeSplit := customSplit(group, "|", ignore)
		for _, cmdline := range pipeSplit {
			args := customSplit(cmdline, " ", ignore)
			for i := range args {
				args[i] = strings.TrimFunc(args[i], func(r rune) bool {
					return r == '\'' || r == '"'
				})
			}
			if cmd, ok = createCommand(args, paths, out, in); ok != nil {
				fmt.Printf("%s\n", ok.Error())
				return nil, nil
			}
			cms = append(cms, cmd)
			in = pipex[0]
			if ok = syscall.Pipe(pipex); ok != nil {
				return nil, ok
			}
			out = pipex[1]
		}
		if ok = syscall.Close(cms[len(cms)-1].Writer()); ok != nil {
			log.Fatal(ok)
		}
		cms[len(cms)-1].SetWriter(1)
		out = pipex[1]
	}
	return cms, nil
}

func checkFile(ut string) (res string, notOk error) {
	stat, notOk := os.Stat(ut)
	if notOk != nil {
		return "", notOk
	} else if stat.IsDir() {
		return "", errors.New(ut + " is directory, can't execute")
	} else if stat.Mode()&0100 == 0 {
		return "", errors.New(ut + " isn't executable, pls make: \n$> chmod +x " + ut)
	}
	return ut, nil
}

func createCommand(args, paths []string, writer, reader int) (res commands.ICommand, notOk error) {
	switch args[0] {
	case "cd":
		return &builtins.Cd{Command: *common.NewCommand(args, writer, reader)}, nil
	case "pwd":
		return &builtins.Pwd{Command: *common.NewCommand(args, writer, reader)}, nil
	case "echo":
		return &builtins.Echo{Command: *common.NewCommand(args, writer, reader)}, nil
	case "kill":
		return &builtins.Kill{Command: *common.NewCommand(args, writer, reader)}, nil
	case "ps":
		return &builtins.Ps{Command: *common.NewCommand(args, writer, reader)}, nil
	case "exec":
		return &builtins.Exec{Command: *common.NewCommand(args, writer, reader)}, nil
	}
	for _, v := range paths {
		if _, notOk = checkFile(v + "/" + args[0]); notOk == nil {
			args[0] = v + "/" + args[0]
			return common.NewCommand(args, writer, reader), nil
		}
	}
	return nil, errors.New(args[0] + ": command not found")
}
