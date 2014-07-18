/* Holds commonly used variables and configuration */

goog.provide('common');

console.log('common loaded');
var GAME = {};
GAME.Define = {
    LEFT: 1,
    DOWN: 2,
    UP: 4,
    RIGHT: 8,
    SELF: 16,
    CHANNEL: 32,
    ALL: 63,
    DIRS: 15
};

// /* Holds some convenience functions for manipulating DOM */
// export module DOMRoutines {
//     /* Returns the element with the given ID */
//     export function GetById(id: string) {
//         return document.getElementById(id)
//     }

//     /* Creates an element */
//     export function CreateElement(type, id, classname): HTMLElement {
//         var element = document.createElement(type)

//         if (id != null) {
//             element.id = id
//         }

//         if (classname != null) {
//             element.className = classname
//         }

//         return element
//     }

//     /* Creates a div */
//     export function CreateDiv(id, classname): HTMLDivElement {
//         return <HTMLDivElement>CreateElement('div', id, classname)
//     }

//     /* Creates an image */
//     export function CreateImg(id, classname, src): HTMLImageElement {
//         var img = <HTMLImageElement>CreateElement("img", id, classname)
//         img.src = src
//         return img
//     }

//     /* Removes all of the child elements of an element. */
//     export function RemoveChildren(element) {
//         while (element.firstChild) {
//             element.removeChild(element.firstChild)
//         }
//     }

//     /* Removes an element from the DOM. */
//     export function RemoveElement(element) {
//         if (element.parentNode) {
//             element.parentNode.removeChild(element)
//         }
//     }

//     /* Adds a classname to an element if not already present. */
//     export function AddClass(element, className) {
//         var classNames = element.className.split(' ')
//         if (!(className in classNames)) {
//             element.className += ' ' + className
//         }
//     }

//     /* Removes a classname from an element. */
//     export function RemoveClass(element, className) {
//         var classNames = element.className.split(' ')
//         var classIndex = classNames.indexOf(className)

//         if (classIndex >= 0) {
//             classNames.splice(classIndex, 1)
//             element.className = classNames.join(' ')
//         }
//     }

//     /* Adds a class to an element if it's not there or removes it if it is. */
//     export function ToggleClass(element, className) {
//        var classNames = element.className.split(' ')
//        var classIndex = classNames.indexOf(className)

//        if (classIndex >= 0) {
//           classNames.splice(classIndex, 1)
//        } else {
//           classNames.push(className)
//        }

//        element.className = classNames.join(' ')
//     }
// }

// /* Encapsulates the connection to the server websocket */
// export class Connection {

//     connection: WebSocket
//     // TODO define serverState more precisely
//     onServerTick: (serverState:any) => void
//     onConnect: () => void
//     player: number
//     debugPrefix

//     constructor(debugPrefix: string,
//                        host: string,
//                      player: number,
//                onServerTick: (serverState: any) => void,
//                   onConnect: () => void) {

//         this.player = player
//         this.connection = new WebSocket(host)
//         this.connection.onopen = this.OnOpen
//         this.connection.onclose = this.OnClose
//         this.connection.onerror = this.OnError
//         this.connection.onmessage = this.OnMessage
//         this.onServerTick = onServerTick
//         this.onConnect = onConnect
//         this.debugPrefix = debugPrefix
//     }

//     /* Sends an object as JSON via the websocket */
//     Send = (object: any) => {
//         object.Player = this.player
//         var msg = JSON.stringify(object)
//         console.log(this.debugPrefix + ": sending message " + msg)
//         this.connection.send(msg)
//     }

//     /* Sends the server the message that this client is ready to start */
//     SendStartGame = () => {
//         this.Send({CommandCode: 0})
//     }

//     /* Sends the server the given tick number */
//     SendTick = (tick) => {
//         this.Send({CommandCode: 1, TickNum: tick})
//     }

//     /* Is run when the websocket connection is opened */
//     OnOpen = (event) => {
//         console.log(this.debugPrefix + ": websocket connection opened ", event)

//         if (this.onConnect) {
//             this.onConnect()
//         }
//     }

//     /* Is run when the websocket connection is closed */
//     OnClose = (event) => {
//         console.log(this.debugPrefix + ": websocket connection closed ", event)
//     }

//     /* Is run when the websocket connection has on error */
//     OnError = (event) => {
//         console.log(this.debugPrefix + ": websocket connection error ", event)
//     }

//     /* Is run when the websocket connection has receieved a message */
//     OnMessage = (event) => {
//         var serverState = JSON.parse(event.data)
//         console.log(this.debugPrefix + ": websocket recieved message ", serverState)

//         if (this.onServerTick) {
//             this.onServerTick(serverState)
//         }
//     }
// }
