package answers

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

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "answerID: %d", id)
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete answerID: %d", id)
}

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	idQuestionStr := r.PathValue("id")
	id, err := strconv.Atoi(idQuestionStr)
	if err != nil {
		http.Error(w, "Invalid questionID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Create questionID: %d", id)
}
