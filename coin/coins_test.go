package coin

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/iov-one/weave/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mustCombineCoins has one return value for tests...
func mustCombineCoins(cs ...Coin) Coins {
	s, err := CombineCoins(cs...)
	if err != nil {
		panic(err)
	}
	return s
}

func TestMakeCoins(t *testing.T) {
	// TODO: verify constructor checks well for errors

	cases := []struct {
		inputs   []Coin
		isEmpty  bool
		isNonNeg bool
		has      []Coin // <= the wallet
		dontHave []Coin // > or outside the wallet
		isErr    bool
	}{
		// empty
		{
			nil,
			true,
			true,
			nil,
			[]Coin{NewCoin(0, 0, "")},
			false,
		},
		// ignore 0
		{
			[]Coin{NewCoin(0, 0, "FOO")},
			true,
			true,
			nil,
			[]Coin{NewCoin(0, 0, "FOO")},
			false,
		},
		// simple
		{
			[]Coin{NewCoin(40, 0, "FUD")},
			false,
			true,
			[]Coin{NewCoin(10, 0, "FUD"), NewCoin(40, 0, "FUD")},
			[]Coin{NewCoin(40, 1, "FUD"), NewCoin(40, 0, "FUN")},
			false,
		},
		// out of order, with negative
		{
			[]Coin{NewCoin(-20, -3, "FIN"), NewCoin(40, 5, "BON")},
			false,
			false,
			[]Coin{NewCoin(40, 4, "BON"), NewCoin(-30, 0, "FIN")},
			[]Coin{NewCoin(40, 6, "BON"), NewCoin(-20, 0, "FIN")},
			false,
		},
		// combine and remove
		{
			[]Coin{NewCoin(-123, -456, "BOO"), NewCoin(123, 456, "BOO")},
			true,
			true,
			nil,
			[]Coin{NewCoin(0, 0, "BOO")},
			false,
		},
		// safely combine
		{
			[]Coin{NewCoin(12, 0, "ADA"), NewCoin(-123, -456, "BOO"), NewCoin(124, 756, "BOO")},
			false,
			true,
			[]Coin{NewCoin(12, 0, "ADA"), NewCoin(1, 300, "BOO")},
			[]Coin{NewCoin(13, 0, "ADA"), NewCoin(1, 400, "BOO")},
			false,
		},
		// verify invalid input cur -> error
		{
			[]Coin{NewCoin(1, 2, "AL2")},
			false, false, nil, nil,
			true,
		},
		// verify invalid input values -> error
		{
			[]Coin{NewCoin(MaxInt+3, 2, "AND")},
			false, false, nil, nil,
			true,
		},
		// if we can combine invalid inputs, then acceptable?
		{
			[]Coin{NewCoin(MaxInt+3, 2, "AND"), NewCoin(-10, 0, "AND")},
			false,
			true,
			[]Coin{NewCoin(MaxInt-8, 0, "AND")},
			nil,
			false,
		},
	}

	for idx, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			s, err := CombineCoins(tc.inputs...)
			if tc.isErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.NoError(t, s.Validate())
			assert.Equal(t, tc.isEmpty, s.IsEmpty())
			assert.Equal(t, tc.isNonNeg, s.IsNonNegative())

			for _, h := range tc.has {
				assert.True(t, s.Contains(h))
			}
			for _, d := range tc.dontHave {
				assert.False(t, s.Contains(d))
			}
		})
	}
}

