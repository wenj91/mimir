// SPDX-License-Identifier: AGPL-3.0-only

package globalerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "an error (err-mimir-missing-metric-name)", Format(ErrIDMissingMetricName, "an error"))
}

func TestFormatWithLimitConfig(t *testing.T) {
	assert.Equal(t, "an error (err-mimir-missing-metric-name). You can adjust the related per-tenant limit setting -my-flag, or contacting your service administrator.", FormatWithLimitConfig(ErrIDMissingMetricName, "my-flag", "an error"))
}
