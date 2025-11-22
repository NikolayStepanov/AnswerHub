package handlers

import "net/http"

const (
	questionsPath = "/questions"
	answersPath   = "/answers"
)

func (h *Handler) setupRoutes() {
	h.setupQuestionsRoutes()
	h.setupAnswersRoutes()
	h.setupAnswersNestedRoutes()
}

func (h *Handler) setupQuestionsRoutes() {
	questionsRouter := h.Questions.RegisterRoutes()
	h.router.Handle(questionsPath+"/", http.StripPrefix(questionsPath, questionsRouter))
}

func (h *Handler) setupAnswersRoutes() {
	answersRouter := h.Answers.RegisterRoutes()
	h.router.Handle(answersPath+"/", http.StripPrefix(answersPath, answersRouter))
}

func (h *Handler) setupAnswersNestedRoutes() {
	answersNestedRouter := h.Answers.RegisterNestedRoutes()
	h.router.Handle(questionsPath+"/{id}"+answersPath+"/", http.StripPrefix(questionsPath, answersNestedRouter))
}
