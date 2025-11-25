package qa

import "net/http"

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodGet+" "+"/{id}", h.GetQuestionWithAnswers)
	r.HandleFunc(http.MethodPost+" "+"/{id}", h.CreateAnswer)
	return r
}
