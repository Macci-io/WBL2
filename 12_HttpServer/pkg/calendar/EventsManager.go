package calendar

import (
	"fmt"
	"sync"
	"time"
)

// EventsManager менеджер событий, по сути выполняет основную роль календаря
type EventsManager struct {
	*sync.RWMutex
	data map[string]*Event
}

// NewEventsManager создаем менеджера
func NewEventsManager() *EventsManager {
	return &EventsManager{&sync.RWMutex{}, make(map[string]*Event)}
}

// Add Добавляет событие в календарь
func (es *EventsManager) Add(e *Event) {
	es.Lock()
	es.data[e.ID] = e
	es.Unlock()
}

// DeleteEvent удаляет событие из календаря если таковой имеется
func (es *EventsManager) DeleteEvent(id string) (ok bool) {
	es.Lock()
	if _, ok = es.data[id]; ok {
		delete(es.data, id)
	}
	es.Unlock()
	return ok
}

// GetEvent находит событие по ID
func (es *EventsManager) GetEvent(id string) *Event {
	es.RLock()
	defer es.RUnlock()
	return es.data[id]
}

// UpdateEvent Обновляет событие если оно имеется
func (es *EventsManager) UpdateEvent(e *Event) (ok error) {
	curEv := es.GetEvent(e.ID)
	if curEv == nil {
		return fmt.Errorf("event id with: %s - not found", e.ID)
	}
	if e.Title != "" {
		curEv.Lock()
		curEv.Title = e.Title
		curEv.Unlock()
	}
	if !e.Date.IsZero() {
		curEv.Lock()
		curEv.Date = e.Date
		curEv.Unlock()
	}
	if e.Description != "" {
		curEv.Lock()
		curEv.Description = e.Description
		curEv.Unlock()
	}
	if e.UserID != "" {
		curEv.Lock()
		curEv.UserID = e.UserID
		curEv.Unlock()
	}
	return nil
}

// GetEvents получаем все события за промежуток времени
func (es *EventsManager) GetEvents(userID string, start, end time.Time) (ev []*Event) {
	es.RLock()
	for _, v := range es.data {
		if v.UserID == userID {
			if (v.Date == start || v.Date.After(start)) && v.Date.Before(end) {
				ev = append(ev, v)
			}
		}
	}
	es.RUnlock()
	return ev
}
