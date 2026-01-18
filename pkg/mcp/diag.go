package mcp

import (
	"context"

	"github.com/chwiee/mcp-k8s/internal/diag"
	"github.com/chwiee/mcp-k8s/internal/k8s"
)

func DiagnoseNamespace(namespace string) (*diag.Result, error) {
	client, err := k8s.NewClient()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return diag.Analyze(ctx, client, namespace)
}
