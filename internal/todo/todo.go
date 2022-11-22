package todo

import (
	"fmt"
	"time"
)

type todo struct {
	Task	string
	Done	bool
	CreatedAt	time.Time
	CompletedAt	time.Time
}

func (t *todo) String() string {
	prefix := "  "

	if t.Done {
		prefix = "X "
	}

	return fmt.Sprintf("%s%s", prefix,t.Task)
}

func (t *todo) VerboseString() string {
	prefix := "  "

	if t.Done {
		prefix = "X "
	}

	return fmt.Sprintf("%s%s %s %s", prefix,t.Task, t.CreatedAt.String(), t.CompletedAt.String())
}