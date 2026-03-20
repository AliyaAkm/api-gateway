package handlers

import (
	"time"

	domain "gateway/internal/domain/gateway"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	serviceName string
	services    []domain.ServiceInfo
}

func New(serviceName string, services []domain.ServiceInfo) *HealthHandler {
	return &HealthHandler{
		serviceName: serviceName,
		services:    services,
	}
}

func (h *HealthHandler) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":      "ok",
		"service":     h.serviceName,
		"timestamp":   time.Now().UTC(),
		"description": "API gateway for the interactive learning platform",
		"services":    h.services,
	})
}
