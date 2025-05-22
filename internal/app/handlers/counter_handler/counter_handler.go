package counter_handler

import (
	"context"
	"net/http"

	http2 "github.com/pizzament/rsc-test/pkg/http"

	"github.com/pizzament/rsc-test/internal/app/utils"
	"github.com/pizzament/rsc-test/internal/model"
)

type service interface {
	AddCount(ctx context.Context, bannerID model.BannerID) error
}

type CounterHandler struct {
	service service
}

func NewCounterHandler(service service) *CounterHandler {
	return &CounterHandler{service: service}
}

func (h CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// парсинг и проверка banner_id
	bannerID, ok := utils.ParseBannerID(w, r)
	if !ok {
		return
	}

	// вызов сервиса с banner_id
	err := h.service.AddCount(r.Context(), model.BannerID(bannerID))
	if err != nil {
		if err = http2.ErrorResponse(w, http.StatusInternalServerError, err.Error()); err != nil {
			return
		}

		return
	}

	// в ответ только статус 200
	w.WriteHeader(http.StatusOK)
}
