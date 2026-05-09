package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type TimeResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request received",
			"ip", r.RemoteAddr,
			"method", r.Method,
			"path", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

// chapter 14
func NewTimeoutMiddleware(duration time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx, cancelFunc := context.WithTimeout(req.Context(), duration)
			defer cancelFunc()
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()

	// Content Negotiation: Revisamos qué formato quiere el cliente
	if r.Header.Get("Accept") == "application/json" {
		resp := TimeResponse{
			DayOfWeek:  now.Weekday().String(),
			DayOfMonth: now.Day(),
			Month:      now.Month().String(),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Default: Plain Text RFC3339
	currentTime := now.Format(time.RFC3339)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(currentTime))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.Handle("/time", LogginMiddleware(http.HandlerFunc(timeHandler)))

	slog.Info("server running on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
