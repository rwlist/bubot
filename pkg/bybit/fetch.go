package bybit

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
)

const (
	tickersURL = "https://api.bybit.com/v2/public/tickers"
)

type Tickers struct {
	RetCode int      `json:"ret_code"`
	RetMsg  string   `json:"ret_msg"`
	ExtCode string   `json:"ext_code"`
	ExtInfo string   `json:"ext_info"`
	Result  []Ticker `json:"result"`
	TimeNow string   `json:"time_now"`
}

type Ticker struct {
	Symbol                 string          `json:"symbol"`
	BidPrice               decimal.Decimal `json:"bid_price"`
	AskPrice               decimal.Decimal `json:"ask_price"`
	LastPrice              decimal.Decimal `json:"last_price"`
	LastTickDirection      string          `json:"last_tick_direction"`
	PrevPrice24H           decimal.Decimal `json:"prev_price_24h"`
	Price24HPcnt           decimal.Decimal `json:"price_24h_pcnt"`
	HighPrice24H           decimal.Decimal `json:"high_price_24h"`
	LowPrice24H            decimal.Decimal `json:"low_price_24h"`
	PrevPrice1H            decimal.Decimal `json:"prev_price_1h"`
	Price1HPcnt            decimal.Decimal `json:"price_1h_pcnt"`
	MarkPrice              decimal.Decimal `json:"mark_price"`
	IndexPrice             decimal.Decimal `json:"index_price"`
	OpenInterest           decimal.Decimal `json:"open_interest"`
	OpenValue              decimal.Decimal `json:"open_value"`
	TotalTurnover          decimal.Decimal `json:"total_turnover"`
	Turnover24H            decimal.Decimal `json:"turnover_24h"`
	TotalVolume            decimal.Decimal `json:"total_volume"`
	Volume24H              decimal.Decimal `json:"volume_24h"`
	FundingRate            decimal.Decimal `json:"funding_rate"`
	PredictedFundingRate   decimal.Decimal `json:"predicted_funding_rate"`
	NextFundingTime        string          `json:"next_funding_time"`
	CountdownHour          int             `json:"countdown_hour"`
	DeliveryFeeRate        string          `json:"delivery_fee_rate"`
	PredictedDeliveryPrice string          `json:"predicted_delivery_price"`
	DeliveryTime           string          `json:"delivery_time"`
}

func FetchTickers() (*Tickers, error) {
	resp, err := http.Get(tickersURL)
	if err != nil {
		return nil, err
	}

	var res Tickers
	err = json.NewDecoder(resp.Body).Decode(&res)
	return &res, err
}
