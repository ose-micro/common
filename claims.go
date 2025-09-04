package common

import (
	"time"
)

// Tenant holds role + permissions for a specific tenant
type Tenant struct {
	Role        string       `json:"role"`
	Tenant      string       `json:"tenant"`
	Permissions []Permission `json:"permissions"`
}

type TokenKind string

const (
	AccessToken  TokenKind = "access"
	RefreshToken TokenKind = "refresh"
	PurposeToken TokenKind = "purpose"
)

// Claims defines our custom JWT payload
type Claims struct {
	Sub       string            `json:"sub"`
	Kind      TokenKind         `json:"typ"`
	Tenants   map[string]Tenant `json:"tenants,omitempty"`
	JTI       string            `json:"jti"`
	ExpiresAt *time.Duration    `json:"exp,omitempty"`
	IssuedAt  *time.Duration    `json:"iat,omitempty"`
	Issuer    string            `json:"iss,omitempty"`
}
