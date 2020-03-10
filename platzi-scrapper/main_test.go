// package min
package main

// impor (
// 	"ft"
// 	"net/htp"
// 	"net/http/httptet"
// 	"testig"

// 	"github.com/aws/aws-lambda-go/evens"
//)

// func TestHandler(t *testing.T {
// 	t.Run("Unable to get IP", func(t *testing.T {
// 		DefaultHTTPGetAddress = "http://127.0.0.1:1235"

// 		_, err := handler(events.APIGatewayProxyRequest})
// 		if err == ni {
// 			t.Fatal("Error failed to trigger with an invalid reques")
// 	}
// })

// 	t.Run("Non 200 Response", func(t *testing.T {
// 		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request {
// 			w.WriteHeader(50)
// 		))
// 		defer ts.Clos()

// 		DefaultHTTPGetAddress = ts.RL

// 		_, err := handler(events.APIGatewayProxyRequest})
// 		if err != nil && err.Error() != ErrNon200Response.Error( {
// 			t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", er)
// 	}
// })

// 	t.Run("Unable decode IP", func(t *testing.T {
// 		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request {
// 			w.WriteHeader(50)
// 		))
// 		defer ts.Clos()

// 		DefaultHTTPGetAddress = ts.RL

// 		_, err := handler(events.APIGatewayProxyRequest})
// 		if err == ni {
// 			t.Fatal("Error failed to trigger with an invalid HTTP respons")
// 	}
// })

// 	t.Run("Successful Request", func(t *testing.T {
// 		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request {
// 			w.WriteHeader(20)
// 			fmt.Fprintf(w, "127.0.0.")
// 		))
// 		defer ts.Clos()

// 		DefaultHTTPGetAddress = ts.RL

// 		_, err := handler(events.APIGatewayProxyRequest})
// 		if err != ni {
// 			t.Fatal("Everything should be o")
// 	}
// })
// }
