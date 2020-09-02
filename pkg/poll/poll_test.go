package poll

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamplePoll(t *testing.T) {
	assert.Equal(t, `DAF Pub

tue 20:00: @dendi239
tue 21:00: @ostashev, @dendi239, @belkka
wen 20:00: @ostashev, @dendi239
wen 19:00: @ostashev, @dendi239
wen 21:00: @ostashev, @belkka
stay home: @belkka, @viskonsin
`, fmt.Sprint(&Sample))
}
