package task

import (
	"time"
)

type Task struct {
	Id                 int       `json:"Id"`
	DateCreated        time.Time `json:"DateCreated"`
	DateCompleted      time.Time `json:"DateCompleted"`
	ResponsibleUser    int       `json:"ResponsibleUser"`
	AssignedUser       int       `json:"AssignedUser"`
	CreatedBy          int       `json:"CreatedBy"`
	LocationId         int       `json:"LocationId"`
	Stage              Stage     `json:"Stage"`
	CancellationReason string    `json:"CancellationReason"`
	Deadline           time.Time `json:"Deadline"`
	Message            string    `json:"Message"`
	Title              string    `json:"Title"`
	Description        string    `json:"Description"`
	Priority           Priority  `json:"Priority"`
}
