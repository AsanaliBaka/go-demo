package link

import (
	"fmt"
	configs "go/adv-demo/configs"
	"go/adv-demo/pkg/event"
	"go/adv-demo/pkg/middleware"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepo
	Config         *configs.Config
	Event          *event.EventBus
}
type LinkHandler struct {
	LinkRepository *LinkRepo
	Event          *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		Event:          deps.Event,
	}

	router.HandleFunc("POST /link", handler.CreateLinkHandler())
	router.HandleFunc("GET /link/{hash}", handler.GoTo())
	router.HandleFunc("DELETE /link/{id}", handler.DeleteLinkHandler())
	router.Handle("PUT /link/{id}", middleware.GetToken(handler.UpdateLinkHandler(), deps.Config))
	router.HandleFunc("GET /link", handler.GetAll())

}

func (l *LinkHandler) CreateLinkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)

		if err != nil {
			return
		}

		link := NewLink(body.Url)

		for {

			linkNew, _ := l.LinkRepository.Get(link.Hash)

			if linkNew == nil {
				break
			}

			link.Hash = RandStringsRunes(6)

		}
		createdlink, err := l.LinkRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JsonWriter(w, createdlink, 201)

	}

}
func (l *LinkHandler) DeleteLinkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		idint, err := strconv.ParseUint(idstring, 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		_, err = l.LinkRepository.GetById(uint(idint))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = l.LinkRepository.Delete(uint(idint))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JsonWriter(w, "as", 200)
	}
}
func (l *LinkHandler) UpdateLinkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}

		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		idstring := r.PathValue("id")
		idint, err := strconv.ParseUint(idstring, 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		link, err := l.LinkRepository.Update(&Link{
			Model: gorm.Model{
				ID: uint(idint),
			},
			Url:  body.Url,
			Hash: body.Hash,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}

		res.JsonWriter(w, link, 201)

	}
}
func (l *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := l.LinkRepository.Get(hash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		go l.Event.Publish(event.Event{
			Type: event.LinkVisitEvent,
			Data: link.ID,
		})

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}

func (l *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "invalid offset", http.StatusBadRequest)
		}

		links := l.LinkRepository.GetLinks(uint(limit), uint(offset))

		counter := l.LinkRepository.Counter()

		data := LinkAnswer{
			Links:   links,
			Counter: int(counter),
		}

		res.JsonWriter(w, data, 200)
	}
}
