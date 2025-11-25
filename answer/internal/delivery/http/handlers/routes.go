package handlers

import "net/http"

const (
	questionsPath = "/questions"
	answersPath   = "/answers"
)

func (h *Handler) setupRoutes() {
	h.setupQuestionsRoutes()
	h.setupAnswersRoutes()
}

func (h *Handler) setupQuestionsRoutes() {
	questionsRouter := h.Questions.RegisterRoutes()

	questionsByIDeRouter := h.RegisterQuestionByIDRoutes()

	h.router.Handle(questionsPath+"/", http.StripPrefix(questionsPath, questionsRouter))
	h.router.Handle(questionsPath+"/{id}", http.StripPrefix(questionsPath, questionsByIDeRouter))
}

func (h *Handler) setupAnswersRoutes() {
	answersRouter := h.Answers.RegisterRoutes()
	h.router.Handle(answersPath+"/", http.StripPrefix(answersPath, answersRouter))
}

func (h *Handler) RegisterQuestionByIDRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodGet+" "+"/{id}", h.QAHandler.GetQuestionWithAnswers)
	r.HandleFunc(http.MethodPost+" "+"/{id}", h.QAHandler.CreateAnswer)
	r.HandleFunc(http.MethodDelete+" "+"/{id}", h.Questions.Delete)

	return r
}
