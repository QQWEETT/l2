package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"q11/internal/app/model"
	"q11/internal/app/store"
	"strconv"
)

const (
	sessionName = "yes"
)

var (
	errQueryParamNotProvided = errors.New("should provided 'date' as YYYY-MM-DD")
)

type server struct {
	router       *http.ServeMux
	logger       *log.Logger
	store        store.Store
	config       *Config
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       http.NewServeMux(),
		store:        store,
		sessionStore: sessionStore,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/create_event", s.handleEventCreate())
	s.router.HandleFunc("/update_event", s.update())
	s.router.HandleFunc("/delete_event", s.delete())
	s.router.HandleFunc("/events_for_day", s.getForDay)
	s.router.HandleFunc("/events_for_week", s.getForWeek)
	s.router.HandleFunc("/events_for_month", s.getForMonth)

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

//----------------------------------------------------------
func (s *server) handleEventCreate() http.HandlerFunc {
	type request struct {
		UID  int    `json:"user_id"`
		Name string `json:"name"`
		Date string `json:"date"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.Event{
			UID:  req.UID,
			Name: req.Name,
			Date: req.Date,
		}
		if err := s.store.Event().Create(u); err != nil {
			s.error(w, r, 503, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

//------

func (s *server) getForDay(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	date := r.URL.Query().Get("date")
	fmt.Println(id)
	fmt.Println(date)
	u, err := s.store.Event().FindByDate(id, date)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
		return
	}
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	session.Values["date"] = u.Date
	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, u)
}

func (s *server) getForWeek(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	date := r.URL.Query().Get("date")
	fmt.Println(id)
	fmt.Println(date)
	u, err := s.store.Event().FindDateByInterval(id, date, 7)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
		return
	}
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, u)
}

func (s *server) getForMonth(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	date := r.URL.Query().Get("date")
	fmt.Println(id)
	fmt.Println(date)
	u, err := s.store.Event().FindDateByInterval(id, date, 31)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
		return
	}
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, u)
}
func (s *server) update() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
		UID  int    `json:"user_id"`
		Date string `json:"date"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.Event{
			Name: req.Name,
			UID:  req.UID,
			Date: req.Date,
		}
		if err := s.store.Event().UpdateEvent(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println(u)
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) delete() http.HandlerFunc {
	type request struct {
		UID  int    `json:"user_id"`
		Date string `json:"date"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.Event{
			UID:  req.UID,
			Date: req.Date,
		}
		if err := s.store.Event().DeleteEvent(u); err != nil {
			s.error(w, r, 503, err)
			return
		}
		fmt.Println(u)
		s.respond(w, r, http.StatusCreated, "Deleted")
	}
}
