package questions

import "net/http"

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodGet+" "+"/", h.GetQuestions)
	r.HandleFunc(http.MethodPost+" "+"/", h.CreateQuestion)
	r.HandleFunc(http.MethodDelete+" "+"/{id}", h.DeleteQuestion)
	return r
}
