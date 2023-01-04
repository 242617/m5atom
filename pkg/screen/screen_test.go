package screen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/242617/m5atom/pkg/screen"
)

func TestAddRemoveObjects(t *testing.T) {
	s, err := screen.New(screen.WithDimensions(1, 1))
	assert.NoError(t, err, "create screen")

	var (
		sample1 screen.Object
		sample2 screen.Object
	)

	s.Add(sample1)
	assert.Equal(t, 1, len(s.Objects()), "length 1")

	s.Add(sample2)
	assert.Equal(t, 2, len(s.Objects()), "length 2")

	s.Remove(sample1)
	assert.Equal(t, 1, len(s.Objects()), "length 1")

	s.Remove(sample2)
	assert.Equal(t, 0, len(s.Objects()), "zero length")
}
