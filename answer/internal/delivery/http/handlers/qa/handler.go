package qa

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"go.uber.org/zap"
)

type Handler struct {
	qa service.QA
}

func NewHandler(qa service.QA) *Handler {
	return &Handler{qa: qa}
}

func (h *Handler) GetQuestionWithAnswers(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	questionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Warn("Invalid questionID", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}
	qa, err := h.qa.GetQuestionWithAnswers(r.Context(), questionID)
	if err != nil {
		logger.Error("Failed to get question with answers", zap.Int64(" questionID", questionID), zap.Error(err))
		http.Error(w, "Failed to get question with answers", http.StatusInternalServerError)
		return
	}

	response := ToQuestionResponse(qa)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	logger.Info("create new answer")
	var req CreateAnswerRequest
	idStr := r.PathValue("id")
	questionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Warn("Invalid questionID", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("Invalid request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	base, err := h.qa.CreateAnswer(r.Context(), questionID, req.UserID, req.Text)
	if err != nil {
		logger.Error("Failed to create answer", zap.Error(err))
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	response := ToBaseResponse(base)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
