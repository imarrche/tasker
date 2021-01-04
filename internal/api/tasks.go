package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/service/web"
	"github.com/imarrche/tasker/internal/store"
)

func (s *Server) taskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		ts, err := s.service.Tasks().GetByColumnID(columnID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
		} else {
			s.respond(w, r, http.StatusOK, ts)
		}
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
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		t := model.Task{Name: req.Name, Description: req.Description, ColumnID: columnID}
		t, err = s.service.Tasks().Create(t)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if web.IsValidationError(err) {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusCreated, t)
		}
	}
}

func (s *Server) taskDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		t, err := s.service.Tasks().GetByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, t)
		}
	}
}

func (s *Server) taskMoveX() http.HandlerFunc {
	type request struct {
		Left bool `json:"left"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.service.Tasks().MoveToColumnByID(id, req.Left)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err == web.ErrInvalidMove {
			s.error(w, r, http.StatusBadRequest, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, nil)
		}
	}
}

func (s *Server) taskMoveY() http.HandlerFunc {
	type request struct {
		Up bool `json:"up"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.service.Tasks().MoveByID(id, req.Up)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err == web.ErrInvalidMove {
			s.error(w, r, http.StatusBadRequest, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, nil)
		}
	}
}

func (s *Server) taskUpdate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		t := model.Task{ID: id, Name: req.Name, Description: req.Description}
		t, err = s.service.Tasks().Update(t)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if web.IsValidationError(err) {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, t)
		}
	}
}

func (s *Server) taskDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.service.Tasks().DeleteByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusNoContent, nil)
		}
	}
}
