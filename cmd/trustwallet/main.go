package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trustwallet/internal/commands"
	"trustwallet/internal/store"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	st := store.NewStore()
	commandHandler := commands.NewCommandHandler(st)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		arrCommandStr := strings.Fields(strings.TrimSuffix(cmdString, "\n"))
		if len(arrCommandStr) == 0 {
			fmt.Println(commands.ErrUnsupportedCommand)
			continue
		}

		var args []string
		if len(arrCommandStr) > 1 {
			args = arrCommandStr[1:]
		}

		result := commandHandler.Handle(arrCommandStr[0], args)
		if result.Error != nil {
			fmt.Println(result.Error)
		} else if result.Data != "" {
			fmt.Println(result.Data)
		}
	}
}
