package util

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func BuildUuid() string {
	uuidStr := uuid.NewV4().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}
