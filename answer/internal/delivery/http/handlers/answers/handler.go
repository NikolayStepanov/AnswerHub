package answers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
)

type Handler struct {
	answerService service.Answer
}

func NewHandler(answer service.Answer) *Handler {
	return &Handler{answerService: answer}
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	var (
		idStr     string
		answerID  int64
		err       error
		answerDTO dto.AnswerDTO
	)

	idStr = r.PathValue("id")
	answerID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}
	answerDTO, err = h.answerService.Get(r.Context(), answerID)
	if err != nil {
		http.Error(w, "Failed get answer", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(&getAnswerResponse{answerDTO})
	if err != nil {
		http.Error(w, "Failed get answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	var (
		idStr    string
		answerID int64
		err      error
	)

	idStr = r.PathValue("id")
	answerID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid answerID", http.StatusBadRequest)
		return
	}
	err = h.answerService.Delete(r.Context(), answerID)
	if err != nil {
		http.Error(w, "Failed delete answer", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
