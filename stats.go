package nicehash

import (
	"encoding/json"
	"errors"
	"time"
)

type GlobalStats struct {
	Algo                  AlgoType `json:"algo"`
	ProfitabilityAboveBtc float32  `json:"profitability_above_btc,string"`
	ProfitabilityAboveLtc float32  `json:"profitability_above_ltc,string"`
	Price                 float64  `json:"price,string"`
	ProfitabilityBtc      float32  `json:"profitability_btc,string"`
	ProfitabilityLtc      float32  `json:"profitability_ltc,string"`
	Speed                 float64  `json:"speed,string"`
}

func (client *NicehashClient) GetStatsGlobalCurrent() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			Error string        `json:"error"`
			Stats []GlobalStats `json:"stats"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.global.current", Algo: AlgoTypeMAX, Location: LocationMAX}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, err
	}
	if stats.Result.Error != "" {
		return nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, nil
}

func (client *NicehashClient) GetStatsGlobalDay() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			Error string        `json:"error"`
			Stats []GlobalStats `json:"stats"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.global.24h", Algo: AlgoTypeMAX, Location: LocationMAX}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, err
	}
	if stats.Result.Error != "" {
		return nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, nil
}

type ProviderStats struct {
	Algo          AlgoType `json:"algo"`
	Balance       float64  `json:"balance,string"`
	AcceptedSpeed float64  `json:"accepted_speed,string"`
	RejectedSpeed float64  `json:"rejected_speed,string"`
}

type ProviderPayments struct {
	Amount float64   `json:"amount,string"`
	Fee    float64   `json:"fee,string"`
	TxID   string    `json:"TXID"`
	Time   time.Time `json:"time"`
}

func (t *ProviderPayments) UnmarshalJSON(data []byte) error {
	var err error
	type Alias ProviderPayments
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Time, err = time.Parse("2006-01-02 15:04:05", aux.Time)
	return err
}

func (client *NicehashClient) GetStatsProvider(addr string) ([]ProviderStats, []ProviderPayments, error) {
	stats := &struct {
		Result struct {
			Error    string             `json:"error"`
			Stats    []ProviderStats    `json:"stats"`
			Payments []ProviderPayments `json:"payments"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.provider", Algo: AlgoTypeMAX, Location: LocationMAX, Addr: addr}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, nil, err
	}
	if stats.Result.Error != "" {
		return nil, nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, stats.Result.Payments, nil
}
