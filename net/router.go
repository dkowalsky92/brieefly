package net

import (
	"net/http"
	"time"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/ctrl/agency"
	"github.com/dkowalsky/brieefly/ctrl/project"
	"github.com/dkowalsky/brieefly/ctrl/user"
	"github.com/dkowalsky/brieefly/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router - hub for networking
type Router struct {
	mux      *chi.Mux
	config   *config.Config
	database *db.DB
}

// BrieeflyRouter - creates a new router
func BrieeflyRouter(db *db.DB, config *config.Config) *Router {
	mux := chi.NewRouter()
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Mount("/api/projects", project.NewRouter(db).Mux)
	mux.Mount("/api/users", user.NewRouter(db).Mux)
	mux.Mount("/api/agencies", agency.NewRouter(db).Mux)

	return &Router{database: db, config: config, mux: mux}
}

// Run - starts the server
func (r *Router) Run() {
	path := config.MyPath(r.config)

	http.ListenAndServe(path, r.mux)
	// r.mux.HandleFunc("/", net.WithStack(RegularCallback, LogMiddleware()))
	// err := http.ListenAndServe(path, r.mux)
	// log.Error(err)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.

	// RESTy routes for "articles" resource
	// r.Route("/user", func(r chi.Router) {
	// 	r.With(paginate).Get("/", listArticles)                           // GET /articles
	// 	r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

	// 	r.Post("/", createArticle)       // POST /articles
	// 	r.Get("/search", searchArticles) // GET /articles/search

	// 	// Regexp url parameters:
	// 	r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug) // GET /articles/home-is-toronto

	// 	// Subrouters:
	// 	r.Route("/{articleID}", func(r chi.Router) {
	// 		r.Use(ArticleCtx)
	// 		r.Get("/", getArticle)       // GET /articles/123
	// 		r.Put("/", updateArticle)    // PUT /articles/123
	// 		r.Delete("/", deleteArticle) // DELETE /articles/123
	// 	})
	// })

	// Mount the admin sub-router

	// go r.cmh.Run()

	// r.muxRouter.HandleFunc("/account", r.GetAccountWithEmail).Methods("GET").Queries("email", "{email:.*}")
	// r.muxRouter.HandleFunc("/account/exist", r.GetAccountWithEmailAndPassword).Methods("GET").Queries("email", "{email:.*}", "password", "{password:.*}")
	// r.muxRouter.HandleFunc("/account", r.CreateAccount).Methods("POST")
	// r.muxRouter.HandleFunc("/account", r.UpdateAccount).Methods("PUT")
	// r.muxRouter.HandleFunc("/account/nearby", r.GetAccountsWithinRadius).Methods("GET").Queries("id", "{id:.*}", "radius", "{radius:.*}", "lat", "{lat:.*}", "lon", "{lon:.*}")
	// r.muxRouter.HandleFunc("/account/location", r.CreateOrUpdateAccountLocation).Methods("POST")
	// r.muxRouter.HandleFunc("/account/location", r.GetAccountLocationsWithinRadius).Methods("GET").Queries("id", "{id:.*}", "radius", "{radius:.*}", "lat", "{lat:.*}", "lon", "{lon:.*}")
	// r.muxRouter.HandleFunc("/account/image", r.CreateImage).Methods("POST")
	// r.muxRouter.HandleFunc("/account/image", r.GetImagesForAccountID).Methods("GET").Queries("id", "{id:.*}")
	// r.muxRouter.HandleFunc("/account/thumbnail", r.CreateThumbnail).Methods("POST")
	// r.muxRouter.HandleFunc("/account/thumbnail", r.GetThumbnailForAccountID).Methods("GET").Queries("id", "{id:.*}")
	// r.muxRouter.HandleFunc("/account/conversation", r.GetConversationWithMessages).Methods("POST")
	// r.muxRouter.HandleFunc("/ws/account/conversation/join", r.cmh.HandleConnection).Methods("GET").Queries("account_id", "{account_id:.*}", "conversation_id", "{conversation_id:.*}")

	// r.muxRouter.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage/"))))
	// r.muxRouter.PathPrefix("/policy/").Handler(http.StripPrefix("/policy/privacy_policy.html", http.FileServer(http.Dir("./policy/privacy_policy.html"))))

	// ip, err := config.MyPath()
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// connStr := ":" + r.conf.Server.Port
	// fmt.Println(connStr)
	// fmt.Println(ip)
	// err = http.ListenAndServeTLS(connStr, r.conf.Server.Certificate, r.conf.Server.Key, r.muxRouter)
	// if err != nil {
	// 	panic(err)
	// }
}
