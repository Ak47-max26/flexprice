package webhookDto

import "github.com/flexprice/flexprice/internal/api/dto"

type InternalCustomerEvent struct {
	CustomerID string `json:"customer_id"`
	TenantID   string `json:"tenant_id"`
}

type CustomerWebhookPayload struct {
	EventType string                `json:"event_type"`
	Customer  *dto.CustomerResponse `json:"customer"`
}

func NewCustomerWebhookPayload(customer *dto.CustomerResponse, eventType string) *CustomerWebhookPayload {
	return &CustomerWebhookPayload{EventType: eventType, Customer: customer}
}
