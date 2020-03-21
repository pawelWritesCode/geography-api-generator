package tools

//Employee represents any employee who can register/deregister and do jobs.
//
//Part of observer design pattern
type Employee interface {

	//RegisterJob adds job to pending jobs list
	RegisterJob(j Job)

	//DeregisterJob removes job from pending jobs list
	DeregisterJob(j Job)

	//DoAll executes all jobs from pending list
	DoAll() error
}
