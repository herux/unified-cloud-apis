package authentication

import (
	"context"

	"cloud.google.com/go/storage"
)

type GCPAuthenticator struct {
	Client *storage.Client
}

func (g *GCPAuthenticator) Init() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	g.Client = client

	return nil
}
