package server

import (
	"httpServer/pkg/calendar"
	"log"
	"net/http"
	"time"
)

// Server наш главный сервер
type Server struct {
	events  *calendar.EventsManager
	routing *http.ServeMux
}

// NewServer создает новый сервер
func NewServer() *Server {
	return &Server{calendar.NewEventsManager(), http.NewServeMux()}
}

// Logger MiddleWare-функция для логирования запросов
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

// Run Запускаем сервер с дефолтными параметрами localhost:8080
func (s *Server) Run() {
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.routing,
	}
	s.routing.HandleFunc("/create_event", Logger(s.CreateEvent))
	s.routing.HandleFunc("/update_event", Logger(s.UpdateEvent))
	s.routing.HandleFunc("/delete_event", Logger(s.DeleteEvent))
	s.routing.HandleFunc("/events_for_day", Logger(s.EventsForDay))
	s.routing.HandleFunc("/events_for_week", Logger(s.EventsForWeek))
	s.routing.HandleFunc("/events_for_month", Logger(s.EventsForMonth))
	log.Println("start server")
	log.Fatal(srv.ListenAndServe())
}
