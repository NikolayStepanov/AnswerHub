package questions

import (
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "questions")
}

func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete questionID: %d", id)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "question created")
}

func (h *Handler) GetQuestionWithAnswers(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Question with answers questionID: %d", id)
}
