package httpapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/squacky/openchat/config"
	"github.com/squacky/openchat/internal/user"
	"github.com/squacky/openchat/pkg/logger"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type HttpAPI struct {
	Host string
	Port int

	userService user.Service
	chatUser    map[string]*websocket.Conn
}

func (api *HttpAPI) GetAddr() string {
	return fmt.Sprintf("%s:%d", api.Host, api.Port)
}

func (api *HttpAPI) HandleErr(w http.ResponseWriter, statuCode int, msg string) {
	w.WriteHeader(statuCode)
	w.Write([]byte(msg))
}

func (api *HttpAPI) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})

	r.Get("/chat", func(w http.ResponseWriter, r *http.Request) {

		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			logger.Error("failed to accept websocket connection", zap.Error(err))
			return
		}

		api.chatUser[from] = c
		fmt.Println("sending message to", to)
		ctx := r.Context()
		for {

			msgType, p, err := c.Read(ctx)
			if err != nil {
				log.Println("Error reading from WebSocket:", err)
				return
			}

			targetClient, ok := api.chatUser[to]
			if ok && targetClient != nil {
				fmt.Println("sending message to", to)
				targetClient.Write(ctx, msgType, p)
			}

			// err = c.Write(ctx, msgType, p)
			// if err != nil {
			// 	log.Println("Error writing to WebSocket:", err)
			// 	return
			// }
		}

	})
	r.Post("/users", api.CreateUser)
	http.ListenAndServe(api.GetAddr(), r)
}

func NewHttpAPI(cfg *config.HttpConfig, userService user.Service) *HttpAPI {
	return &HttpAPI{
		Host:        cfg.Host,
		Port:        cfg.Port,
		userService: userService,
		chatUser:    make(map[string]*websocket.Conn),
	}
}
