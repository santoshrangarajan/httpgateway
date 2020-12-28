package main

////https://github.com/dlsniper/gopherconuk
import (
	"log"
	"net/http"
)

// NumbersService is ...
type NumbersService struct {
	logger    *log.Logger
	dbService *DbService
}

// NewNumbersService is ...
func NewNumbersService(logger *log.Logger, dbService *DbService) *NumbersService {
	return &NumbersService{
		logger:    logger,
		dbService: dbService,
	}
}

// Home is ...
func (h *NumbersService) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	h.logger.Println("Home called")
	h.dbService.open()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}

// SetupRoutes is ...
func (h *NumbersService) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/home", h.Home)
}
