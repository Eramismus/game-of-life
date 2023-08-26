package graphics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/eramismus/game-of-life/src/gol"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

func StartWebSocketServer(
	port int,
	endpoint string,
	grid gol.Grid,
	neighbour_map map[gol.Position][]gol.Position,
	json_grid gol.JsonGrid,
) {
	fmt.Println("Starting a server")

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		// fmt.Println(&grid)
		go func(grid *gol.Grid, neighbour_map map[gol.Position][]gol.Position, json_grid *gol.JsonGrid) {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				fmt.Printf("Message: %s\n", msg)
				fmt.Printf("OpCode: %b\n", op)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				for i := 0; i < 100; i++ {
					gol.UpdateGrid(grid, neighbour_map, json_grid)
					data, err := json.Marshal(json_grid)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
						return
					}
					fmt.Printf("Sending back: %s\n", data)
					err = wsutil.WriteServerMessage(conn, ws.OpText, data)
					if err != nil {
						return
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}(&grid, neighbour_map, &json_grid)
	})

	port_string := fmt.Sprintf(":%d", port)
	fmt.Println("Started a web-socket server")
	err := http.ListenAndServe(port_string, nil)
	if err != nil {
		panic(err)
	}
}
