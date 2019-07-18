package goToolCommon

import (
	"context"
	"sync"
	"time"
)

type IObject interface {
	GetKey() string
	GetObject() interface{}
}

type IObjectManager interface {
	Register() chan<- IObject
	Unregister() chan<- string

	IsEmpty() bool
	Length() int
	GetIdList() []string
	HasId(id string) bool
	GetObject(id string) interface{}

	IsClosed() bool
	Close()
}

func NewObject(key string, o interface{}) IObject {
	return &object{
		key:    key,
		object: o,
	}
}

type object struct {
	key    string
	object interface{}
}

func (o *object) GetKey() string {
	return o.key
}

func (o *object) GetObject() interface{} {
	return o.object
}

func NewObjectManager() IObjectManager {
	ctx, cancel := context.WithCancel(context.Background())
	m := objectManager{
		list:         make(map[string]interface{}),
		chRegister:   make(chan IObject),
		chUnregister: make(chan string),

		ctx:    ctx,
		cancel: cancel,

		isClosed: false,
	}
	m.start()
	return &m
}

type objectManager struct {
	list         map[string]interface{}
	chRegister   chan IObject
	chUnregister chan string
	lock         sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc

	isClosed bool
}

func (m *objectManager) start() {
	go func() {
		for {
			select {
			case o := <-m.chRegister:
				m.register(o)
			case k := <-m.chUnregister:
				m.unregister(k)
			case <-m.ctx.Done():
				m.Close()
				return
			}
		}
	}()
}

func (m *objectManager) Register() chan<- IObject {
	return m.chRegister
}

func (m *objectManager) register(object IObject) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.list[object.GetKey()] = object.GetObject()
}

func (m *objectManager) Unregister() chan<- string {
	return m.chUnregister
}

func (m *objectManager) unregister(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.list, key)
}

func (m *objectManager) IsEmpty() bool {
	return len(m.list) == 0
}

func (m *objectManager) Length() int {
	return len(m.list)
}

func (m *objectManager) GetIdList() []string {
	l := make([]string, 0)
	for k := range m.list {
		l = append(l, k)
	}
	return l
}

func (m *objectManager) GetObject(id string) interface{} {
	o, ok := m.list[id]
	if ok {
		return o
	} else {
		return nil
	}
}

func (m *objectManager) Close() {
	if !m.isClosed {
		return
	}
	m.isClosed = true
	m.cancel()
	time.Sleep(time.Second)
	close(m.chRegister)
	close(m.chUnregister)
}

func (m *objectManager) IsClosed() bool {
	return m.isClosed
}

func (m *objectManager) HasId(id string) bool {
	for k := range m.list {
		if k == id {
			return true
		}
	}
	return false
}
