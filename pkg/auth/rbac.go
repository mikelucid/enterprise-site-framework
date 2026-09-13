package auth

import "slices"

type RBAC struct {
	policies map[string][]string
}

func NewRBAC() *RBAC {
	return &RBAC{policies: map[string][]string{
		"admin":   {"*"},
		"manager": {"site.read", "site.write", "payment.read"},
		"user":    {"site.read"},
		"viewer":  {"site.read"},
	}}
}

func (r *RBAC) HasPermission(roles []string, permission string) bool {
	for _, role := range roles {
		perms := r.policies[role]
		if slices.Contains(perms, "*") || slices.Contains(perms, permission) {
			return true
		}
	}
	return false
}
