package graphics

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func StartWebSocketServer(port int, endpoint string) {
	fmt.Println("Starting a server")
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					return
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					return
				}
			}
		}()
	})

	http.HandleFunc("/", serveMainPage)

	port_string := fmt.Sprintf(":%d", port)
	http.ListenAndServe(port_string, nil)
}
