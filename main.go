package main

import (
	"fmt"
	"os"

	"com.github.michaelruocco/cron-parser/domain"
)

func main() {
	args := removeProgramName(os.Args)
	santizedArgs, err := domain.Sanitize(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := domain.Parse(santizedArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	formattedResult := domain.Format(result)
	fmt.Println(formattedResult)
}

func removeProgramName(args []string) []string {
	return args[1:]
}
