package tlsgo_test

import (
	"testing"

	"github.com/robin-samuel/tlsgo"
)

func TestClient(t *testing.T) {
	session := tlsgo.NewSession(tlsgo.Chrome112)
	response, err := session.Get("https://www.robinsamuel.dev", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Status)
}
