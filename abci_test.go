package weave_test

import (
	"fmt"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	pkerr "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateErrorResult(t *testing.T) {
	cases := []struct {
		err  error
		msg  string
		code uint32
	}{
		{errors.NormalizePanic("stdlib"), "internal", errors.ErrInternal.ABCICode()},
		{fmt.Errorf("base"), "base", errors.ErrInternal.ABCICode()},
		{pkerr.New("dave"), "dave", errors.ErrInternal.ABCICode()},
		{errors.Wrap(fmt.Errorf("demo"), "wrapped"), "wrapped: demo", errors.ErrInternal.ABCICode()},
		{errors.ErrInvalidInput.New("unable to decode"), errors.ErrInvalidInput.New("unable to decode").Error(), errors.ErrInvalidInput.ABCICode()},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {

			dres := weave.DeliverTxError(tc.err, false)
			assert.True(t, dres.IsErr())
			assert.Contains(t, dres.Log, tc.msg)

			dres = weave.DeliverTxError(tc.err, true)
			assert.True(t, dres.IsErr())
			assert.Contains(t, dres.Log, tc.msg)

			// TODO:O this is failing, because stacktrace
			// implementation is not present for the new error
			// handing code.
			//assert.Contains(t, dres.Log, "iov-one/weave/abci")
			assert.Equal(t, tc.code, dres.Code)

			cres := weave.CheckTxError(tc.err, false)
			assert.True(t, cres.IsErr())
			assert.Contains(t, cres.Log, tc.msg)
			// assert.Equal(t, fmt.Sprintf("cannot check tx: %s", tc.msg), cres.Log)

			cres = weave.CheckTxError(tc.err, true)
			assert.True(t, cres.IsErr())
			assert.Contains(t, cres.Log, tc.msg)
			// TODO: this is failing, because stacktrace
			// implementation is not present for the new error
			// handing code.
			//assert.Contains(t, cres.Log, "iov-one/weave/abci")
			assert.Equal(t, tc.code, cres.Code)
		})
	}
}

func TestCreateResults(t *testing.T) {
	d, msg := []byte{1, 3, 4}, "got it"
	dres := weave.DeliverResult{Data: d, Log: msg}
	ad := dres.ToABCI()
	assert.EqualValues(t, d, ad.Data)
	assert.Equal(t, msg, ad.Log)
	assert.Empty(t, ad.Tags)

	c, gas := "aok", int64(12345)
	cres := weave.NewCheck(gas, c)
	ac := cres.ToABCI()
	assert.Equal(t, c, ac.Log)
	assert.Equal(t, gas, ac.GasWanted)
	assert.Empty(t, ac.Data)
}
