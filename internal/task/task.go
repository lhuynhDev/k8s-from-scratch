package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Sheduled
	Running
	Finished
)

type Task struct {
	ID            uuid.UUID
	State         State
	Name          string
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
}
