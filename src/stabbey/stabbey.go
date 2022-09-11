package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"

	"stabbey/game"
	"stabbey/interfaces"
	"stabbey/order"
	"stabbey/player"
	"stabbey/runtime"
	"stabbey/serializable"
	"stabbey/signalhandlers"
	"stabbey/spectator"
)

/* Constants */
const FILE_COMBINED_HTML string = "resources/html/combined.html"
const FILE_HAND_HTML string = "resources/html/hand.html"
const FILE_MAP_HTML string = "resources/html/map.html"
const FILE_SETUP_HTML string = "resources/html/setup.html"
const HTTP_ROOT string = "/"
const HTTP_CONNECT string = "/connect"
const HTTP_CONNECT_HAND string = "/connect/hand"
const HTTP_CONNECT_MAP string = "/connect/map"
const HTTP_WEBSOCKET string = "/ws"
const HTTP_WEBSOCKET_SPECTATOR string = "/ws-spectate"

/* JS to use for compiled mode of map.html */
const SCRIPT_COMPILED_MAP string = "" +
	"<script src=\"/resources/js/compiled/map.js\"></script>\n"

/* JS to use for uncompiled mode of map.html */
const SCRIPT_UNCOMPILED_MAP string = "" +
	"<script src=\"/resources/js/3rd-party/base.js\"></script>\n" +
	"<script type=\"text/javascript\"> goog.require(\"st.map\"); </script>\n"

/* JS to use for compiled mode of hand.html */
const SCRIPT_COMPILED_HAND string = "" +
	"<script src=\"/resources/js/compiled/hand.js\"></script>\n"

/* JS to use for uncompiled mode of hand.html */
const SCRIPT_UNCOMPILED_HAND string = "" +
	"<script src=\"/resources/js/3rd-party/base.js\"></script>\n" +
	"<script type=\"text/javascript\"> goog.require(\"st.hand\"); </script>\n"

/* CLI Flags */
var ADDR = flag.String("addr", ":8080", "http service address")
var COMPILED_JS = flag.Bool("compiledjs", false, "run JS in compiled mode")

/* Important Globals */
var GAME *game.Game
var RUNTIME *runtime.Runtime

/* Template for hand.html */
type HandPageTemplate struct {
	Me       int
	Host     string
	JsToLoad string
}

/* Template for map.html */
type MapPageTemplate struct {
	Host     string
	JsToLoad string
}

/* Start the server and connect url paths to functions */
func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	go signalhandlers.Interrupt()
	go signalhandlers.Quit()

	http.HandleFunc("/resources/js/", JavascriptHandler)
	http.HandleFunc("/resources/css/", CssHandler)
	http.HandleFunc("/resources/img/", ImgHandler)
	http.HandleFunc(HTTP_ROOT, InitSetup)
	http.HandleFunc(HTTP_CONNECT, ConnectSetup)
	http.HandleFunc(HTTP_CONNECT_HAND, ConnectHandSetup)
	http.HandleFunc(HTTP_CONNECT_MAP, ConnectMapSetup)
	http.Handle(HTTP_WEBSOCKET, websocket.Handler(PlayerConnect))
	http.Handle(HTTP_WEBSOCKET_SPECTATOR, websocket.Handler(SpectatorConnect))

	log.Println("Starting Server")
	GAME = game.NewGame()
	RUNTIME = runtime.New(GAME)

	if err := http.ListenAndServe(*ADDR, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

/* Send the basic page that just has the buttons to start/join games */
func InitSetup(w http.ResponseWriter, req *http.Request) {
	if tmpl, e := template.ParseFiles(FILE_SETUP_HTML); e != nil {
		log.Fatal("Parse error:", e)
	} else {
		tmpl.Execute(w, nil)
	}
}

/* Serve javascript files manually in order to set the content type */
func JavascriptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, r.URL.Path[1:])
}

/* Serve javascript files manually in order to set the content type */
func CssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, r.URL.Path[1:])
}

/* Serve image files manually in order to set the content type */
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, r.URL.Path[1:])
}

