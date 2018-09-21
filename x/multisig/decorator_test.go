package multisig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x"
)

func TestDecorator(t *testing.T) {
	var helpers x.TestHelpers
	db := store.MemStore()

	// create some keys
	_, a := helpers.MakeKey()
	_, b := helpers.MakeKey()
	_, c := helpers.MakeKey()

	// the contract we'll be using in our tests
	contractID := withContract(t, db, CreateContractMsg{
		Sigs:                newSigs(a, b, c),
		ActivationThreshold: 2,
		AdminThreshold:      3,
	})

	// helper to create a ContractTx
	multisigTx := func(payload, multisig []byte) ContractTx {
		tx := helpers.MockTx(helpers.MockMsg(payload))
		return ContractTx{Tx: tx, MultisigID: multisig}
	}

	cases := []struct {
		tx      weave.Tx
		signers []weave.Condition
		perms   []weave.Condition
		err     error
	}{
		// doesn't support multisig interface
		{
			helpers.MockTx(helpers.MockMsg([]byte{1, 2, 3})),
			[]weave.Condition{a},
			nil,
			nil,
		},
		// Correct interface but no content
		{
			multisigTx([]byte("john"), nil),
			[]weave.Condition{a},
			nil,
			nil,
		},
		// with multisig contract
		{
			multisigTx([]byte("foo"), contractID),
			[]weave.Condition{a, b},
			[]weave.Condition{MultiSigCondition(contractID)},
			nil,
		},
		// with multisig contract but not enough signatures to activate
		{
			multisigTx([]byte("foo"), contractID),
			[]weave.Condition{a},
			nil,
			ErrUnauthorizedMultiSig(contractID),
		},
		// with invalid multisig contract ID
		{
			multisigTx([]byte("foo"), []byte("bad id")),
			[]weave.Condition{a, b},
			nil,
			ErrContractNotFound([]byte("bad id")),
		},
	}

	// the handler we're chaining with the decorator
	h := new(MultisigCheckHandler)
	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			ctx, auth := newContextWithAuth(tc.signers...)

			d := NewDecorator(auth)
			stack := helpers.Wrap(d, h)

			_, err := stack.Check(ctx, db, tc.tx)
			if tc.err != nil {
				require.EqualError(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.perms, h.Perms)
			}

			_, err = stack.Deliver(ctx, db, tc.tx)
			if tc.err != nil {
				require.EqualError(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.perms, h.Perms)
			}
		})
	}
}

//---------------- helpers --------

// MultisigCheckHandler stores the seen permissions on each call
// for this extension's authenticator (ie. multisig.Authenticate)
type MultisigCheckHandler struct {
	Perms []weave.Condition
}

var _ weave.Handler = (*MultisigCheckHandler)(nil)

func (s *MultisigCheckHandler) Check(ctx weave.Context, store weave.KVStore,
	tx weave.Tx) (res weave.CheckResult, err error) {
	s.Perms = Authenticate{}.GetConditions(ctx)
	return
}

func (s *MultisigCheckHandler) Deliver(ctx weave.Context, store weave.KVStore,
	tx weave.Tx) (res weave.DeliverResult, err error) {
	s.Perms = Authenticate{}.GetConditions(ctx)
	return
}

// ContractTx fulfills the MultiSigTx interface to satisfy the decorator
type ContractTx struct {
	weave.Tx
	MultisigID []byte
}

var _ MultiSigTx = ContractTx{}
var _ weave.Tx = ContractTx{}

func (p ContractTx) GetMultisig() []byte {
	return p.MultisigID
}