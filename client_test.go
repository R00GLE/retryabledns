package retryabledns

import (
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/require"
)

func TestConsistentResolve(t *testing.T) {
	client := New(Options{BaseResolvers: []string{"8.8.8.8:53", "1.1.1.1:53"}, MaxRetries: 5})

	var last string
	for i := 0; i < 10; i++ {
		d, err := client.Resolve("example.com")
		require.Nil(t, err, "could not resolve dns")

		if last != "" {
			require.Equal(t, last, d.A[0], "got another data from previous")
		} else {
			last = d.A[0]
		}
	}
}

func TestQueryMultiple(t *testing.T) {
	client := New(Options{BaseResolvers: []string{"8.8.8.8:53", "1.1.1.1:53"}, MaxRetries: 5})

	d, _ := client.QueryMultiple("example.com", []uint16{dns.TypeA, dns.TypeAAAA})

	// From current dig result
	require.True(t, len(d.A) > 0)
	require.True(t, len(d.AAAA) > 0)
}

func TestTrace(t *testing.T) {
	client := New(Options{BaseResolvers: []string{"8.8.8.8:53", "1.1.1.1:53"}, MaxRetries: 5})

	_, err := client.Trace("www.projectdiscovery.io", dns.TypeA, 100)
	require.Nil(t, err, "could not resolve dns")
}
