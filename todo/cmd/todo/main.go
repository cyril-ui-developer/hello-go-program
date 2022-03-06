package main

import(
	"fmt"
	"flag"
	"os"

	"github.com/cyril-ui-developer/hello-go-program/todo"
)

var todoFileName = ".todo.json"

func main (){

if os.Getenv("TODO_FILENAME") != "" {
	todoFileName = os.Getenv("TODO_FILENAME")
}	

 flag.Usage = func(){
	 fmt.Fprintf(flag.CommandLine.Output(),
	"%s tool. A todo CLI developed by CY \n", os.Args[0])

	flag.PrintDefaults()
 } 

 task := flag.String("task", "", "Add a new task to te list")
 list := flag.Bool("list", false, "List all the tasks")

 flag.Parse()

 l := &todo.List{}

 if err := l.Get(todoFileName); err != nil {
	 fmt.Fprintln(os.Stderr, err)
	 os.Exit(1)
 }

 switch {
 case *list:
	fmt.Print(l)
	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
default:
			fmt.Fprintln(os.Stderr, "Invalid option")
			os.Exit(1)
 }
}