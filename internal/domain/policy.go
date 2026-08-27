package domain

import "fmt"

type Permission string

const (
	PermissionDownload Permission = "download"
	PermissionDispatch Permission = "dispatch"
	PermissionExport   Permission = "export"
)

func CanAccess(role Role, sensitivity Sensitivity, p Permission) bool {
	if role == RoleLead {
		return true
	}
	if sensitivity == Sensitive {
		return role == RoleCurator && p == PermissionDispatch
	}
	if sensitivity == Restricted {
		return role == RoleCurator || role == RoleReviewer
	}
	return role != RoleAnnotator || p != PermissionExport
}
func ValidateAIUse(source, boundary string) error {
	if source == "" || boundary == "" {
		return fmt.Errorf("%w: AI provenance required", ErrInvalid)
	}
	return nil
}
func ValidateLanguage(code string) error {
	if len(code) < 2 || len(code) > 8 {
		return fmt.Errorf("%w: language code", ErrInvalid)
	}
	return nil
}
