package car_trading

import (
	"context"
	"github.com/andReyM228/lib/errs"
	"github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (s Service) getTx(ctx context.Context, txHash string) (*bankTypes.MsgSend, error) {
	tx, err := s.chain.GetTx(ctx, txHash)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	var msg types.Msg
	if err := s.chain.GetUnpacker().UnpackAny(tx.Tx.Body.Messages[0], &msg); err != nil {
		return nil, errs.BadRequestError{Cause: "invalid msg format"}
	}

	msgSend := msg.(*bankTypes.MsgSend)

	//msg, ok := tx.GetTx().GetMsgs()[0].(*types.MsgSend)
	//if !ok {
	//	return nil, errs.BadRequestError{Cause: "invalid msg format"}
	//}

	return msgSend, nil
}
