package webhookDto

import "github.com/flexprice/flexprice/internal/api/dto"

type InternalPaymentEvent struct {
	PaymentID string `json:"payment_id"`
	TenantID  string `json:"tenant_id"`
}

type PaymentWebhookPayload struct {
	EventType string               `json:"event_type"`
	Payment   *dto.PaymentResponse `json:"payment"`
}

func NewPaymentWebhookPayload(payment *dto.PaymentResponse, eventType string) *PaymentWebhookPayload {
	return &PaymentWebhookPayload{EventType: eventType, Payment: payment}
}
