package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents a ToDo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add Creates a new todo and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks a ToDo item as Completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete Method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists", i)
	}
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as a JSON and saves it
// using the provided filename
func (l *List) Save(filename string) error {
	js, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0o644)
}

// Get method opens the provided filename, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, l)
}
