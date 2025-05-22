package utils

import (
	"log"
	"net/http"
	"strconv"

	http2 "github.com/pizzament/rsc-test/pkg/http"
)

// ParseBannerID содержит общие функции парсинга и проверки banner_id, при некорректных значениях отправляется статус 400.
func ParseBannerID(w http.ResponseWriter, r *http.Request) (int16, bool) {
	bannerIDRaw := r.PathValue("banner_id")
	bannerID, err := strconv.ParseInt(bannerIDRaw, 10, 16)

	if err != nil {
		if err = http2.ErrorResponse(w, http.StatusBadRequest, "bannerID must be valid"); err != nil {
			log.Println("json.Encode failed ", err)

			return 0, false
		}

		return 0, false
	}

	if bannerID < 0 {
		if err = http2.ErrorResponse(w, http.StatusBadRequest, "bannerID must be more than zero"); err != nil {
			log.Println("json.Encode failed ", err)

			return 0, false
		}

		return 0, false
	}

	return int16(bannerID), true
}
