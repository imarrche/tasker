package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func (s *Server) handleProjectList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ps []model.Project
		ps, err := s.service.Projects().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, ps)
	}
}

func (s *Server) handleProjectCreate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		p := model.Project{
			Name:        request.Name,
			Description: request.Description,
		}
		p, err := s.service.Projects().Create(p)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, p)
	}
}

func (s *Server) handleProjectDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		p, err := s.service.Projects().GetByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *Server) handleProjectUpdate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		p := model.Project{
			ID:          id,
			Name:        request.Name,
			Description: request.Description,
		}
		p, err = s.service.Projects().Update(p)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *Server) handleProjectDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		err = s.service.Projects().DeleteByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusNoContent, nil)
	}
}
