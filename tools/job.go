package tools

import (
	"context"
)

//Job represents any job that can be executed
type Job interface {

	//Execute triggers execution of job
	//
	//Job is/is not executed according to current ctx.
	//execution of this method should be safe for concurrently execution
	Execute(ctx context.Context, ch1 chan error)
}
