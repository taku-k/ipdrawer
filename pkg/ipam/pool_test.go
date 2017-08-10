package ipam

import (
	"net"
	"testing"

	"golang.org/x/net/context"

	"github.com/taku-k/ipdrawer/pkg/storage"
)

func Test_getPools(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := newTestIPManager(r)

	ctx := context.Background()

	pool := &IPPool{
		Start:  net.ParseIP("10.0.0.1"),
		End:    net.ParseIP("10.0.0.254"),
		Status: POOL_AVAILABLE,
	}

	n := &Network{
		Prefix: &net.IPNet{
			IP:   net.ParseIP("10.0.0.0"),
			Mask: net.CIDRMask(24, 32),
		},
		Status: NETWORK_AVAILABLE,
	}

	_ = m.CreateNetwork(ctx, n)
	_ = m.CreatePool(ctx, n, pool)

	pools, err := getPools(r, n)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(pools) == 0 {
		t.Fatalf("Must be not zero")
	}
	if !pools[0].Start.Equal(pool.Start) || !pools[0].End.Equal(pool.End) {
		t.Fatalf("Got unexpected pool: %v", pools[0])
	}
}