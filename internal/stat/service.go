package stat

import (
	"go/adv-demo/pkg/event"
	"log"
)

type StatServiceDeps struct {
	Event    *event.EventBus
	StatRepo *StatRepo
}

type StatService struct {
	Event    *event.EventBus
	StatRepo *StatRepo
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		Event:    deps.Event,
		StatRepo: deps.StatRepo,
	}
}

func (s *StatService) AddClic() {
	for msg := range s.Event.Subscribe() {
		if msg.Type == event.LinkVisitEvent {
			id, ok := msg.Data.(uint)

			if !ok {
				log.Fatal("Bad id ")
				continue
			}

			s.StatRepo.AddClic(id)
		}
	}
}
