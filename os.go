package bandaid

import (
	"context"
	"os"
)

func ExpandEnvVars(ctx context.Context, value string) (string, error) {
	if value == "" {
		return "", nil
	}

	expandedCommand := os.ExpandEnv(value)
	return expandedCommand, nil
}
