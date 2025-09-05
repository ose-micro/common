package claims

import (
	"time"
)

// Claims defines our custom JWT dto
type Claims struct {
	Sub       string            `json:"sub"`
	Kind      TokenKind         `json:"typ"`
	Tenants   map[string]Tenant `json:"tenants,omitempty"`
	JTI       string            `json:"jti"`
	ExpiresAt *time.Duration    `json:"exp,omitempty"`
	IssuedAt  *time.Duration    `json:"iat,omitempty"`
	Issuer    string            `json:"iss,omitempty"`
}
