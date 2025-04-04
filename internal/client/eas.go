package client

import "github.com/fintreal/eas-sdk-go/eas"

type EASClient struct {
	eas.EASClient
	AccountId   string
	AccountName string
}

func NewEASClient(token string, accountName string) (*EASClient, error) {
	client := eas.NewEASClient(token)
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
