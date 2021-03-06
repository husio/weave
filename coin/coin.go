package coin

import (
	"regexp"

	"github.com/iov-one/weave/errors"
)

//-------------- Coin -----------------------

// IsCC is the RegExp to ensure valid currency codes
var IsCC = regexp.MustCompile(`^[A-Z]{3,4}$`).MatchString

const (
	// MaxInt is the largest whole value we accept
	MaxInt int64 = 999999999999999 // 10^15-1
	// MinInt is the lowest whole value we accept
	MinInt = -MaxInt

	// FracUnit is the smallest numbers we divide by
	FracUnit int64 = 1000000000 // fractional units = 10^9
	// MaxFrac is the highest possible fractional value
	MaxFrac = FracUnit - 1
	// MinFrac is the lowest possible fractional value
	MinFrac = -MaxFrac
)

// NewCoin creates a new coin object
func NewCoin(whole int64, fractional int64, ticker string) Coin {
	return Coin{
		Whole:      whole,
		Fractional: fractional,
		Ticker:     ticker,
	}
}

// NewCoinp returns a pointer to a new coin.
func NewCoinp(whole, fractional int64, ticker string) *Coin {
	c := NewCoin(whole, fractional, ticker)
	return &c
}

// ID returns a coin ticker name.
func (c Coin) ID() string {
	return c.Ticker
}

// Split divides the value of a coin into given amount of pieces and returns a
// single piece.
// It might be that a precise splitting is not possible. Any leftover of a
// fractional value is returned as well.
// For example splitting 4 EUR into 3 pieces will result in a single piece
// being 1.33 EUR and 1 cent returned as the rest (leftover).
//   4 = 1.33 x 3 + 1
func (c Coin) Divide(pieces int64) (Coin, Coin, error) {
	// This is an invalid use of the method.
	if pieces <= 0 {
		zero := Coin{Ticker: c.Ticker}
		return zero, zero, errors.ErrHuman.New("pieces must be greater than zero")
	}

	// When dividing whole and there is a leftover then convert it to
	// fractional and split as well.
	fractional := c.Fractional
	if leftover := c.Whole % pieces; leftover != 0 {
		fractional += leftover * FracUnit
	}

	one := Coin{
		Ticker:     c.Ticker,
		Whole:      c.Whole / pieces,
		Fractional: fractional / pieces,
	}
	rest := Coin{
		Ticker:     c.Ticker,
		Whole:      0, // This we can always divide.
		Fractional: fractional % pieces,
	}
	return one, rest, nil
}

// Multiply returns the result of a coin value multiplication. This method can
// fail if the result would overflow maximum coin value.
func (c Coin) Multiply(times int64) (Coin, error) {
	if times == 0 || (c.Whole == 0 && c.Fractional == 0) {
		return Coin{Ticker: c.Ticker}, nil
	}

	whole, err := mul64(c.Whole, times)
	if err != nil {
		return Coin{}, err

	}
	frac, err := mul64(c.Fractional, times)
	if err != nil {
		return Coin{}, err
	}

	// Normalize if fractional value overflows.
	if frac > FracUnit {
		if n := whole + frac/FracUnit; n < whole {
			return Coin{}, errors.ErrOverflow
		} else {
			whole = n
		}
		frac = frac % FracUnit
	}

	res := Coin{
		Ticker:     c.Ticker,
		Whole:      whole,
		Fractional: frac,
	}
	return res, nil
}

// mul64 multiplies two int64 numbers. If the result overflows the int64 size
// the ErrOverflow is returned.
func mul64(a, b int64) (int64, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	c := a * b
	if c/a != b {
		return c, errors.ErrOverflow
	}
	return c, nil
}

// Add combines two coins.
// Returns error if they are of different
// currencies, or if the combination would cause
// an overflow
func (c Coin) Add(o Coin) (Coin, error) {
	// If any of the coins represents no value and does not have a ticker
	// set then it has no influence on the addition result.
	if c.Ticker == "" && c.IsZero() {
		return o, nil
	}
	if o.Ticker == "" && o.IsZero() {
		return c, nil
	}

	if !c.SameType(o) {
		err := ErrInvalidCurrency.Newf("adding %s to %s", c.Ticker, o.Ticker)
		return Coin{}, err
	}

	c.Whole += o.Whole
	c.Fractional += o.Fractional
	return c.normalize()
}

