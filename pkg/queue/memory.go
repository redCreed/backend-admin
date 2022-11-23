package queue

import (
	"sync"
	"time"
)

type queue chan Messager

type Memory struct {
	syncMap sync.Map
	wait    sync.WaitGroup
	mutex   sync.RWMutex
	poolNum int
}

func NewMemory(num int) Queue {
	return &Memory{poolNum: num}
}

func (m *Memory) string() string {
	return "memory"
}

func (m *Memory) makeQueue() queue {
	if m.poolNum <= 0 {
		return make(queue)
	}

	return make(queue, m.poolNum)
}

func (m *Memory) Add(messager Messager) error {
	q := m.getQueue(messager.GetKey())

	//推送
	go func(m Messager, q queue) {
		q <- m
	}(messager, q)
	return nil
}

func (m *Memory) getQueue(key string) queue {
	var q queue
	value, ok := m.syncMap.Load(key)
	if !ok {
		q = m.makeQueue()
		m.syncMap.Store(key, q)
	} else {
		q = value.(queue)
	}
	return q
}

func (m *Memory) Register(key string, consumerFunc ConsumerFunc) {
	qu := m.getQueue(key)
	t := time.NewTicker(5 * time.Second)
	var err error
	maxLen := 50
	data := make([]Messager, 0, maxLen)
	go func(q queue, cf ConsumerFunc) {
		for {
			select {
			case msg := <-q:
				if len(data) < maxLen {
					data = append(data, msg)
					continue
				}
			case <-t.C:
				//定时时间到执行消费操作
				if len(data) == 0 {
					continue
				}
			}

			if err = cf(data); err != nil {
				for _, v := range data {
					if v.GetErrorCount() < 3 {
						v.SetErrorCount(v.GetErrorCount() + 1)
						q <- v
					}
				}
			}

			//重新初始化
			data = make([]Messager, 0, maxLen)
		}
	}(qu, consumerFunc)
	return
}

func (m *Memory) Run() {

}

func (m *Memory) Close() {

}
