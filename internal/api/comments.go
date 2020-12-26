package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func (s *Server) commentList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
		}

		cs, err := s.service.Comments().GetByTaskID(taskID)
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

func (s *Server) commentCreate() http.HandlerFunc {
	type request struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := model.Comment{
			TaskID: taskID,
			Text:   request.Text,
		}
		if c, err = s.service.Comments().Create(c); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, c)
	}
}

func (s *Server) commentDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentID, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		c, err := s.service.Comments().GetByID(commentID)
		if err == store.ErrNotFound {
			s.error(w, r, http.StatusNotFound, nil)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *Server) commentUpdate() http.HandlerFunc {
	type request struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		commentID, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		request := &request{}
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		c := model.Comment{
			ID:   commentID,
			Text: request.Text,
		}
		if c, err = s.service.Comments().Update(c); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *Server) commentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentID, err := strconv.Atoi(mux.Vars(r)["comment_id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		err = s.service.Comments().DeleteByID(commentID)
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
