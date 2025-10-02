package taskmanagementservice

import (
	"time"
)

type Stage int

const (
	Created = iota
	Assigned
	InWork
	Review
	Completed
	SentForRevision
	Canceled
	Expired
)

type Priority int

const (
	Low = iota
	Medium
	High
)

type Task struct {
	Id                 int       `json:"id"`
	DateCreated        time.Time `json:"DateCreated"`
	DateCompleted      time.Time `json:"DateCompleted"`
	ResponsibleUser    int       `json:"ResponsibleUser"`
	AssignedUser       int       `json:"AssignedUser"`
	CreatedBy          int       `json:"CreatedBy"`
	Stage              Stage     `json:"Stage"`
	CancellationReason string    `json:"CancellationReason"`
	Deadline           time.Time `json:"Deadline"`
	Message            string    `json:"Message"`
	Description        string    `json:"Description"`
	Priority           Priority  `json:"Priority"`
}
