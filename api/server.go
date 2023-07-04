package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type IRepository interface {
	SaveDevice(device domain.Device) error
	GetDevice(id uuid.UUID) (domain.Device, error)
	GetAllDevices() ([]domain.Device, error)
	SaveTransaction(domain.Transaction) error
	UpdateLastSignatureAndCounter(deviceID uuid.UUID, lastSignature string) error
	GetTransactions(id uuid.UUID) ([]domain.Transaction, error)
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
	repository    IRepository
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string, repository IRepository) *Server {
	return &Server{
		listenAddress: listenAddress,
		repository:    repository,
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	r := mux.NewRouter()

	// r.Handle("/api/v0/health", http.HandlerFunc(s.Health))
	// r.Handle("/api/v0/devices", http.HandlerFunc(s.Devices))
	// r.Handle("/api/v0/sign", http.HandlerFunc(s.Sign))
	// r.Handle("/api/v0/transactions/{deviceID}", http.HandlerFunc(s.Transactions))

	r.HandleFunc("/api/v0/health", s.Health).Methods("GET")
	r.HandleFunc("/api/v0/devices", s.GetAllDevices).Methods("GET")
	r.HandleFunc("/api/v0/devices", s.CreateDevice).Methods("POST")
	r.HandleFunc("/api/v0/devices/{deviceID}/sign", s.Sign).Methods("POST")
	r.HandleFunc("/api/v0/devices/{deviceID}/transactions", s.Transactions).Methods("GET")

	// TODO: register further HandlerFuncs here ...

	return http.ListenAndServe(s.listenAddress, r)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, errors []string) {
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: errors,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

// WriteAPIResponse takes an HTTP status code and a generic data struct
// and writes those as an HTTP response in a structured format.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)

	response := Response{
		Data: data,
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}
