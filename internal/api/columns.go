package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func (s *Server) handleColumnList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		cs, err := s.service.Columns().GetByProjectID(projectID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusOK, cs)
	}
}

func (s *Server) handleColumnCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := model.Column{
			Name:      request.Name,
			ProjectID: projectID,
		}
		c, err = s.service.Columns().Create(c)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, c)
	}
}

func (s *Server) handleColumnDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		c, err := s.service.Columns().GetByID(columnID)
		if c.ProjectID != projectID {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *Server) handleColumnMove() http.HandlerFunc {
	type request struct {
		Left bool `json:"left"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		c, err := s.service.Columns().GetByID(columnID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}
		if c.ProjectID != projectID {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}

		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.service.Columns().MoveByID(columnID, request.Left)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *Server) handleColumnUpdate() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c, err := s.service.Columns().GetByID(columnID)
		if c.ProjectID != projectID {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		c.Name = request.Name
		c, err = s.service.Columns().Update(c)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *Server) handleColumnDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.Atoi(mux.Vars(r)["project_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}
		columnID, err := strconv.Atoi(mux.Vars(r)["column_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		c, err := s.service.Columns().GetByID(columnID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if c.ProjectID != projectID {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err = s.service.Columns().DeleteByID(c.ID); err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusNoContent, nil)
	}
}
