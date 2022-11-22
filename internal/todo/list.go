package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type List []todo

func  (list *List) Add(task string) {
	t := todo{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Now(),
	}

	*list = append(*list, t)
}

func (list *List) Complete(i int) error {
	if(i < 0 || i >= len(*list)) {
		return fmt.Errorf("item %d does not exist", i)
	}

	(*list)[i].Done = true
	(*list)[i].CompletedAt = time.Now()

	return nil
}

func (list *List) Delete(i int) error {
	if(i < 0 || i >= len(*list)) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*list = append((*list)[:i], (*list)[i+1:]...)

	return nil
}

func (list *List) Save(filename string) error {
	js, err := json.Marshal(list)

	if err != nil { return err }

	return os.WriteFile(filename, js, 0644)
}

func (list *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if(err != nil) {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, list)
}
