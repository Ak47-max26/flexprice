package webhookDto

import "github.com/flexprice/flexprice/internal/api/dto"

type InternalEntitlementEvent struct {
	EntitlementID string `json:"entitlement_id"`
	TenantID      string `json:"tenant_id"`
}

type EntitlementWebhookPayload struct {
	EventType   string                   `json:"event_type"`
	Entitlement *dto.EntitlementResponse `json:"entitlement"`
}

func NewEntitlementWebhookPayload(entitlement *dto.EntitlementResponse, eventType string) *EntitlementWebhookPayload {
	return &EntitlementWebhookPayload{EventType: eventType, Entitlement: entitlement}
}
