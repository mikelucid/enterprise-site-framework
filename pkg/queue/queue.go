package queue

type Job struct {
	Name    string
	Payload []byte
}

type Dispatcher interface {
	Dispatch(job Job) error
}
