package main

import(
	"fmt"
	"flag"
	"os"

	"github.com/cyril-ui-developer/hello-go-program/todo"
)

const todoFileName = ".todo.json"

func main (){
 //fmt.Print("Todo CLI")
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
	//  for _, item := range *l {
    //   fmt.Println(item.Task)
	//  }
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