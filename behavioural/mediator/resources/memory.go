package resources

import (
	"fmt"
	"sync"
)

type MemorySizeInMB int64

type Memory struct {
	waitQueue []Application
	total     MemorySizeInMB
	used      MemorySizeInMB
	runQueue  []Application
	mutex     sync.Mutex
}

var once sync.Once
var memoryInstance *Memory

func GetMemory() *Memory {
	if memoryInstance == nil {
		once.Do(func() {
			memoryInstance = &Memory{
				waitQueue: []Application{},
				total:     MemorySizeInMB(1024 * 4),
				used:      MemorySizeInMB(0),
				runQueue:  []Application{},
				mutex:     sync.Mutex{},
			}
		})
	}
	return memoryInstance
}

func GetAppNameList(appList []Application) []string {
	res := []string{}
	for _, app := range appList {
		res = append(res, app.GetAppName())
	}
	return res
}

func (m *Memory) IsAvailable(app Application) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if app.MemoryRequired()+m.used > m.total {
		m.waitQueue = append(m.waitQueue, app)
		return false
	} else {
		m.runQueue = append(m.runQueue, app)
		m.used = m.used + app.MemoryRequired()
		app.GotMemory()
		return true
	}
}

func (m *Memory) NotifyFree(app Application) {
	fmt.Printf("Done: %s\n", app.GetAppName())
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// free memory used by running app
	m.FreeMemoryUsed(app)
	if len(m.waitQueue) > 0 {
		// add app from wait queue to run queue
		m.CheckWaitQueue()
	}
}

func (m *Memory) FreeMemoryUsed(app Application) {
	length := len(m.runQueue)
	for i, appRunning := range m.runQueue {
		if (appRunning.GetAppName()) == app.GetAppName() {
			m.runQueue[length-1], m.runQueue[i] = m.runQueue[i], m.runQueue[length-1]
			m.runQueue = m.runQueue[:length-1]
			m.used = m.used - app.MemoryRequired()
			return
		}
	}
}

func (m *Memory) CheckWaitQueue() {
	firstApp := m.waitQueue[0]
	if !(firstApp.MemoryRequired()+m.used > m.total) {
		m.runQueue = append(m.runQueue, firstApp)
		m.waitQueue = m.waitQueue[1:]
		m.used = m.used + firstApp.MemoryRequired()
		firstApp.GotMemory()
	}
}
