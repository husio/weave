package distribution

import (
	"fmt"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

// Initializer fulfils the Initializer interface to load data from the genesis
// file
type Initializer struct{}

var _ weave.Initializer = (*Initializer)(nil)

// FromGenesis will parse initial account info from genesis and save it to the
// database
func (*Initializer) FromGenesis(opts weave.Options, db weave.KVStore) error {
	type recipient struct {
		Address weave.Address `json:"address"`
		Weight  int32         `json:"weight"`
	}
	var revenues []struct {
		Admin      weave.Address `json:"admin"`
		Recipients []recipient   `json:"recipients"`
	}
	if err := opts.ReadOptions("distribution", &revenues); err != nil {
		return errors.Wrap(err, "cannot load distribution")
	}

	bucket := NewRevenueBucket()
	for i, r := range revenues {
		recipients := make([]*Recipient, 0, len(r.Recipients))
		for _, rc := range r.Recipients {
			recipients = append(recipients, &Recipient{
				Address: rc.Address,
				Weight:  rc.Weight,
			})
		}
		revenue := Revenue{
			Admin:      r.Admin,
			Recipients: recipients,
		}
		if err := revenue.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("revenue #%d is invalid", i))
		}
		if _, err := bucket.Create(db, &revenue); err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot store #%d revenue", i))
		}
	}
	return nil
}
