package nft

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

const (
	PathAddApprovalMsg    = "nft/approval/add"
	PathRemoveApprovalMsg = "nft/approval/remove"
)

type ApprovalMsg interface {
	GetT() string
	Identified
}

func (*AddApprovalMsg) Path() string {
	return PathAddApprovalMsg
}

func (*RemoveApprovalMsg) Path() string {
	return PathRemoveApprovalMsg
}

func (m AddApprovalMsg) Validate() error {
	if err := weave.Address(m.Address).Validate(); err != nil {
		return err
	}
	if !isValidAction(m.Action) {
		return errors.ErrInternal.New("invalid action")
	}
	if !isValidTokenID(m.ID) {
		return errors.ErrInternal.New("invalid token ID")
	}
	return m.Options.Validate()
}

func (m RemoveApprovalMsg) Validate() error {
	if err := weave.Address(m.Address).Validate(); err != nil {
		return err
	}
	if !isValidAction(m.Action) {
		return errors.ErrInternal.New("invalid action")
	}
	if !isValidTokenID(m.ID) {
		return errors.ErrInternal.New("invalid token ID")
	}
	return nil
}
