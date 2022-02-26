package timer

import "time"

//定时任务
type Task struct {
	Expire   time.Time //绝对时间
	CallBack func()    //回调函数
}

//创建定时器
func NewTask(seconds time.Duration, callBack func()) *Task {
	task := new(Task)
	task.Expire = time.Now().Add(seconds * time.Second)
	task.CallBack = callBack

	return task
}
