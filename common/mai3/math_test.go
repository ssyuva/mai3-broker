package mai3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGss(t *testing.T) {
	f := func(x float64) float64 {
		tmp := x - 2
		return tmp * tmp
	}
	aa, bb := Gss(f, 1, 5, 1e-8, 15)
	assert.Equal(t, 1.9987067953999975, aa)
	assert.Equal(t, 2.001639345143427, bb)
}
