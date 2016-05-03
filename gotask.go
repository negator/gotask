package gotask

import (
	"sync"
)

type GoTaskGroup struct {
	tasks []*GoTask
}

func (g *GoTaskGroup) Append(t *GoTask) {
	g.tasks = append(g.tasks, t)
}
func (g *GoTaskGroup) AppendF(f func()) {
	g.Append(task(f))
}

func (g *GoTaskGroup) AwaitAll() {
	for _, f := range g.tasks {
		f.Await()
	}
}

type GoTask struct {
	wg *sync.WaitGroup
}

func task(f func()) *GoTask {
	var wg sync.WaitGroup
	wg.Add(1)
	r := GoTask{wg: &wg}
	go func() {
		f()
		wg.Done()
	}()
	return &r
}
