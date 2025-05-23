package stats_handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pizzament/rsc-test/internal/app/utils"
	"github.com/pizzament/rsc-test/internal/model"
	http2 "github.com/pizzament/rsc-test/pkg/http"
)

type service interface {
	ReceiveStats(ctx context.Context, bannerID model.BannerID, from time.Time, to time.Time) ([]model.Stat, error)
}

type StatsHandler struct {
	service service
}

func NewStatsHandler(service service) *StatsHandler {
	return &StatsHandler{service: service}
}

func (h StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// парсинг и проверка banner_id
	bannerID, ok := utils.ParseBannerID(w, r)
	if !ok {
		return
	}

	// декодинг JSON
	var req StatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//  парсинг времени
	from, to, err := req.ParseTimes()
	if err != nil {
		http.Error(w, "Time parsing error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// проверка валидности данных
	if from.After(to) {
		http.Error(w, "'from' must be before 'to'", http.StatusBadRequest)
		return
	}

	// вызов сервиса с полученным параметром
	stats, err := h.service.ReceiveStats(r.Context(), model.BannerID(bannerID), from, to)
	if err != nil {
		if err = http2.ErrorResponse(w, http.StatusInternalServerError, err.Error()); err != nil {
			return
		}

		return
	}

	// формирование ответа
	response := StatsResponse{
		Stats: make([]StatItem, len(stats)),
	}

	for i, stat := range stats {
		response.Stats[i] = StatItem{
			TS: stat.Timestamp,
			V:  stat.Count,
		}
	}

	// отправка ответа
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println("response json.Encode failed")
		return
	}
}
