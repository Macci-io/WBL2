package server

import (
	"httpServer/pkg/calendar"
	"time"
)

// EventModel промежуточная модель для временного сохранения данных с клиента
type EventModel struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Validate валидация ключевых параметров
func (e EventModel) Validate() bool {
	if _, ok := time.Parse("2006-01-02", e.Date); ok != nil {
		return false
	}
	return !(e.UserID == "" || e.Title == "")
}

// ToEvent конвертация в целевую структуру
func (e *EventModel) ToEvent() *calendar.Event {
	event := calendar.NewEvent()
	event.ID = e.ID
	event.Title = e.Title
	event.UserID = e.UserID
	event.Description = e.Description
	event.Date, _ = time.Parse("2006-01-02", e.Date)
	return event
}
