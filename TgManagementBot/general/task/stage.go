package task

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
