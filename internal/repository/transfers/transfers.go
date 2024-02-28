package transfers

import (
	"encoding/json"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/rabbit"
	"user_service/internal/repository"

	"github.com/andReyM228/lib/log"
)

type Repository struct {
	log    log.Logger
	rabbit rabbit.Rabbit
}

func NewRepository(rabbit rabbit.Rabbit, log log.Logger) Repository {
	return Repository{
		rabbit: rabbit,
		log:    log,
	}
}

func (r Repository) Issue(ToAddress, Memo string, Amount int64) (string, error) {
	result, err := r.rabbit.Request(bus.SubjectTxServiceIssue, bus.IssueRequest{
		ToAddress: ToAddress,
		Amount:    Amount,
		Memo:      Memo,
	})
	if err != nil {
		return "", err
	}

	if result.StatusCode != 200 {
		return "", repository.InternalServerError{}
	}

	var txResp bus.TxResponse

	if err := json.Unmarshal(result.Payload, &txResp); err != nil {
		return "", repository.InternalServerError{}
	}

	return txResp.TxHash, nil
}

func (r Repository) Withdraw(ToAddress, Memo string, Amount int64) (string, error) {
	result, err := r.rabbit.Request(bus.SubjectTxServiceWithdraw, bus.WithdrawRequest{
		ToAddress: ToAddress,
		Amount:    Amount,
		Memo:      Memo,
	})
	if err != nil {
		return "", err
	}

	if result.StatusCode != 200 {
		return "", repository.InternalServerError{}
	}

	var txResp bus.TxResponse

	if err := json.Unmarshal(result.Payload, &txResp); err != nil {
		return "", repository.InternalServerError{}
	}

	return txResp.TxHash, nil
}
