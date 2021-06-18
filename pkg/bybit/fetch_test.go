package bybit

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchTickers(t *testing.T) {
	res, err := FetchTickers()
	assert.NoError(t, err)

	spew.Dump(res)
}
