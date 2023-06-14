package tlsgo_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/robin-samuel/tlsgo"
)

func TestClient(t *testing.T) {
	session := tlsgo.NewSession(tlsgo.Chrome112, false)
	session.Proxy = "http://localhost:8888"
	params := url.Values{"q": {"golang"}}
	response, err := session.Get("https://www.robinsamuel.dev", &tlsgo.Options{Params: params})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response.Request.Raw)
	fmt.Println(response.Raw)
	t.Log(response.StatusCode)
}
