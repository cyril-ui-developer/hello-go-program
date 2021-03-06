package todo

import(
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task	string
	Done	bool
	CreatedAt	time.Time
	CompletedAt	time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:	task,
		Done:	false,
		CreatedAt:	time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Save(filename string) error {
 js, err := json.Marshal((l))

 if err != nil {
	 return err
 }

 return ioutil.WriteFile(filename, js, 0644)
}


func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
        return err
		}
	}
	if len(file) == 0{
		return nil
	}

	return json.Unmarshal(file, l)
}