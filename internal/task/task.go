package task

type State int

const (
	Pending State = iota
	Sheduled
	Running
	Finished
)
