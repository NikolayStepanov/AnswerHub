package handlers

import "net/http"

func (h *Handler) setupRoutes() {
	h.setupQuestionsRoutes()
	h.setupAnswersRoutes()
}

func (h *Handler) setupQuestionsRoutes() {
	questionsRouter := h.Questions.RegisterRoutes()
	h.router.Handle("/questions/", http.StripPrefix("/questions", questionsRouter))
}

func (h *Handler) setupAnswersRoutes() {
	answersRouter := h.Answers.RegisterRoutes()
	h.router.Handle("/answers/", http.StripPrefix("/answers", answersRouter))
}
