package types

import (
	"cosmossdk.io/errors"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequest{}

func NewMsgRequest(creator string, amount uint64) *MsgRequest {
	return &MsgRequest{
		Creator: creator,
		Amount:  amount,
	}
}

// ValidateBasic runs stateless checks on the message
func (msg *MsgRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	 // Ensure the amount is non-zero.
	 if msg.Amount == 0 {
        return errors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be greater than zero")
    }
	return nil
}

