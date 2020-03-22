//package worker implements worker who can do some kind of jobs
package worker

import (
	"context"
	"generator/backend-go/tools"
)

//Worker can register / deregister jobs and do them all later
type Worker struct {
	jobList []tools.Job
}

//New returns new Worker instance
func New() *Worker {
	return &Worker{jobList: []tools.Job{}}
}

//RegisterJob adds job to pending list
func (w *Worker) RegisterJob(j tools.Job) {
	w.jobList = append(w.jobList, j)
}

//DeregisterJob removes job from pending list
func (w *Worker) DeregisterJob(j tools.Job) {
	w.jobList = removeFromJobList(w.jobList, j)
}

//DoAll executes all jobs from current job list
//
//All jobs are executed concurrently. If one fails, rest of them is cancelled immediately.
func (w *Worker) DoAll() error {
	var err error
	ch1 := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer w.clearJobList()

	//Concurrently executing jobs
	for _, job := range w.jobList {
		go job.Execute(ctx, ch1)
	}

	//Stopping works if at least one fail
	for i := 0; i < len(w.jobList); i++ {
		err = <-ch1

		if err != nil {
			cancel()
			return err
		}
	}

	return nil
}

//clearJobList clears pending jobs list
func (w *Worker) clearJobList() {
	w.jobList = []tools.Job{}
}

//removeFromJobList helps removing job from job list
func removeFromJobList(jobList []tools.Job, job tools.Job) []tools.Job {
	newJobList := []tools.Job{}
	for _, currentJob := range jobList {
		if currentJob != job {
			newJobList = append(newJobList, currentJob)
		}
	}

	return newJobList
}
