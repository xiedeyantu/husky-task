package core

// task status
const (
	TaskDistributed = "Distributed"
	TaskRunning     = "Running"
	TaskSuccess     = "Success"
	TaskFailed      = "Failed"
)

// executor type
const (
	ExecutorLeader = "leader"
	ExecutorWorker = "worker"
)

// dispatche status
const (
	Need    = "Need"
	NotNeed = "NotNeed"
)