/* Connecting to the combined map and hand view */
func ConnectSetup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, FILE_COMBINED_HTML)
}

/* Create the game, add players as they join */
func ConnectHandSetup(w http.ResponseWriter, r *http.Request) {
	var curPlayer *player.Player

	curPlayer = player.New(GAME)
	pX, pY := GAME.GetRandomEmptySpace()
	curPlayer.SetPosition(0, pX, pY)
	GAME.AddPlayer(curPlayer, curPlayer)
	log.Printf("Added player %v to game at %v, %v ", curPlayer.GetPlayerId(),
		pX, pY)

	if tmpl, e := template.ParseFiles(FILE_HAND_HTML); e != nil {
		log.Fatal("Parse error:", e)
	} else {
		var script string

		if *COMPILED_JS {
			script = SCRIPT_COMPILED_HAND
		} else {
			script = SCRIPT_UNCOMPILED_HAND
		}

		tmpl.Execute(w, HandPageTemplate{curPlayer.GetPlayerId(),
			"ws://" + r.Host + HTTP_WEBSOCKET, script})
	}
}

/* Connecting to the map does not count as a player connecting */
func ConnectMapSetup(w http.ResponseWriter, r *http.Request) {
	if tmpl, e := template.ParseFiles(FILE_MAP_HTML); e != nil {
		log.Fatal("Parse error:", e)
	} else {
		var script string

		if *COMPILED_JS {
			script = SCRIPT_COMPILED_MAP
		} else {
			script = SCRIPT_UNCOMPILED_MAP
		}

		tmpl.Execute(w, MapPageTemplate{"ws://" +
			r.Host + HTTP_WEBSOCKET_SPECTATOR,
			script})
	}
}

/* Initialize a websocket connection and pair it with the handshaking player */
func PlayerConnect(ws *websocket.Conn) {
	websocketConnection := struct {
		CommandCode int
		Player      int
	}{}

	var message string
	if err := websocket.Message.Receive(ws, &message); err != nil {
		log.Fatal("Reading Socket Error:", err)
	} else {
		err := json.Unmarshal([]byte(message), &websocketConnection)
		if err != nil {
			log.Fatal("Decoding Message Error:", err)
		} else {
			log.Printf("WebSocket Connected: %+v", websocketConnection)
			p := GAME.GetPlayer(websocketConnection.Player)
			p.SetWebSocketConnection(ws)
			p.SendMessage(serializable.NewGameInfo().Json())
			PlayerKeepReading(p, ws)
		}
	}
}

/* Keep the player's websocket alive and continue reading from it forever */
func PlayerKeepReading(p interfaces.Player, ws *websocket.Conn) {
	for {
		var message string
		playerOrder := struct {
			CommandCode int
			TickNum     int
			Queue       []string
		}{}

		if e := websocket.Message.Receive(ws, &message); e != nil {
			break
		} else {
			log.Println("Received message on websocket:", message)
			err := json.Unmarshal([]byte(message), &playerOrder)
			if err != nil {
				log.Fatalln("Error Decoding Order: ", string(message))
				break
			}

			o := order.NewOrder(playerOrder.CommandCode, playerOrder.TickNum,
				playerOrder.Queue, p, GAME)
			RUNTIME.AddOrder(o)
		}
	}
	log.Printf("Closing socket for %v", p.GetPlayerId())
	ws.Close()
}

/* Initialize a spectator connection; disconnected from players */
func SpectatorConnect(ws *websocket.Conn) {
	var message string
	if err := websocket.Message.Receive(ws, &message); err != nil {
		log.Fatal("Reading Socket Error:", err)
	} else {
		spec := spectator.New()
		spec.SetWebSocketConnection(ws)
		spec.SendMessage(serializable.NewGameInfo().Json())
		GAME.AddSpectator(spec)
		log.Println("Adding board spectator")

		/* Keep reading */
		for {
			if websocket.Message.Receive(ws, &message) != nil {
				break
			}
		}
	}

	/* TODO should probably remove spectators from the game, or limit how many
	 * can be added */
	log.Printf("Closing spectator websocket")
	ws.Close()
}
