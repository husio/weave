package currency

import (
	"encoding/json"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/store"
)

func TestGenesisKey(t *testing.T) {
	const genesis = `
		{
			"currencies": [
				{"ticker": "MCR", "name": "my currency", "sig_figs": 9},
				{"ticker": "DOGE", "name": "Doge Coin", "sig_figs": 4}
			]
		}
	`

	var opts weave.Options
	if err := json.Unmarshal([]byte(genesis), &opts); err != nil {
		t.Fatalf("cannot unmarshal genesis: %s", err)
	}

	db := store.MemStore()
	var ini Initializer
	if err := ini.FromGenesis(opts, db); err != nil {
		t.Fatalf("cannot load genesis: %s", err)
	}

	bucket := NewTokenInfoBucket()
	obj, err := bucket.Get(db, "MCR")
	if err != nil {
		t.Fatalf("cannot fetch token information: %s", err)
	} else if obj == nil {
		t.Fatal("token information not found")
	}

	info := obj.Value().(*TokenInfo)
	if info.Name != "my currency" {
		t.Errorf("invalid token name: %q", info.Name)
	}
	if info.SigFigs != 9 {
		t.Errorf("invalid token sig figs: %d", info.SigFigs)
	}
}
