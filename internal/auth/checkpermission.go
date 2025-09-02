package auth

import (
	"context"
	"slices"
)

func CheckPermission(ctx context.Context, requiredRole string) bool {
	roles, ok := ctx.Value("roles").([]string)
	if !ok {
		return false
	}
	return slices.Contains(roles, requiredRole)
}
