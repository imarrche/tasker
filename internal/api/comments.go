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

func (s *Server) commentList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cs, err := s.service.Comments().GetByTaskID(taskID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, cs)
		}
	}
}

func (s *Server) commentCreate() http.HandlerFunc {
	type request struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := model.Comment{Text: req.Text, TaskID: taskID}
		c, err = s.service.Comments().Create(c)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if web.IsValidationError(err) {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusCreated, c)
		}
	}
}

func (s *Server) commentDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c, err := s.service.Comments().GetByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, c)
		}
	}
}

func (s *Server) commentUpdate() http.HandlerFunc {
	type request struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var req request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := model.Comment{ID: id, Text: req.Text}
		c, err = s.service.Comments().Update(c)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if web.IsValidationError(err) {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusOK, c)
		}
	}
}

func (s *Server) commentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.service.Comments().DeleteByID(id)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
		} else if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		} else {
			s.respond(w, r, http.StatusNoContent, nil)
		}
	}
}
