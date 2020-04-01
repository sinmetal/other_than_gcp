package backend

import (
	"context"
	"testing"
)

func TestSheetsService_ReadTest(t *testing.T) {
	ctx := context.Background()

	ss := NewSheetsService(ctx)
	_, err := ss.ReadTest(ctx, "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms", "A1")
	if err != nil {
		t.Fatal(err)
	}
}
