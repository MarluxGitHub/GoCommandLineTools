package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"marlux/commandlinetools/internal/todo"
	"os"
	"strconv"
	"strings"
)

const todoFileName = "todo.json"

func main() {

	add := flag.Bool(("add"), false, "Task to be Included in the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", -1, "Complete a task")
	del := flag.Int("delete", -1, "Delete a task")
	verbose := flag.Bool("verbose", false, "Verbose output")
	openTodo := flag.Bool("open", false, "Show only OpenTodos")

	flag.Parse()

	l := todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch  {
		case *list:
			for key, item := range l {
				if *openTodo && item.Done {
					continue
				}

				if *verbose {
					fmt.Println(strconv.Itoa(key) + ".: " + item.VerboseString())
				} else {
					fmt.Println(strconv.Itoa(key) + ".: " + item.String())
				}
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
		case *add:
			t, err := getTask(os.Stdin, flag.Args()...)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			l.Add(t)

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *del >= 0:
			if err := l.Delete(*del); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
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

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)

	s.Scan()

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be empty")
	}

	return s.Text(), nil
}