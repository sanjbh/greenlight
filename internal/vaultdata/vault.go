package vaultdata

import (
	"crypto/tls"
	"errors"
	"net/http"

	vault "github.com/hashicorp/vault/api"
)

type DbData struct {
	Username string
	Password string
	DBHost   string
}

func GetDataFromVault(address string) (*DbData, error) {
	config := vault.DefaultConfig()

	config.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	config.Address = address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}

	secret, err := client.Logical().Read("secret/data/greenlight/dbconfig")
	if err != nil {
		return nil, err
	}

	if secret == nil {
		return nil, errors.New("secret not found")
	}

	// fmt.Printf("%+v\n", secret)

	data, ok := secret.Data["data"].(map[string]any)
	if !ok {
		return nil, errors.New("failed to get data from secret")
	}

	return &DbData{
		Username: data["username"].(string),
		Password: data["password"].(string),
		DBHost:   data["dbhost"].(string),
	}, nil

}
