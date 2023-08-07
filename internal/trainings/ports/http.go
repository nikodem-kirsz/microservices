package ports

import (
	"net/http"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) GetTrainings(w http.ResponseWriter, r *http.Request) {
}

func (h HttpServer) CreateTraining(w http.ResponseWriter, r *http.Request) {
}

func (h HttpServer) CancelTraining(w http.ResponseWriter, r *http.Request, trainingUUID string) {
}

func (h HttpServer) RescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string) {
}

func (h HttpServer) RequestRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string) {
}

func (h HttpServer) ApproveRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string) {
}

func (h HttpServer) RejectRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string) {
}
