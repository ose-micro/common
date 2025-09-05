package claims

import (
	commonv1 "github.com/ose-micro/common/gen/go/ose/micro/common/v1"
)

func BuildClaimGrpc(claim Claims) *commonv1.Claims {
	return &commonv1.Claims{
		Sub:       claim.Sub,
		Kind:      buildKindGrpc(claim.Kind),
		Tenants:   buildTenantGrpc(claim.Tenants),
		Jti:       claim.JTI,
		Issuer:    claim.Issuer,
		ExpiresAt: claim.ExpiresAt,
		IssuedAt:  claim.IssuedAt,
	}
}

func BuildClaimDto(claims commonv1.Claims) *Claims {
	kind := buildKindDto(claims.Kind)
	tenants := make(map[string]Tenant)

	for key, tenant := range claims.Tenants {
		tenants[key] = buildTenantDto(*tenant)
	}

	return &Claims{
		Sub:       claims.Sub,
		Kind:      kind,
		Tenants:   tenants,
		Issuer:    claims.Issuer,
		ExpiresAt: claims.ExpiresAt,
		IssuedAt:  claims.IssuedAt,
		JTI:       claims.Jti,
	}
}

func buildKindDto(kind commonv1.TokenKind) TokenKind {
	switch kind {
	case commonv1.TokenKind_TokenKind_AccessToken:

		return AccessToken
	case commonv1.TokenKind_TokenKind_RefreshToken:
		return RefreshToken
	case commonv1.TokenKind_TokenKind_PurposeToken:
		return PurposeToken
	default:
		return ""
	}
}

func buildTenantDto(tenant commonv1.Tenant) Tenant {
	permissions := make([]Permission, len(tenant.Permissions))
	for _, permission := range tenant.Permissions {
		permissions = append(permissions, Permission{
			Resource: permission.Resource,
			Action:   permission.Action,
		})
	}

	return Tenant{
		Role:        tenant.Role,
		Tenant:      tenant.Tenant,
		Permissions: permissions,
	}
}

func buildKindGrpc(kind TokenKind) commonv1.TokenKind {
	switch kind {
	case AccessToken:
		return commonv1.TokenKind_TokenKind_AccessToken
	case RefreshToken:
		return commonv1.TokenKind_TokenKind_RefreshToken
	case PurposeToken:
		return commonv1.TokenKind_TokenKind_PurposeToken
	default:
		return commonv1.TokenKind_TokenKind_Undefined
	}
}

func buildTenantGrpc(payload map[string]Tenant) map[string]*commonv1.Tenant {
	tenants := make(map[string]*commonv1.Tenant)
	for key, tenant := range payload {
		permissions := make([]*commonv1.Permission, 0)
		for _, permission := range tenant.Permissions {
			permissions = append(permissions, &commonv1.Permission{
				Action:   permission.Action,
				Resource: permission.Resource,
			})
		}

		tenants[key] = &commonv1.Tenant{
			Role:        tenant.Role,
			Tenant:      tenant.Tenant,
			Permissions: permissions,
		}
	}

	return tenants
}
