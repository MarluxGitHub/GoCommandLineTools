package main

import (
	"flag"
	"fmt"
	"marlux/commandlinetools/internal/todo"
	"os"
)

const todoFileName = "todo.json"

func main() {

	task := flag.String(("task"), "", "Task to be Included in the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", -1, "Complete a task")

	flag.Parse()

	l := todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch  {
		case *list:
			for _, item := range l {
				fmt.Printf("%s\n", item.Task)
			}
		case *complete >= 0:
			if err := l.Complete(*complete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *task != "":
			l.Add(*task)

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		default:
			fmt.Println("Please specify a valid flag")
			fmt.Println("---------------------------")

			flag.Usage()
			os.Exit(1)
	}

}