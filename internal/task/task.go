package task

import "github.com/google/uuid"

type State int

const (
	Pending State = iota
	Sheduled
	Running
	Finished
)

type Task struct {
	ID    uuid.UUID
	State State
	Name  string
}
