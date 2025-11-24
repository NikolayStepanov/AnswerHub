package qahandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service/answer"
	"github.com/NikolayStepanov/AnswerHub/internal/service/question"
)

type Handler struct {
	questionService question.Service
	answerService   answer.Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetQuestionWithAnswers(w http.ResponseWriter, r *http.Request) {
	var (
		idStr            string
		questionID       int64
		questionDTO      dto.QuestionDTO
		answersDTO       []dto.AnswerDTO
		answersResponses []dto.AnswerResponseDTO
		err              error
	)

	idStr = r.PathValue("id")
	questionID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	questionDTO, err = h.questionService.Get(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to get question", http.StatusInternalServerError)
		return
	}

	answersDTO, err = h.answerService.GetAnswersByQuestionID(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to get question", http.StatusInternalServerError)
		return
	}
	answersResponses = dto.ToAnswerResponseList(answersDTO)
	jsonResponse, err := json.Marshal(&getQuestionResponse{questionDTO, answersResponses})
	if err != nil {
		http.Error(w, "Failed to get question", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) DeleteQuestionWithAnswers(w http.ResponseWriter, r *http.Request) {
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
	err = h.answerService.DeleteAnswersByQuestionID(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
	}

	err = h.questionService.Delete(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
