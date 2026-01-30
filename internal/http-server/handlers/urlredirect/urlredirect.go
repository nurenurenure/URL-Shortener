package urlredirect

import (
	"URLshort/internal/lib/api/response"
	"URLshort/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Request struct {
	Alias string `json:"alias"`
}

type Response struct {
	response.Response
	URL string `json:"url" validate:"required,url"`
}

type URLredirect interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlredirect URLredirect) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Error("alias is empty")
			http.Error(w, "alias is required", http.StatusBadRequest)
			return
		}
		log.Info("redirect request", slog.String("alias", alias))
		red_url, err := urlredirect.GetURL(alias)
		if err != nil {
			log.Error("failed to get URL", sl.Err(err), slog.String("alias", alias))

			// Если URL не найден - возвращаем 404
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		log.Info("redirecting",
			slog.String("from", alias),
			slog.String("to", red_url))

		http.Redirect(w, r, red_url, http.StatusMovedPermanently)
	}

}
