package webhookDto

import "github.com/flexprice/flexprice/internal/api/dto"

type InternalCommunicationEvent struct {
	InvoiceID string `json:"invoice_id"`
	TenantID  string `json:"tenant_id"`
}

type CommunicationWebhookPayload struct {
	EventType string               `json:"event_type"`
	Invoice   *dto.InvoiceResponse `json:"invoice"`
}

func NewCommunicationWebhookPayload(invoice *dto.InvoiceResponse, eventType string) *CommunicationWebhookPayload {
	return &CommunicationWebhookPayload{EventType: eventType, Invoice: invoice}
}
