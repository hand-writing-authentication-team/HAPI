package utils

import (
	"strings"

	"github.com/google/uuid"
)

func JobIDGenerator() string {
	randomUUID := uuid.New()
	jobID := strings.Replace(randomUUID.String(), "-", "", -1)
	return jobID
}
