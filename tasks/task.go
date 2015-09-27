package tasks

type Task interface {
	Run() error
}

// Tasks implement the Resettable interface if they want to run any sort of reset logic.
// This includes things like clearing intermediate results.
type Resettable interface {
	/* If there was a retriable error, this will be called before Run */
	ResetForRetry()
}

// Tasks implement the Messy interface if they have side effects outside of the normal
// database transactions.  This includes things like touching files, etc.  CleanUp is
// always called exactly once, at the end of the task, regardless of success.
type Messy interface {
	CleanUp()
}