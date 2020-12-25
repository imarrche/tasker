package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func (s *Server) taskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		ts, err := s.service.Tasks().GetByColumnID(columnID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusOK, ts)
	}
}

func (s *Server) taskCreate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		t := model.Task{
			Name:        request.Name,
			Description: request.Description,
			ColumnID:    columnID,
		}
		t, err = s.service.Tasks().Create(t)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, t)
	}
}

func (s *Server) taskDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		t, err := s.service.Tasks().GetByID(taskID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusOK, t)
	}
}

func (s *Server) taskMoveX() http.HandlerFunc {
	type request struct {
		Left bool `json:"left"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		err = s.service.Tasks().MoveToColumnByID(taskID, request.Left)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *Server) taskMoveY() http.HandlerFunc {
	type request struct {
		Up bool `json:"up"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		err = s.service.Tasks().MoveByID(taskID, request.Up)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *Server) taskUpdate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		t := model.Task{
			ID:          taskID,
			Name:        request.Name,
			Description: request.Description,
		}
		t, err = s.service.Tasks().Update(t)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, nil)
		}

		s.respond(w, r, http.StatusOK, t)
	}
}

func (s *Server) taskDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		err = s.service.Tasks().DeleteByID(taskID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusNoContent, nil)
	}
}
