from money import Money


class Bank:
    def __init__(self):
        self.exchange_rates = {}

    def add_exchange_rate(self, currency_from, currency_to, rate):
        key = currency_from + "->" + currency_to
        self.exchange_rates[key] = rate

    def convert(self, a_money, a_currency):
        if a_money.currency == a_currency:
            return Money(a_money.amount, a_currency), None

        key = a_money.currency + "->" + a_currency
        if key in self.exchange_rates:
            return Money(a_money.amount * self.exchange_rates[key], a_currency), None

        return None, key
