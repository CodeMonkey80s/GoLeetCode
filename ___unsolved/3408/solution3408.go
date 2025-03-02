package solution3408

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []*task

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// if pq[i].priority == pq[j].priority {
	// 	return pq[i].userId < pq[j].userId
	// }
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	t := x.(*task)
	t.index = n
	*pq = append(*pq, t)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	t := old[n-1]
	old[n-1] = nil
	t.index = -1
	*pq = old[0 : n-1]
	return t
}

func (pq *PriorityQueue) Update(t *task, p int) {
	t.priority = p
	heap.Fix(pq, t.index)
}

type TaskManager struct {
	tasksById       map[int]*task
	tasksByPriority PriorityQueue
}

type task struct {
	taskId   int
	userId   int
	priority int
	index    int
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		tasksById:       make(map[int]*task, len(tasks)),
		tasksByPriority: make([]*task, 0, len(tasks)),
	}
	for i := 0; i < len(tasks); i++ {
		t := task{
			taskId:   tasks[i][1],
			userId:   tasks[i][0],
			priority: tasks[i][2],
			index:    i,
		}
		tm.tasksById[t.taskId] = &t
		tm.tasksByPriority = append(tm.tasksByPriority, &t)
	}
	heap.Init(&tm.tasksByPriority)
	for _, t := range tm.tasksByPriority {
		fmt.Printf("r: %+v\n", t)
	}
	fmt.Println()
	return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	t := task{
		taskId:   taskId,
		userId:   userId,
		priority: priority,
		index:    len(tm.tasksByPriority) + 1,
	}
	tm.tasksById[t.taskId] = &t
	tm.tasksByPriority = append(tm.tasksByPriority, &t)
	heap.Push(&tm.tasksByPriority, &t)
	// heap.Fix(&tm.tasksByPriority, t.index)
	// tm.tasksByPriority.Update(&t, t.index)
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	tm.tasksById[taskId].priority = newPriority
	// heap.Fix(&tm.tasksByPriority, tm.tasksById[taskId].index)
}

func (tm *TaskManager) Rmv(taskId int) {
	// delete(tm.tasksById, taskId)
	// heap.Pop(&tm.tasksByPriority)
}

func (tm *TaskManager) ExecTop() int {
	// return tm.tasksByPriority.Pop().(*task).userId
	return 0
}
