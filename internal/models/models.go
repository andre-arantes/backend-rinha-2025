package models

import "time"

type Money struct {
	Value          int
	ConversionRate int
}

// Check if conversion rate is working properly in this context
func NewMoneyEntry(amount int, conversionRate int) *Money {
	return &Money{
		Value:          amount,
		ConversionRate: conversionRate,
	}
}

func (m *Money) Parse() string {
	return ""
}

type PaymentRequest struct {
	CorrelationID string    `json:"correlation_id"`
	Amount        int16     `json:"amount"`
	RequestAt     time.Time `json:"request_at"`
}
