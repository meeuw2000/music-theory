// An interval is the difference between two pitches.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterval(t *testing.T) {
	assert.Equal(t, 15, len(IntervalOrder))
}
