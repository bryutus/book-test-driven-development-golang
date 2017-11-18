package chapter15

type sum struct {
	augend Expression
	addend Expression
}

func newSum(augend Money, addend Expression) Expression {
	return sum{augend: augend, addend: addend}
}

func (s sum) reduce(bank Bank, to Currency) Money {
	amount := s.augend.reduce(bank, to).getAmount() + s.addend.reduce(bank, to).getAmount()
	return newMoney(to, amount)
}

func (s sum) plus(addend Expression) Expression {
	return nil
}

func (s sum) getAugend() Expression {
	return s.augend
}

func (s sum) getAddend() Expression {
	return s.addend
}