// Negative returns the opposite coins value
//   c.Add(c.Negative()).IsZero() == true
func (c Coin) Negative() Coin {
	return Coin{
		Ticker:     c.Ticker,
		Whole:      -1 * c.Whole,
		Fractional: -1 * c.Fractional,
	}
}

// Subtract given amount.
func (c Coin) Subtract(amount Coin) (Coin, error) {
	return c.Add(amount.Negative())
}

// Compare will check values of two coins, without
// inspecting the currency code. It is up to the caller
// to determine if they want to check this.
// It also assumes they were already normalized.
//
// Returns 1 if c is larger, -1 if o is larger, 0 if equal
func (c Coin) Compare(o Coin) int {
	if c.Whole > o.Whole {
		return 1
	}
	if c.Whole < o.Whole {
		return -1
	}
	// same integer, compare fractional
	if c.Fractional > o.Fractional {
		return 1
	}
	if c.Fractional < o.Fractional {
		return -1
	}
	// actually the same...
	return 0
}

// Equals returns true if all fields are identical
func (c Coin) Equals(o Coin) bool {
	return c.Ticker == o.Ticker &&
		c.Whole == o.Whole &&
		c.Fractional == o.Fractional
}

// IsEmpty returns true on null or zero amount
func IsEmpty(c *Coin) bool {
	return c == nil || c.IsZero()
}

// IsZero returns true amounts are 0
func (c Coin) IsZero() bool {
	return c.Whole == 0 && c.Fractional == 0
}

// IsPositive returns true if the value is greater than 0
func (c Coin) IsPositive() bool {
	return c.Whole > 0 ||
		(c.Whole == 0 && c.Fractional > 0)
}

// IsNonNegative returns true if the value is 0 or higher
func (c Coin) IsNonNegative() bool {
	return c.Whole >= 0 && c.Fractional >= 0
}

// IsGTE returns true if c is same type and at least
// as large as o.
// It assumes they were already normalized.
func (c Coin) IsGTE(o Coin) bool {
	if !c.SameType(o) || c.Whole < o.Whole {
		return false
	}
	if (c.Whole == o.Whole) &&
		(c.Fractional < o.Fractional) {
		return false
	}
	return true
}

// SameType returns true if they have the same currency
func (c Coin) SameType(o Coin) bool {
	return c.Ticker == o.Ticker
}

// Clone provides an independent copy of a coin pointer
func (c *Coin) Clone() *Coin {
	if c == nil {
		return nil
	}
	return &Coin{
		Ticker:     c.Ticker,
		Whole:      c.Whole,
		Fractional: c.Fractional,
	}
}

// Validate ensures that the coin is in the valid range
// and valid currency code. It accepts negative values,
// so you may want to make other checks in your business
// logic
func (c Coin) Validate() error {
	if !IsCC(c.Ticker) {
		return ErrInvalidCurrency.New(c.Ticker)
	}
	if c.Whole < MinInt || c.Whole > MaxInt {
		return ErrInvalidCoin.New(outOfRange)
	}
	if c.Fractional < MinFrac || c.Fractional > MaxFrac {
		return ErrInvalidCoin.New(outOfRange)
	}
	// make sure signs match
	if c.Whole != 0 && c.Fractional != 0 &&
		((c.Whole > 0) != (c.Fractional > 0)) {
		return ErrInvalidCoin.New("mismatched sign")
	}

	return nil
}

// normalize will adjust the fractional parts to
// correspond to the range and the integer parts.
//
// If the normalized coin is outside of the range,
// returns an error
func (c Coin) normalize() (Coin, error) {
	// keep fraction in range
	for c.Fractional < MinFrac {
		c.Whole--
		c.Fractional += FracUnit
	}
	for c.Fractional > MaxFrac {
		c.Whole++
		c.Fractional -= FracUnit
	}

	// make sure the signs correspond
	if (c.Whole > 0) && (c.Fractional < 0) {
		c.Whole--
		c.Fractional += FracUnit
	} else if (c.Whole < 0) && (c.Fractional > 0) {
		c.Whole++
		c.Fractional -= FracUnit
	}

	// return error if integer is out of range
	if c.Whole < MinInt || c.Whole > MaxInt {
		return Coin{}, ErrInvalidCoin.New(outOfRange)
	}
	return c, nil
}
