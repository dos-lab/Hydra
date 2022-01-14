package types

import "time"

type SchedulerRecord struct {
	DoScheduleRecords []*DoScheduleOnceRecord
	// Extra 每个调度器独特的信息，如kMeans调度器可记录它内部kMeans执行的过程，BranchAndBound的参数，执行速度等。
	Extra interface{}
}

type DoScheduleOnceRecord struct {
	Duration time.Duration
}

type Record struct {
	Scheduler Scheduler
	Cluster Cluster
	FinishedJobs []Job
	SchedulerRecord SchedulerRecord
}