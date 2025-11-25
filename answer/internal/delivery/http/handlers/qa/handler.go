package qa

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
)

type Handler struct {
	qa service.QA
}

func NewHandler(qa service.QA) *Handler {
	return &Handler{qa: qa}
}

func (h *Handler) GetQuestionWithAnswers(w http.ResponseWriter, r *http.Request) {
	var (
		idStr               string
		questionID          int64
		questionResponseDTO dto.QuestionResponseDTO
		err                 error
	)

	idStr = r.PathValue("id")
	questionID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}
	questionResponseDTO, err = h.qa.GetQuestionWithAnswers(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed to get question", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(&getQuestionResponse{questionResponseDTO})
	if err != nil {
		http.Error(w, "Failed to get question", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	var (
		req     createAnswerRequest
		err     error
		baseDTO dto.BaseDTO
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Debug(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	baseDTO, err = h.qa.CreateAnswer(r.Context(), req.QuestionID, req.UserID, req.Text)
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(&createAnswerResponse{BaseDTO: baseDTO})
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
