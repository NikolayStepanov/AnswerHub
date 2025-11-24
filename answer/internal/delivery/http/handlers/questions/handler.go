package questions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service/question"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
)

type Handler struct {
	questionService question.Service
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
	var (
		idStr      string
		questionID int64
		err        error
	)

	idStr = r.PathValue("id")
	questionID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	err = h.questionService.Delete(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		req     createRequest
		baseDTO dto.BaseDTO
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Debug(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	baseDTO, err = h.questionService.Create(r.Context(), req.Text)
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(&createResponse{BaseDTO: baseDTO})
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
