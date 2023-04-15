package tlsgo_test

import (
	"fmt"
	"testing"

	"github.com/robin-samuel/tlsgo"
)

func TestClient(t *testing.T) {
	session := tlsgo.NewSession(tlsgo.Chrome112)
	session.Proxy = "http://localhost:8888"
	response, err := session.Get("https://www.robinsamuel.dev", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response.Request.Raw)
	fmt.Println(response.Raw)
	t.Log(response.StatusCode)
}
