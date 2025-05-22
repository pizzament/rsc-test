package stats_handler

import (
	"context"
	"net/http"

	"github.com/pizzament/rsc-test/internal/app/utils"
	"github.com/pizzament/rsc-test/internal/model"
	http2 "github.com/pizzament/rsc-test/pkg/http"
)

type service interface {
	ReceiveStats(ctx context.Context, bannerID model.BannerID) ([]model.Stat, error)
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

	// вызов сервиса с полученным параметром
	_, err := h.service.ReceiveStats(r.Context(), model.BannerID(bannerID))
	if err != nil {
		if err = http2.ErrorResponse(w, http.StatusPreconditionFailed, err.Error()); err != nil {
			return
		}

		return
	}

	// формирование ответа
}
