package authentication

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type AzureAuthenticator struct {
	Credential *azidentity.DefaultAzureCredential
}

func (a *AzureAuthenticator) Init() error {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}

	a.Credential = cred
	return nil
}
