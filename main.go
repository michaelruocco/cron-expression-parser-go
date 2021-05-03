package main

import (
	"fmt"
	"os"

	"com.github.michaelruocco/cron-parser/domain"
)

func main() {
	args := removeProgramName(os.Args)
	output, err := domain.Run(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func removeProgramName(args []string) []string {
	return args[1:]
}
