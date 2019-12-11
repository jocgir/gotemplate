package template

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
)

// Console starts the current template in interactive mode
func (t *Template) Console() {
	shell := ishell.New()
	shell.Println("gotemplate interactive console")
	if user, err := user.Current(); err == nil {
		historyFile := path.Join(user.HomeDir, ".gotemplate")
		os.MkdirAll(historyFile, 0755)
		shell.SetHistoryPath(path.Join(historyFile, "history"))
	}

	run := func(c *ishell.Context) {
		result, err := t.ProcessContent(strings.Join(c.RawArgs, " "), "")
		if err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		}
		fmt.Println(color.HiWhiteString(result))
	}

	shell.AddCmd(&ishell.Cmd{
		Name:    "razor",
		Aliases: []string{"r"},
		Func: func(c *ishell.Context) {
			c.RawArgs = c.RawArgs[1:]
			result, _ := t.applyRazor([]byte(strings.Join(c.RawArgs, " ")))
			fmt.Println(color.CyanString(string(result)))
			run(c)
		},
		Help: "Evaluate razor expression and return the go template evaluation",
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "history",
		Aliases: []string{"r"},
		Func: func(c *ishell.Context) {
			result, _ := t.applyRazor([]byte(strings.Join(c.RawArgs[1:], " ")))
			fmt.Println(string(result))
		},
		Help: "List all available functions",
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Func: func(c *ishell.Context) { t.PrintFunctions(false, false, false, c.Args...) },
		Help: "List all available functions",
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "long",
		Func: func(c *ishell.Context) { t.PrintFunctions(true, true, true, c.Args...) },
		Help: "List all available functions",
	})

	shell.NotFound(run)

	shell.Run()
	shell.Close()
}
