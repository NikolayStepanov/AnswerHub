package handlers

import (
	"net/http"

	"github.com/NikolayStepanov/AnswerHub/internal/delivery/http/handlers/answers"
	"github.com/NikolayStepanov/AnswerHub/internal/delivery/http/handlers/qa"
	"github.com/NikolayStepanov/AnswerHub/internal/delivery/http/handlers/questions"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
)

type Handler struct {
	router    *http.ServeMux
	Questions *questions.Handler
	Answers   *answers.Handler
	QAHandler *qa.Handler
}

func NewHandler(answer service.Answer, question service.Question, qaService service.QA) *Handler {
	handler := &Handler{
		router: http.NewServeMux(),
	}

	handler.Answers = answers.NewHandler(answer)
	handler.Questions = questions.NewHandler(question)
	handler.QAHandler = qa.NewHandler(qaService)

	handler.setupRoutes()
	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) GetRouter() *http.ServeMux {
	return h.router
}
