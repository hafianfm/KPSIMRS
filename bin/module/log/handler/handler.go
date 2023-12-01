package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/go-chi/chi/v5"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/usecase"
)

type LogResponse struct {
	Status string          `json:"status"`
	Data   []model.LogData `json:"data"`
}

type httpHandler struct {
	logUsecase usecase.LogUsecase
}

func InitLogHttpHandler(r *chi.Mux, uc usecase.LogUsecase) {
	handler := &httpHandler{
		logUsecase: uc,
	}

	r.Get("/api/v1/logs", handler.GetLogsHandler)
    r.Post("/api/v1/logs", handler.PushToBrokerLog)

}

func (h *httpHandler) GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// mengambil status
	status := r.URL.Query().Get("Filter")

	// mengambil search
	search := r.URL.Query().Get("Search")

	logs, err := h.logUsecase.GetLogs(status, search)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(LogResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), http.StatusText(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LogResponse{
		Status: fmt.Sprintf("Success (%s)", http.StatusText(http.StatusOK)),
		Data:   logs,
	})
}

func (h *httpHandler) PushToBrokerLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer producer.Close()
	var data []model.LogData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := json.Marshal(data)

	msg := &sarama.ProducerMessage{
		Topic: "log6",
		Value: sarama.StringEncoder(bytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
