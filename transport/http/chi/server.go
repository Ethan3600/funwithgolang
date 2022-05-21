package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ethan3600/funwithgolang/application"
	cpu_controller "github.com/Ethan3600/funwithgolang/controllers/cpu"
	"github.com/Ethan3600/funwithgolang/dtos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func NewServer(app application.AppContext) {
	r := chi.NewRouter()

	registerDefaultMiddleware(r)

	r.Get("/api/v1/cpu", func(w http.ResponseWriter, r *http.Request) {
		times := r.URL.Query().Get("times")

		intTimes, _ := strconv.Atoi(times)
		nums := cpu_controller.GetCpuIntensiveWork(intTimes)

		resp := dtos.R{
			"status": dtos.Success,
			"data":   nums,
		}

		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	})

	http.ListenAndServe(":1323", r)
}

func registerDefaultMiddleware(r *chi.Mux) {
	// e.Use(middleware.CORS())
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Mount("/", middleware.Profiler())
}
