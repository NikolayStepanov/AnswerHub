package questions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"go.uber.org/zap"
)

type Handler struct {
	questionService service.Question
}

func NewHandler(question service.Question) *Handler {
	return &Handler{questionService: question}
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.questionService.List(r.Context())
	if err != nil {
		logger.Error("Failed to get questions", zap.Error(err))
		http.Error(w, "Failed to get questions", http.StatusBadRequest)
		return
	}

	response := ToQuestionList(questions)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req createRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("Failed to decode request", zap.Error(err))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	base, err := h.questionService.Create(r.Context(), req.Text)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusBadRequest)
		return
	}

	response := ToBaseResponse(base)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	questionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Warn("Invalid question ID", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	if err := h.questionService.Delete(r.Context(), questionID); err != nil {
		logger.Error("Failed to delete question", zap.Int64("questionID", questionID), zap.Error(err))
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	logger.Info("Question deleted successfully", zap.Int64("questionID", questionID))
	w.WriteHeader(http.StatusNoContent)
}
