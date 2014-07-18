// These values are filled in the HTML file
goog.require("common");
goog.provide("hand");
// import Common = require("Common")
// var DOM = Common.DOMRoutines
// var Defines = Common.Definitions

// var BUTTON_PREFIX = '/resources/img/card/btn_'
// var HAND = []

// class CardButton {

//     buttonName: string
//     pressed: boolean = false
//     root: HTMLImageElement

//     constructor(type, onClick) {
//         var buttonName = BUTTON_PREFIX + type

//         if (!onClick) {
//             this.root = DOM.CreateImg('', 'button', buttonName + '_disable.png')
//             return
//         }

//         this.root = DOM.CreateImg('', 'button', buttonName + '.png')

//         this.root.onmousedown = function() {
//             this.root.src = buttonName + '_press.png'
//             this.pressed = true
//         }

//         this.root.onmouseout = function() {
//             this.root.src = buttonName + '.png'
//             this.pressed = false
//         }

//         this.root.onmouseup = function() {
//             if (this.pressed) {
//                 this.pressed = false
//                 this.root.onmouseout()
//                 onClick()
//             }
//         }
//     }
// }

// class Card {

//     root: HTMLDivElement
//     container: HTMLDivElement
//     artImg: HTMLImageElement
//     maskImg: HTMLImageElement
//     cardTitle: HTMLDivElement
//     cardCount: HTMLDivElement
//     buttons: CardButton[] = []

//     constructor(name, actions, image, count) {

//         this.root = DOM.CreateDiv('', 'card-wrapper')
//         this.container = DOM.CreateDiv('', 'card')
//         this.root.appendChild(this.container)
//         this.artImg = DOM.CreateImg('', 'art', image)
//         this.maskImg = DOM.CreateImg('', 'mask', '/resources/img/card/card.png')
//         this.cardTitle = DOM.CreateDiv('', 'title')
//         this.cardTitle.innerHTML = name
//         this.cardCount = DOM.CreateDiv('', 'count')
//         this.cardCount.innerHTML = count
//         this.container.appendChild(this.artImg)
//         this.container.appendChild(this.maskImg)
//         this.container.appendChild(this.cardTitle)
//         this.container.appendChild(this.cardCount)

//         this.buttons.push(new CardButton('up',
//             actions & Defines.CardActions.UP ? this.ClickUp : null))
//         this.buttons.push(new CardButton('left',
//             actions & Defines.CardActions.LEFT ? this.ClickLeft : null))
//         this.buttons.push(new CardButton('right',
//             actions & Defines.CardActions.RIGHT ? this.ClickRight : null))
//         this.buttons.push(new CardButton('down',
//             actions & Defines.CardActions.DOWN ? this.ClickDown : null))
//         this.buttons.push(new CardButton('self',
//             actions & Defines.CardActions.SELF ? this.ClickSelf : null))
//         this.buttons.push(new CardButton('channel',
//             actions & Defines.CardActions.CHANNEL ? this.ClickChannel : null))

//         for (var i = 0; i < this.buttons.length; i++) {
//             var buttonRoot = this.buttons[i].root
//             buttonRoot.style.top = (178 + 147 * i) + 'px'
//             this.container.appendChild(buttonRoot)
//         }

//         this.container.onclick = function() {
//             if (HAND[HAND.length - 1] != self) {
//                 var before = []
//                 for (var i = 0; HAND[i] != self; i++) {
//                     before.push(HAND[i])
//                 }

//                 var after = []
//                 for (i++; i < HAND.length; i++) {
//                     after.push(HAND[i])
//                 }

//                 HAND = before.concat(after)
//                 HAND.push(self)
//                 this.ArrangeHand()
//             }
//         }
//     }

//     ClickLeft = () => {
//     }

//     ClickDown = () => {
//     }

//     ClickUp = () => {
//     }

//     ClickRight = () => {
//     }

//     ClickSelf = () => {
//     }

//     ClickChannel = () => {
//     }
// }

// class Main {

//     clientTick: number = 0
//     containerDiv: HTMLElement = DOM.GetById('cards')
//     connection: Common.Connection

//     constructor() {
//         this.connection = new Common.Connection("Hand",
//             GAME.HOST, GAME.PLAYER, this.HandleServerTick, null)

//         window.onresize = this.WindowOnResize

//         DOM.GetById("start-link").onclick = () => {
//             this.WindowOnResize()
//             DOM.RemoveElement(DOM.GetById("start-link"))
//             this.connection.SendStartGame()
//             this.connection.SendTick(++this.clientTick)
//         }
//     }

//     WindowOnResize = () => {
//         this.containerDiv.style.zoom = "" + window.innerWidth / 780.
//     }

//     BuildHand = () => {
//         DOM.RemoveChildren(this.containerDiv)
//         for (var i = 0; i < HAND.length; i++) {
//             this.containerDiv.appendChild(HAND[i].root)
//         }
//         this.ArrangeHand()
//     }

//     ArrangeHand = () => {
//         for (var i = 0; i < HAND.length; i++) {
//             var style = HAND[i].root.style
//             style.left = (4 * i) + 'px'
//             style.top = (90 * i) + 'px'
//             style.zIndex = i
//         }
//     }

//     HandleServerTick = (serverState) => {
//         if (serverState.Players) {
//             HAND = []
//             var actions = serverState.Players[GAME.PLAYER].AvailableActions

//             for (var i = 0; i < actions.length; i++) {
//                 var action = actions[i]
//                 var dirs = Defines.CardActions.CHANNEL

//                 if (action.AvailableDirections[0]) {
//                     dirs |= Defines.CardActions.LEFT
//                 }

//                 if (action.AvailableDirections[1]) {
//                     dirs |= Defines.CardActions.RIGHT
//                 }

//                 if (action.AvailableDirections[2]) {
//                     dirs |= Defines.CardActions.UP
//                 }

//                 if (action.AvailableDirections[3]) {
//                     dirs |= Defines.CardActions.DOWN
//                 }

//                 if (action.AvailableDirections[4]) {
//                     dirs |= Defines.CardActions.SELF
//                 }

//                 HAND.push(new Card(action.ShortDescription, dirs,
//                     '/resources/img/card/card_art_001.png', '&infin'))
//             }
//             this.BuildHand()
//         }
//     }
// }
// new Main()
