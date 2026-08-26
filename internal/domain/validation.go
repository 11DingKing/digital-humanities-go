package domain

import (
	"fmt"
	"strings"
	"time"
)

func ValidateProjectName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("%w: project name", ErrInvalid)
	}
	return nil
}

func ValidateCorpusTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("%w: corpus title", ErrInvalid)
	}
	return nil
}

func ValidateLicense(license string) error {
	if strings.TrimSpace(license) == "" {
		return fmt.Errorf("%w: license", ErrInvalid)
	}
	return nil
}

func ValidateSegment(segment string) error {
	if len(strings.TrimSpace(segment)) < 1 {
		return fmt.Errorf("%w: empty segment", ErrInvalid)
	}
	return nil
}

func ValidateQuota(quota int64) error {
	if quota < 1 {
		return fmt.Errorf("%w: quota", ErrInvalid)
	}
	return nil
}

func ValidatePriority(priority int) error {
	if priority < 1 || priority > 10 {
		return fmt.Errorf("%w: priority", ErrInvalid)
	}
	return nil
}

func ValidateConcurrency(n int) error {
	if n < 1 || n > 32 {
		return fmt.Errorf("%w: concurrency", ErrInvalid)
	}
	return nil
}

func ValidateExpiry(expiry time.Time) error {
	if !expiry.After(time.Now()) {
		return fmt.Errorf("%w: expiry", ErrInvalid)
	}
	return nil
}

func ValidateRequestID(id string) error {
	if len(id) < 4 {
		return fmt.Errorf("%w: request id", ErrInvalid)
	}
	return nil
}

func ValidateRole(role Role) error {
	switch role {
	case RoleLead, RoleCurator, RoleAnnotator, RoleReviewer:
		return nil
	}
	return fmt.Errorf("%w: role", ErrInvalid)
}

func ValidateSensitivity(s Sensitivity) error {
	switch s {
	case Public, Restricted, Sensitive:
		return nil
	}
	return fmt.Errorf("%w: sensitivity", ErrInvalid)
}

func ValidateStatus(status ProjectStatus) error {
	switch status {
	case Draft, Active, Archived:
		return nil
	}
	return fmt.Errorf("%w: status", ErrInvalid)
}

func ValidateMetadataField1(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField2(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField3(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField4(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField5(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField6(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField7(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField8(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField9(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField10(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField11(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField12(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField13(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField14(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField15(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField16(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField17(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField18(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField19(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField20(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField21(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField22(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField23(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField24(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField25(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField26(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField27(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField28(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField29(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField30(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField31(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField32(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField33(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField34(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField35(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField36(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField37(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField38(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField39(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField40(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField41(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField42(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField43(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField44(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}

func ValidateMetadataField45(value string) error {
	if len(value) > 4096 {
		return fmt.Errorf("%w: metadata too long", ErrInvalid)
	}
	if strings.IndexByte(value, 0) >= 0 {
		return fmt.Errorf("%w: metadata contains NUL", ErrInvalid)
	}
	return nil
}
