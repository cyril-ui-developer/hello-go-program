package todo


import(
	"testing"
	//"io/ioutil"
	//"os"

	//"github.com/cyril-ui-developer/hello-go-program/todo"
)

func TestAdd(t *testing.T){
	l := List{}

	taskNameOne := "New Task One"
	taskNameTwo := "New Task Two"

	l.Add(taskNameOne)
	l.Add(taskNameTwo)

	if l[0].Task != taskNameOne {
		t.Errorf("Expected. %q, got %q instead,", taskNameOne, l[0].Task)
	}

	if l[1].Task != taskNameTwo {
		t.Errorf("Expected. %q, got %q instead,", taskNameTwo, l[1].Task)
	}
}