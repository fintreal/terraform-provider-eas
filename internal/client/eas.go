package client

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/eas"
)

type EASClient struct {
	eas.EASClient
	AccountId   string
	AccountName string
}

func NewEASClient(token string, accountName string) (*EASClient, error) {
	client := eas.NewEASClient(token)

	me, err := client.Me.Get()

	if me == nil || err != nil {
		return nil, fmt.Errorf("token is invalid")
	}

	account, err := client.Account.GetByName(accountName)

	if err != nil {
		return nil, err
	}

	return &EASClient{
		EASClient:   *client,
		AccountId:   account.Id,
		AccountName: account.Name,
	}, err
}
