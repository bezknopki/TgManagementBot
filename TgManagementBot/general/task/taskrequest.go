package task

type TaskRequest struct {
	TaskID        string `json:"task_id"`
	CorrelationID string `json:"correlation_id"`
}
