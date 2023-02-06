package main

import "fmt"

type Task struct {
	name        string
	description string
	complete    bool
}

type TaskList struct {
	tasks []*Task
}

func (t *TaskList) addTask(task *Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) printTasks() {
	for _, task := range t.tasks {
		fmt.Println(task.name, task.description, task.complete)
	}
}

func (t *TaskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...) // first part of the slice without the index + second part of the slice without the index
}

func (t *Task) toPrint() {
	fmt.Println(t.name, t.description, t.complete)
}

func (t *Task) toComplete() {
	t.complete = true
}

func main() {
	task1 := Task{name: "Task 1", description: "This is the task 1", complete: false}
	task1.toPrint()
	task1.toComplete()
	task1.toPrint()
	task2 := Task{name: "Task 2", description: "This is the task 2", complete: false}
	// add tasks
	taskList := TaskList{}
	taskList.addTask(&task1) // we use & because we are passing a pointer
	taskList.addTask(&task2)
	fmt.Println(taskList)

	// print tasks
	taskList.printTasks()
	// delete task
	taskList.deleteTask(0)

	// print tasks
	taskList.printTasks()
}
