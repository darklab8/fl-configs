package exe_mapped

import (
	"github.com/darklab8/darklab_flconfigs/flconfigs/settings/logger"
	"github.com/darklab8/darklab_goutils/goutils/worker"
	"github.com/darklab8/darklab_goutils/goutils/worker/worker_logger"
	"github.com/darklab8/darklab_goutils/goutils/worker/worker_types"
)

func launchWorker(worker_id worker_types.WorkerID, tasks <-chan worker.ITask) {
	logger.Log.Info("worker started", worker_logger.WorkerID(worker_id))
	for task := range tasks {

		task_err := make(chan error, 1)
		go func() {
			task_err <- task.RunTask(worker_id)
		}()

		<-task_err

		task.SetAsDone()
	}
	logger.Log.Info("worker finished", worker_logger.WorkerID(worker_id))
}
