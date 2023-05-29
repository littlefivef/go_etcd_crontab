package common

const (
	// JOB_SAVE_DIR 任务保存目录
	JOB_SAVE_DIR = "/cron/jobs/"

	//任务强杀目录
	JOB_KILLER_DIR = "/cron/killer/"

	//任务锁目录
	JOB_LOCK_DIR = "/cron/lock/"

	服务注册地址
	JOB_WORKER_DIR = "/cron/worker/"

	//保存任务事件
	JOB_EVENT_SAVE = 1

	//删除任务事件
	JOB_EVENT_DELETE = 2

	//强杀任务事件
	JOB_EVENT_KILL = 3
)
