package answers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"go.uber.org/zap"
)

type Handler struct {
	answerService service.Answer
}

func NewHandler(answer service.Answer) *Handler {
	return &Handler{answerService: answer}
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	answerID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Warn("Invalid answerID", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}

	answer, err := h.answerService.Get(r.Context(), answerID)
	if err != nil {
		logger.Error("Failed to get answer", zap.Int64("answerID", answerID), zap.Error(err))
		http.Error(w, "Failed to get answer", http.StatusInternalServerError)
		return
	}

	response := ToAnswerResponse(answer)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	answerID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Warn("Invalid answerID in request", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}

	if err := h.answerService.Delete(r.Context(), answerID); err != nil {
		logger.Error("Failed to delete answer", zap.Int64("answerID", answerID), zap.Error(err))
		http.Error(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}
	logger.Info("Answer deleted successfully", zap.Int64("questionID", answerID))
	w.WriteHeader(http.StatusNoContent)
}
