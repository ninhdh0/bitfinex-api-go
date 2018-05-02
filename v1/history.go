package bitfinex

import "time"

type HistoryService struct {
	client *Client
}

type Balance struct {
	Currency    string
	Amount      string
	Balance     string
	Description string
	Timestamp   string
}

func (s *HistoryService) Balance(currency, wallet string, since, until time.Time, limit int) ([]Balance, error) {

	payload := map[string]interface{}{"currency": currency}

	if !since.IsZero() {
		payload["since"] = since.Unix()
	}
	if !until.IsZero() {
		payload["until"] = until.Unix()
	}
	if limit != 0 {
		payload["limit"] = limit
	}

	req, err := s.client.newAuthenticatedRequest("POST", "history", payload)

	if err != nil {
		return nil, err
	}

	var v []Balance

	_, err = s.client.do(req, &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

type Movement struct {
	ID               int64   `json:"id"`
	Txid             string  `json:"txid"`
	Currency         string  `json:"currency"`
	Method           string  `json:"method"`
	Type             string  `json:"type"`
	Amount           float64 `json:"amount,string"`
	Description      string  `json:"description"`
	Address          string  `json:"address"`
	Status           string  `json:"status"`
	Timestamp        float64 `json:"timestamp,string"`
	TimestampCreated float64 `json:"timestamp_created,string"`
	Fee              float64 `json:"fee,string"`
}

func (s *HistoryService) Movements(currency, method string, since, until time.Time, limit int) ([]Movement, error) {

	payload := map[string]interface{}{"currency": currency}

	if method != "" {
		payload["method"] = method
	}
	if !since.IsZero() {
		payload["since"] = since.Unix()
	}
	if !until.IsZero() {
		payload["until"] = until.Unix()
	}
	if limit != 0 {
		payload["limit"] = limit
	}

	req, err := s.client.newAuthenticatedRequest("POST", "history/movements", payload)

	if err != nil {
		return nil, err
	}

	var v []Movement

	_, err = s.client.do(req, &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

type PastTrade struct {
	Price       float64 `json:"price,string"`
	Amount      float64 `json:"amount,string"`
	Timestamp   float64 `json:"timestamp,string"`
	Exchange    string  `json:"exchange"`
	Type        string  `json:"type"`
	FeeCurrency string  `json:"fee_currency"`
	FeeAmount   float64 `json:"fee_amount,string"`
	TradeID     int64   `json:"tid"`
	OrderID     int64   `json:"order_id"`
}

func (s *HistoryService) Trades(pair string, since, until time.Time, limit int, reverse bool) ([]PastTrade, error) {
	payload := map[string]interface{}{"symbol": pair}

	if !since.IsZero() {
		payload["timestamp"] = since.Unix()
	}
	if !until.IsZero() {
		payload["until"] = until.Unix()
	}
	if limit != 0 {
		payload["limit_trades"] = limit
	}
	if reverse {
		payload["reverse"] = 1
	}

	req, err := s.client.newAuthenticatedRequest("POST", "mytrades", payload)

	if err != nil {
		return nil, err
	}

	var v []PastTrade

	_, err = s.client.do(req, &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}
