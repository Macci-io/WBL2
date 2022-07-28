package processing

import (
	"bufio"
	"log"
	"microshell/shell/commands"
	"microshell/shell/parse"
	"os"
	"strings"
)

//ReadLine точка входа в нашу программу
func ReadLine() {
	const prefix = "\033[1;31m $> \033[0m "
	var (
		ok    error
		line  []byte
		cmds  []commands.ICommand
		paths []string
	)
	env := os.Getenv("PATH")
	if env != "" {
		paths = strings.Split(env, ":")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		print(prefix)
		if line, _, ok = reader.ReadLine(); ok != nil {
			log.Fatal(ok)
		} else if string(line) == "\\quit" {
			break
		} else if cmds, ok = parse.CreateCommands(string(line), paths); ok != nil {
			log.Fatal(ok)
		}
		commands.ExecuteAll(cmds)
	}
}
