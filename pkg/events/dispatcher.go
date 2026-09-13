package events

import "sync"

type Handler func(payload any)

type Dispatcher struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

func NewDispatcher() *Dispatcher { return &Dispatcher{handlers: map[string][]Handler{}} }

func (d *Dispatcher) Subscribe(event string, h Handler) {
	d.mu.Lock()
	d.handlers[event] = append(d.handlers[event], h)
	d.mu.Unlock()
}

func (d *Dispatcher) Dispatch(event string, payload any) {
	d.mu.RLock()
	hs := d.handlers[event]
	d.mu.RUnlock()
	for _, h := range hs {
		h(payload)
	}
}