// TestCombine checks combine and equals
// and thereby checks add
func TestCombine(t *testing.T) {
	cases := []struct {
		a, b  Coins
		comb  Coins
		isErr bool
	}{
		// empty
		{
			mustCombineCoins(), mustCombineCoins(), mustCombineCoins(), false,
		},
		// one plus one
		{
			mustCombineCoins(NewCoin(MaxInt, 5, "ABC")),
			mustCombineCoins(NewCoin(-MaxInt, -4, "ABC")),
			mustCombineCoins(NewCoin(0, 1, "ABC")),
			false,
		},
		// multiple
		{
			mustCombineCoins(NewCoin(7, 8, "FOO"), NewCoin(8, 9, "BAR")),
			mustCombineCoins(NewCoin(5, 4, "APE"), NewCoin(2, 1, "FOO")),
			mustCombineCoins(NewCoin(5, 4, "APE"), NewCoin(8, 9, "BAR"), NewCoin(9, 9, "FOO")),
			false,
		},
		// overflows
		{
			mustCombineCoins(NewCoin(MaxInt, 0, "ADA")),
			mustCombineCoins(NewCoin(2, 0, "ADA")),
			Coins{},
			true,
		},
	}

	for idx, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {

			ac := tc.a.Count()
			bc := tc.b.Count()

			res, err := tc.a.Combine(tc.b)
			// don't modify original Coinss
			assert.Equal(t, ac, tc.a.Count())
			assert.Equal(t, bc, tc.b.Count())
			if tc.isErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NoError(t, res.Validate())
			assert.True(t, tc.comb.Equals(res))
			// result should only be the same as an input
			// if the other input was empty
			assert.Equal(t, tc.a.IsEmpty(),
				tc.b.Equals(res))
			assert.Equal(t, tc.b.IsEmpty(),
				tc.a.Equals(res))
		})
	}
}

func TestCoinsNormalize(t *testing.T) {
	coinp := func(w, f int64, t string) *Coin {
		c := NewCoin(w, f, t)
		return &c
	}

	cases := map[string]struct {
		coins     Coins
		wantCoins Coins
		wantErr   error
	}{
		"nil coins": {
			coins:     nil,
			wantCoins: nil,
		},
		"empty coins": {
			coins:     make(Coins, 0),
			wantCoins: nil,
		},
		"one zero coin": {
			coins:     Coins{coinp(0, 0, "BTC")},
			wantCoins: nil,
		},
		"one non zero coin": {
			coins:     Coins{coinp(1, 1, "BTC")},
			wantCoins: Coins{coinp(1, 1, "BTC")},
		},
		"coins sum to zero": {
			coins: Coins{
				coinp(1, 1, "BTC"),
				coinp(-1, -1, "BTC"),
			},
			wantCoins: nil,
		},
		"coins sum to non zero": {
			coins: Coins{
				coinp(1, 1, "BTC"),
				coinp(2, 2, "BTC"),
			},
			wantCoins: []*Coin{
				coinp(3, 3, "BTC"),
			},
		},
		"unordered coins": {
			coins: Coins{
				coinp(2, 0, "B"),
				coinp(3, 0, "C"),
				coinp(1, 0, "A"),
			},
			wantCoins: []*Coin{
				coinp(1, 0, "A"),
				coinp(2, 0, "B"),
				coinp(3, 0, "C"),
			},
		},
		"unordered and split value coins": {
			coins: Coins{
				coinp(1, 0, "B"),
				coinp(1, 0, "C"),
				coinp(1, 0, "B"),
				coinp(1, 0, "A"),
				coinp(1, 0, "C"),
				coinp(1, 0, "C"),
			},
			wantCoins: []*Coin{
				coinp(1, 0, "A"),
				coinp(2, 0, "B"),
				coinp(3, 0, "C"),
			},
		},
		"multiple coins sum to zero": {
			coins: Coins{
				coinp(1, 0, "DOGE"),

				coinp(1, 0, "BTC"),
				coinp(-1, 0, "BTC"),

				coinp(-1, 0, "ETH"),
				coinp(2, 0, "ETH"),
				coinp(-1, 0, "ETH"),

				coinp(-1, 0, "DOGE"),
			},
			wantCoins: nil,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			got, err := NormalizeCoins(tc.coins)
			if !errors.Is(tc.wantErr, err) {
				t.Fatalf("want %+v error, got %+v", tc.wantErr, err)
			}
			if tc.wantErr == nil {
				if !reflect.DeepEqual(got, tc.wantCoins) {
					t.Logf(" got: %s", Coins(got))
					t.Logf("want: %s", Coins(tc.wantCoins))
					t.Fatal("unexpected result")
				}
			}
		})
	}
}

