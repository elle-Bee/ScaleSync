package api

import (
	"ScaleSync/pkg/metrics"
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type WarehouseHandler struct {
	Service service.WarehouseService
}

// CreateWarehouse handler with metrics instrumentation
func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("create_warehouse"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("create_warehouse").Inc()

	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		metrics.ApiFailures.WithLabelValues("create_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateWarehouse(&warehouse); err != nil {
		metrics.ApiFailures.WithLabelValues("create_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("create_warehouse").Inc()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(warehouse)
}

// GetWarehouses handler with metrics instrumentation
func (h *WarehouseHandler) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("get_warehouses"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("get_warehouses").Inc()

	warehouses, err := h.Service.GetWarehouses()
	if err != nil {
		metrics.ApiFailures.WithLabelValues("get_warehouses").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("get_warehouses").Inc()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(warehouses)
}

// GetWarehouse handler with metrics instrumentation
func (h *WarehouseHandler) GetWarehouse(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("get_warehouse"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("get_warehouse").Inc()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		metrics.ApiFailures.WithLabelValues("get_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	warehouse, err := h.Service.GetWarehouse(id)
	if err != nil {
		metrics.ApiFailures.WithLabelValues("get_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("get_warehouse").Inc()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(warehouse)
}

// UpdateWarehouse handler with metrics instrumentation
func (h *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("update_warehouse"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("update_warehouse").Inc()

	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		metrics.ApiFailures.WithLabelValues("update_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateWarehouse(&warehouse); err != nil {
		metrics.ApiFailures.WithLabelValues("update_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("update_warehouse").Inc()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(warehouse)
}

// DeleteWarehouse handler with metrics instrumentation
func (h *WarehouseHandler) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("delete_warehouse"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("delete_warehouse").Inc()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		metrics.ApiFailures.WithLabelValues("delete_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteWarehouse(id); err != nil {
		metrics.ApiFailures.WithLabelValues("delete_warehouse").Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics.ApiSuccesses.WithLabelValues("delete_warehouse").Inc()
	w.WriteHeader(http.StatusNoContent)
}
