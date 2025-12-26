package handlers

const (
	answersPath = "/answers"
)

func (h *Handler) setupRoutes() {
	h.setupQuestionsRoutes()
	h.setupAnswersRoutes()
}

func (h *Handler) setupQuestionsRoutes() {
	h.router.HandleFunc("GET /questions/", h.Questions.GetQuestions)
	h.router.HandleFunc("POST /questions/", h.Questions.CreateQuestion)

	h.router.HandleFunc("GET /questions/{id}", h.QAHandler.GetQuestionWithAnswers)
	h.router.HandleFunc("DELETE /questions/{id}", h.Questions.Delete)

	h.router.HandleFunc("POST /questions/{id}/answers/", h.QAHandler.CreateAnswer)
}

func (h *Handler) setupAnswersRoutes() {
	answersRouter := h.Answers.RegisterRoutes()
	h.router.Handle("/answers/", answersRouter)
}