func BenchmarkCoinsNormalize(b *testing.B) {
	coinp := func(w, f int64, t string) *Coin {
		c := NewCoin(w, f, t)
		return &c
	}

	benchmarks := map[string]Coins{
		"nil coins":      nil,
		"zero len coins": make(Coins, 0),
		"one coin":       Coins{coinp(1, 0, "ETH")},
		"two normalized coins": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
		},
		"two unordered coins": Coins{
			coinp(1, 0, "B"),
			coinp(1, 0, "C"),
		},
		"two split coins": Coins{
			coinp(1, 0, "BTC"),
			coinp(1, 0, "BTC"),
		},
		"four normalized": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
			coinp(1, 0, "C"),
			coinp(1, 0, "D"),
		},
		"four not normalized": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "C"),
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
		},
		"six not normalized": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "C"),
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
			coinp(-1, 0, "B"),
			coinp(1, 0, "D"),
		},
		"six normalized": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
			coinp(1, 0, "C"),
			coinp(1, 0, "D"),
			coinp(-1, 0, "E"),
			coinp(-1, 0, "F"),
		},
		"twelve normalized": Coins{
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
			coinp(1, 0, "C"),
			coinp(1, 0, "D"),
			coinp(-1, 0, "E"),
			coinp(-1, 0, "F"),
			coinp(-1, 0, "G"),
			coinp(-1, 0, "H"),
			coinp(-1, 0, "I"),
			coinp(-1, 0, "J"),
			coinp(-1, 0, "K"),
			coinp(-1, 0, "L"),
		},
		"twelve not normalized": Coins{
			coinp(-1, 0, "G"),
			coinp(-1, 0, "H"),
			coinp(-1, 0, "A"),
			coinp(-1, 0, "H"),
			coinp(1, 0, "A"),
			coinp(1, 0, "B"),
			coinp(1, 0, "C"),
			coinp(1, 0, "D"),
			coinp(-1, 0, "E"),
			coinp(-1, 0, "F"),
			coinp(-1, 0, "A"),
			coinp(-1, 0, "H"),
		},
	}

	for benchName, coins := range benchmarks {
		b.Run(benchName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NormalizeCoins(coins)
			}
		})
	}
}

func TestCoinsIsNormalized(t *testing.T) {
	coinp := func(w, f int64, t string) *Coin {
		c := NewCoin(w, f, t)
		return &c
	}

	cases := map[string]struct {
		coins Coins
		want  bool
	}{
		"nil": {
			coins: nil,
			want:  true,
		},
		"empty": {
			coins: []*Coin{},
			want:  true,
		},
		"one non zero coin": {
			coins: []*Coin{coinp(1, 0, "BTC")},
			want:  true,
		},
		"one zero coin": {
			coins: []*Coin{coinp(0, 0, "BTC")},
			want:  false,
		},
		"normalized": {
			coins: []*Coin{
				coinp(1, 0, "A"),
				coinp(0, 1, "B"),
				coinp(0, 1, "C"),
				coinp(1, 0, "D"),
			},
			want: true,
		},
		"unordered": {
			coins: []*Coin{
				coinp(1, 0, "A"),
				coinp(0, 1, "C"),
				coinp(0, 1, "B"),
				coinp(1, 0, "D"),
			},
			want: false,
		},
		"repeating currency": {
			coins: []*Coin{
				coinp(1, 0, "A"),
				coinp(0, 1, "A"),
				coinp(0, 1, "B"),
				coinp(1, 0, "C"),
			},
			want: false,
		},
		"one currency nil": {
			coins: []*Coin{
				coinp(1, 0, "A"),
				nil,
				coinp(0, 1, "B"),
				coinp(1, 0, "C"),
			},
			want: false,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if got := isNormalized(tc.coins); got != tc.want {
				t.Fatal("unexpected result")
			}
		})
	}
}
