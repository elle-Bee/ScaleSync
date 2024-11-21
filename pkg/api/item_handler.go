package api

import (
	"ScaleSync/pkg/metrics"
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type ItemHandler struct {
	Service service.ItemService
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("create_item"))
	defer timer.ObserveDuration() // Record the duration when the function exits

	metrics.ApiRequests.WithLabelValues("create_item").Inc()

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		metrics.ApiFailures.WithLabelValues("create_item").Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateItem(&item); err != nil {
		metrics.ApiFailures.WithLabelValues("create_item").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("create_item").Inc()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("get_items"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("get_items").Inc()

	items, err := h.Service.GetItems()
	if err != nil {
		metrics.ApiFailures.WithLabelValues("get_items").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("get_items").Inc()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
