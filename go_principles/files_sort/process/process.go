package process

import (
	"strconv"
	"strings"
)

type Process struct {
	Id uint64
	Priority int64
	Time uint64
	Status string
}

func (process *Process) String() string {
	builder := strings.Builder{}

	builder.WriteString("Process { Id: ")
	builder.WriteString(strconv.FormatUint(process.Id, 10))
	builder.WriteString(", Priority: ")
	builder.WriteString(strconv.FormatInt(process.Priority, 10))
	builder.WriteString(", Time: ")
	builder.WriteString(strconv.FormatUint(process.Time, 10))
	builder.WriteString(", Status: ")
	builder.WriteString(process.Status)
	builder.WriteString(" }")

	return builder.String()
}

type ByPriority []Process

func (p ByPriority) Len() int {
	return len(p)
}

func (p ByPriority) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPriority) Less(i, j int) bool {
	return p[i].Priority < p[j].Priority
}