package todo_test

import (
	"testing"

	"github.com/VJ-2303/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %q, got %q instead", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}
}
