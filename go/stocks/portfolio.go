package stocks

import "errors"

type Portifolio []Money

func (p Portifolio) Add(money Money) Portifolio {
	p = append(p, money)
	return p
}

func (p Portifolio) Evaluate(currency string) (Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedAmount, ok := convert(m, currency); ok {
			total += convertedAmount
		} else {
			failedConversions = append(failedConversions, m.currency+"->"+currency)
		}
	}
	if len(failedConversions) == 0 {
		return NewMoney(total, currency), nil
	}
	failures := "["
	for _, f := range failedConversions {
		failures += f + ","
	}
	failures += "]"
	return NewMoney(0, ""), errors.New("Missing exchange rate(s):" + failures)
}

func convert(money Money, currency string) (float64, bool) {
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount, true
	}
	key := money.currency + "->" + currency
	rate, ok := exchangeRates[key]
	return money.amount * rate, ok
}
