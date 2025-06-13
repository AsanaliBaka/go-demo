package stat

import (
	conifgs "go/adv-demo/configs"
	"go/adv-demo/pkg/middleware"
	"go/adv-demo/pkg/res"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepo *StatRepo
	Config   *conifgs.Config
}

type StatHandler struct {
	StatRepo *StatRepo
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepo: deps.StatRepo,
	}

	router.Handle("GET /stat", middleware.GetToken(handler.GetStat(), deps.Config))

}

func (s *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		from, err := time.Parse("2006-10-01", r.URL.Query().Get("from"))

		if err != nil {
			http.Error(w, "invalid from param", http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-10-01", r.URL.Query().Get("to"))

		if err != nil {
			http.Error(w, "invalid to param", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")

		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "invalid by param", http.StatusBadRequest)
			return
		}

		stats := s.StatRepo.GetStat(by, from, to)
		res.JsonWriter(w, stats, 200)

	}
}
