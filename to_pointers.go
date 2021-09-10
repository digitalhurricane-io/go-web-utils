package utils

import (
	"time"
)

// StringP copy string and return pointer
func StringP(s string) *string {
	ss := s
	return &ss
}

// BoolP copy bool and return pointer
func BoolP(b bool) *bool {
	bb := b
	return &bb
}

// Int64P copy int64 and return pointer
func Int64P(i int64) *int64 {
	ii := i
	return &ii
}

// Float64P copy float64 and return pointer
func Float64P(i float64) *float64 {
	ii := i
	return &ii
}

// IntP copy int and return pointer
func IntP(i int) *int {
	ii := i
	return &ii
}

// TimeP copy time and return pointer
func TimeP(t time.Time) *time.Time {
	tt := t
	return &tt
}

func SafeTimeDeref(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Now().UTC()
}

func SafeStringDeref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func SafeBoolDeref(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func SafeInt64Deref(i *int64) int64 {
	if i != nil {
		return *i
	}
	return 0
}

func SafeIntDeref(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

func SafeFloat64Deref(i *float64) float64 {
	if i != nil {
		return *i
	}
	return 0
}
