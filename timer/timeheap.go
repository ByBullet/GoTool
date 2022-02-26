package timer

import (
	"time"
)

//时间堆
type Timer struct {
	arrays   []*Task //定时器(最小堆数据结构)
	capacity int     //堆容量
	size     int     //堆元素数量
}

//创建一个时间堆(空堆)
func NewTimer() *Timer {
	return new(Timer)
}

func (th *Timer) AddTaskCall(seconds time.Duration, callBack func()) {
	th.AddTask(NewTask(seconds, callBack))
}

//添加一个定时任务，指往堆中添加一个节点
//时间复杂度O(logN)，N为节点个数
func (th *Timer) AddTask(timer *Task) {
	if timer == nil {
		return
	}

	//将节点元素添加到最后一个叶子节点，若该节点到期时间小于其父节点则与其交换，执行此上滤操作
	if th.size < th.capacity {
		th.arrays[th.size] = timer
	} else {
		th.arrays = append(th.arrays, timer)
		th.capacity = len(th.arrays) //更新容量
	}
	th.size++

	//最后一个叶子节点下标
	curIndex := th.size - 1

	for curIndex > 0 && (curIndex-1)/2 >= 0 {
		curNode := th.arrays[curIndex]          //当前节点
		parentNode := th.arrays[(curIndex-1)/2] //父节点

		//当前节点的时间-父节点时间 < 0 则表示当前节点比父节点小，交换后继续上滤
		if curNode.Expire.Sub(parentNode.Expire) < 0 {
			th.arrays[curIndex] = parentNode
			th.arrays[(curIndex-1)/2] = curNode

			curIndex = (curIndex - 1) / 2
		} else {
			break
		}
	}
}

func (th *Timer) Empty() bool {
	if th.size <= 0 {
		return true
	}
	return false
}

//获得堆顶部定时器
func (th *Timer) Top() *Task {
	if th.Empty() {
		return nil
	}
	return th.arrays[0]
}

//删除定时器
func (th *Timer) Remove(task *Task) {
	if task == nil {
		return
	}

	//将回调函数置空，延时删除
	task.CallBack = nil
}

//删除根节点后需要保持堆序
func (th *Timer) Pop() {
	if th.Empty() {
		return
	}

	if th.arrays[0] != nil {
		th.arrays[0] = th.arrays[th.size-1] //将原来的堆顶元素替换为堆数组中的最后一个元素
		th.size--
		th.percolateDown(0) //对新的堆顶元素执行下滤操作，保证堆序
	}
}

//对从hole开始的元素进行下滤操作，保证堆序
func (th *Timer) percolateDown(hole int) {
	temp := th.arrays[hole]
	child := 0
	for hole*2+1 <= th.size-1 {
		child = hole*2 + 1
		if child < (th.size-1) &&
			(th.arrays[child+1].Expire.Sub(th.arrays[child].Expire) < 0) {
			child++
		}

		if th.arrays[child].Expire.Sub(temp.Expire) < 0 {
			th.arrays[hole] = th.arrays[child]
		} else {
			break
		}
		hole = child
	}
	th.arrays[hole] = temp
}

//心跳起搏函数，定时到期时执行
func (th *Timer) Tick() {
	topTask := th.arrays[0]
	curTime := time.Now()

	for !th.Empty() {
		if topTask == nil {
			break
		}

		//定时器未到期
		if topTask.Expire.Sub(curTime) > 0 {
			break
		}

		//执行定时到期任务
		if topTask.CallBack != nil {
			topTask.CallBack()
		}

		//将堆顶元素删除，同时生成新的堆顶定时器
		th.Pop()
		topTask = th.arrays[0]
	}
}

func (th *Timer) Start() {
	for !th.Empty() {
		th.Tick()
	}
}
