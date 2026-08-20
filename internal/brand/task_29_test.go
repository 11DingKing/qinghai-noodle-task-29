package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask29(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	p := ProductListing{ID: "p", StoreID: "store-1", SKU: "yak", OriginRegion: "青海", IngredientLotIDs: []string{"l"}, OnlinePriceCents: 1000, StorePriceCents: 1000}
	_, err := s.PublishStoreCatalog(context.Background(), compliantStore(now), activeLicense(now), []ProductListing{p}, map[string]IngredientLot{"l": {ID: "l", OriginRegion: "青海"}})
	require.NoError(t, err)
}
