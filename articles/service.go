package articles

import mesg "github.com/mesg-foundation/go-service"

// list of task keys:
const (
	createTaskKey = "create"
	getTaskKey    = "get"
	listTaskKey   = "list"
)

// listenTasks starts listening for service's tasks.
func (s *ArticlesService) listenTasks() error {
	return s.s.Listen(
		mesg.Task(createTaskKey, s.createHandler),
		mesg.Task(listTaskKey, s.listHandler),
		mesg.Task(getTaskKey, s.getHandler),
	)
}

// output key for errors.
const errOutputKey = "error"

// errorOutput is the error output data.
type errorOutput struct {
	Message string `json:"message"`
}

// newErrorOutput returns a new error output from given err.
func newErrorOutput(err error) (outputKey string, outputData mesg.Data) {
	return errOutputKey, errorOutput{Message: err.Error()}
}
