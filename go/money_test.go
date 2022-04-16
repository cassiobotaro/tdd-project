package main

import "testing"

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * float64(multiplier),
		currency: m.currency,
	}
}

func (m Money) Divide(divisor int) Money {
	return Money{
		amount:   m.amount / float64(divisor),
		currency: m.currency,
	}
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got: %+v", expected, actual)
	}
}

type Portifolio []Money

func (p Portifolio) Add(money Money) Portifolio {
	p = append(p, money)
	return p
}

func (p Portifolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}
	return Money{amount: total, currency: currency}
}

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	actualResult := fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	assertEqual(t, expectedResult, actualResult)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualResult := originalMoney.Divide(4)
	expectedResult := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

func TestAddition(t *testing.T) {
	var portifolio Portifolio
	var portifolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portifolio = portifolio.Add(fiveDollars)
	portifolio = portifolio.Add(tenDollars)
	portifolioInDollars = portifolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portifolioInDollars)
}
