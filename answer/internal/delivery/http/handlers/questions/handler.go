package questions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
)

type Handler struct {
	questionService service.Question
}

func NewHandler(question service.Question) *Handler {
	return &Handler{questionService: question}
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		questions    []dto.QuestionDTO
		jsonResponse []byte
	)
	questions = make([]dto.QuestionDTO, 0)
	questions, err = h.questionService.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonResponse, err = json.Marshal(&List{questions})
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		req          createRequest
		baseDTO      dto.BaseDTO
		jsonResponse []byte
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

	jsonResponse, err = json.Marshal(&createResponse{BaseDTO: baseDTO})
	if err != nil {
		http.Error(w, "Invalid create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		idStr      string
		questionID int64
		err        error
	)

	idStr = r.PathValue("id")
	questionID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid questionID", http.StatusBadRequest)
		return
	}

	err = h.questionService.Delete(r.Context(), questionID)
	if err != nil {
		http.Error(w, "Failed delete question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
