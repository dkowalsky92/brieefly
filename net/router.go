package net

import (
	"net/http"

	"github.com/dkowalsky/brieefly/log"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/database"
)

// Router - hub for networking
type Router struct {
	mux      *http.ServeMux
	config   *config.Config
	database *database.Database
}

// NewRouter - creates a new router
func NewRouter(db *database.Database, config *config.Config) *Router {
	mux := http.NewServeMux()
	return &Router{database: db, config: config, mux: mux}
}

func RegularCallback(w http.ResponseWriter, r *http.Request) {
	log.Info("Performing callback.")
}

func LogMiddleware() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Info("Performing middleware.")
			f(w, r)
		}
	}
}

// Run - starts the server
func (r *Router) Run() {
	path := config.MyPath(r.config)
	r.mux.HandleFunc("/", WithStack(RegularCallback, LogMiddleware()))
	err := http.ListenAndServeTLS(path, r.config.Server.Certificate, r.config.Server.Key, r.mux)
	log.Error(err)

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
