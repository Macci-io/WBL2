package calendar

import (
	"sync"
	"time"
)

// Event наша целевая структура для хранения единицы информации
type Event struct {
	sync.Mutex
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// NewEvent создаем пустой event
func NewEvent() *Event {
	return new(Event)
}
