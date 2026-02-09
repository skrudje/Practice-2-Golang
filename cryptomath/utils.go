package cryptomath

import (
	"math/big"
)

// хранит результат расширенного алгоритма евклида
type GCDResult struct {
	D *big.Int // НОД
	X *big.Int // кэф X
	Y *big.Int // кэф Y
}

// возведение в степень по модулю
func PowerMod(base, exp, mod *big.Int) *big.Int {
	res := big.NewInt(1)
	b := new(big.Int).Set(base)
	e := new(big.Int).Set(exp)
	m := new(big.Int).Set(mod)

	b.Mod(b, m)

	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// пока степень > 0
	for e.Cmp(zero) > 0 {
		// если степень нечетная (exp % 2 == 1)
		if new(big.Int).Mod(e, two).Cmp(one) == 0 {
			res.Mul(res, b)
			res.Mod(res, m)
		}
		// base = base * base % m
		b.Mul(b, b)
		b.Mod(b, m)
		// exp = exp / 2
		e.Div(e, two)
	}
	return res
}

// Возвращает d, x, y такие, что ax + by = d
func ExtendedGCD(a, b *big.Int) GCDResult {
	zero := big.NewInt(0)
	
	// базовый случай рекурсии: если b == 0
	if b.Cmp(zero) == 0 {
		return GCDResult{
			D: new(big.Int).Set(a),
			X: big.NewInt(1),
			Y: big.NewInt(0),
		}
	}

	// рекурсивный вызов ExtendedGCD(b, a % b)
	rem := new(big.Int).Mod(a, b) // a % b
	res := ExtendedGCD(b, rem)

	// x = y1
	// y = x1 - (a/b) * y1
	x := new(big.Int).Set(res.Y)
	
	div := new(big.Int).Div(a, b) // a / b
	mul := new(big.Int).Mul(div, res.Y) // (a/b) * y1
	y := new(big.Int).Sub(res.X, mul) // x1 - ...

	return GCDResult{D: res.D, X: x, Y: y}
}

// првоерка на простоту.
func IsPrime(n *big.Int) bool {
	return n.ProbablyPrime(20) // дает высокую гарантию
}

// ModInverse - поиск обратного элемента (таск 3)
func ModInverse(a, m *big.Int) *big.Int {
	res := ExtendedGCD(a, m)
	one := big.NewInt(1)
	// если нод != 1, обратного элемента нет
	if res.D.Cmp(one) != 0 {
		return nil
	}

	// приводим х к положительному значению по модулю m
	// (res.X % m + m) % m
	x := res.X
	x.Mod(x, m)
	if x.Sign() < 0 {
		x.Add(x, m)
	}
	return x
}