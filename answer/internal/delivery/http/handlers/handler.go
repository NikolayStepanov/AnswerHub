package handlers

import (
	"net/http"

	"github.com/NikolayStepanov/AnswerHub/v2/internal/delivery/http/handlers/answers"
	"github.com/NikolayStepanov/AnswerHub/v2/internal/delivery/http/handlers/questions"
)

type Handler struct {
	router    *http.ServeMux
	Questions *questions.Handler
	Answers   *answers.Handler
}

func NewHandler() *Handler {
	handler := &Handler{
		router: http.NewServeMux(),
	}

	handler.Answers = answers.NewHandler()
	handler.Questions = questions.NewHandler()

	handler.setupRoutes()
	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) GetRouter() *http.ServeMux {
	return h.router
}
