package main

import (
	"fmt"
	conifgs "go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/stat"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/event"
	"go/adv-demo/pkg/middleware"

	"net/http"
)

func App() http.Handler {

	conf := conifgs.LoadConfig()
	router := http.NewServeMux()
	db := db.NewDb(conf)
	evenbus := event.NewEventBus()
	// repo
	linkRepo := link.NewLinkRepo(db)
	userRepo := user.NewUserRepo(db)
	statRepo := stat.NewStatRepo(db)
	//service
	authService := auth.NewAuthService(userRepo)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		Event:    evenbus,
		StatRepo: statRepo,
	})
	//handler
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepo,
		Config:         conf,
		Event:          evenbus,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	stat.NewStatHandler(
		router, stat.StatHandlerDeps{
			StatRepo: statRepo,
			Config:   conf,
		})

	go statService.AddClic()
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	return stack(router)
}
func main() {

	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is work")
	server.ListenAndServe()

}
