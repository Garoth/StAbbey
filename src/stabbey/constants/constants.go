package constants

/* TODO this package could probably not exist (constants could be moved to
 * their respective subpackages and used only there) */

const FILE_SETUP_HTML string = "resources/html/setup.html"
const FILE_MAIN_HTML string  = "resources/html/main.html"

const HTTP_ROOT string       = "/"
const HTTP_CONNECT string    = "/connect"
const HTTP_WEBSOCKET string  = "/ws"
const HTTP_TEST string       = "/test"

const FORMVAL_GAMEKEY string = "gamekey"

const BOARD_WIDTH int        = 16
const BOARD_HEIGHT int       = 12
const BOARD_NUM_LAYERS int   = 8

const ENTITY_TYPE_PLAYER     = "player"
const ENTITY_TYPE_MONSTER    = "monster"
