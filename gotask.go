package gotask

import (
	"reflect"
	"sync"
)

type GoTask struct {
	/*Make sure done is called only once*/
	once sync.Once
	done chan bool
}

/*Block until task is finished*/
func (t *GoTask) Await() {
	t.once.Do(func() {
		<-t.done
	})
}

func goTask(f func()) *GoTask {
	r := GoTask{done: make(chan bool, 1)}
	go func() {
		f()
		r.done <- true
	}()
	return &r
}

type GoTaskGroup struct {
	once  sync.Once
	tasks []*GoTask
}

/*Append returns a new GoTaskGroup with the set of tasks or function appended*/
func (g *GoTaskGroup) Append(t ...*GoTask) GoTaskGroup {
	tasks := append(g.tasks, t...)
	return GoTaskGroup{tasks: tasks}
}
func (g *GoTaskGroup) AppendF(f func()) GoTaskGroup {
	return g.Append(goTask(f))
}

func (g *GoTaskGroup) AwaitAll() {
	g.once.Do(func() {
		cases := make([]reflect.SelectCase, len(g.tasks))
		for i, task := range g.tasks {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(task.done)}
		}
		for r := len(cases); r > 0; r-- {
			reflect.Select(cases)
		}
	})
}
