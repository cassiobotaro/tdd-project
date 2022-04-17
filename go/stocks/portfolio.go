package stocks

import "errors"

type Portifolio []Money

func (p Portifolio) Add(money Money) Portifolio {
	p = append(p, money)
	return p
}

func (p Portifolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedCurrency, err := bank.Convert(m, currency); err == nil {
			total += convertedCurrency.amount
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}
	if len(failedConversions) == 0 {
		totalMoney := NewMoney(total, currency)
		return &totalMoney, nil
	}
	failures := "["
	for _, f := range failedConversions {
		failures += f + ","
	}
	failures += "]"
	return nil, errors.New("Missing exchange rate(s):" + failures)
}
