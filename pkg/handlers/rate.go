package handlers

import (
	"errors"
	"fmt"
	"github.com/rwlist/bubot/pkg/base"
	"github.com/rwlist/bubot/pkg/bybit"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Rate struct{}

func (r *Rate) HandleCommand(req *base.Request) {
	cmd := req.Command
	if cmd == nil {
		return
	}

	ticker := strings.TrimPrefix(cmd.Text, "rate")
	if cmd.Text == ticker {
		// no prefix
		return
	}
	ticker = strings.Trim(ticker, "_ ")

	info, err := r.fetchTicker(ticker)
	if err != nil {
		log.WithError(err).Error("failed to fetch ticker")
		_, _ = req.Response.Reply(fmt.Sprintf("failed to fetch: %s", err))
		return
	}

	strResponse := info.LastPrice.String()
	if strings.HasSuffix(ticker, "USDT") || strings.HasSuffix(ticker, "USD") {
		strResponse += "$"
	}

	_, _ = req.Response.Reply(strResponse)
}

func (r *Rate) fetchTicker(symbol string) (*bybit.Ticker, error) {
	tickers, err := bybit.FetchTickers()
	if err != nil {
		return nil, err
	}

	for _, ticker := range tickers.Result {
		if strings.EqualFold(symbol, ticker.Symbol) {
			return &ticker, nil
		}
	}

	return nil, errors.New("ticker not found")
}
