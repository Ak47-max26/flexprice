package webhookDto

import "github.com/flexprice/flexprice/internal/api/dto"

type InternalInvoiceEvent struct {
	InvoiceID string `json:"invoice_id"`
	TenantID  string `json:"tenant_id"`
}

type InvoiceWebhookPayload struct {
	EventType string               `json:"event_type"`
	Invoice   *dto.InvoiceResponse `json:"invoice"`
}

func NewInvoiceWebhookPayload(invoice *dto.InvoiceResponse, eventType string) *InvoiceWebhookPayload {
	return &InvoiceWebhookPayload{EventType: eventType, Invoice: invoice}
}
