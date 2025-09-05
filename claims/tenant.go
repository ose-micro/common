package claims

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
