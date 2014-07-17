define(["require", "exports"], function(require, exports) {
    /* Holds commonly used variables and configuration */
    (function (Definitions) {
        (function (CardActions) {
            CardActions[CardActions["LEFT"] = 1] = "LEFT";
            CardActions[CardActions["DOWN"] = 2] = "DOWN";
            CardActions[CardActions["UP"] = 4] = "UP";
            CardActions[CardActions["RIGHT"] = 8] = "RIGHT";
            CardActions[CardActions["SELF"] = 16] = "SELF";
            CardActions[CardActions["CHANNEL"] = 32] = "CHANNEL";
            CardActions[CardActions["ALL"] = 63] = "ALL";
            CardActions[CardActions["DIRS"] = 15] = "DIRS";
        })(Definitions.CardActions || (Definitions.CardActions = {}));
        var CardActions = Definitions.CardActions;
    })(exports.Definitions || (exports.Definitions = {}));
    var Definitions = exports.Definitions;

    /* Holds some convenience functions for manipulating DOM */
    (function (DOMRoutines) {
        /* Returns the element with the given ID */
        function GetById(id) {
            return document.getElementById(id);
        }
        DOMRoutines.GetById = GetById;

        /* Creates an element */
        function CreateElement(type, id, classname) {
            var element = document.createElement(type);

            if (id != null) {
                element.id = id;
            }

            if (classname != null) {
                element.className = classname;
            }

            return element;
        }
        DOMRoutines.CreateElement = CreateElement;

        /* Creates a div */
        function CreateDiv(id, classname) {
            return CreateElement('div', id, classname);
        }
        DOMRoutines.CreateDiv = CreateDiv;

        /* Creates an image */
        function CreateImg(id, classname, src) {
            var img = CreateElement("img", id, classname);
            img.src = src;
            return img;
        }
        DOMRoutines.CreateImg = CreateImg;

        /* Removes all of the child elements of an element. */
        function RemoveChildren(element) {
            while (element.firstChild) {
                element.removeChild(element.firstChild);
            }
        }
        DOMRoutines.RemoveChildren = RemoveChildren;

        /* Removes an element from the DOM. */
        function RemoveElement(element) {
            if (element.parentNode) {
                element.parentNode.removeChild(element);
            }
        }
        DOMRoutines.RemoveElement = RemoveElement;

        /* Adds a classname to an element if not already present. */
        function AddClass(element, className) {
            var classNames = element.className.split(' ');
            if (!(className in classNames)) {
                element.className += ' ' + className;
            }
        }
        DOMRoutines.AddClass = AddClass;

        /* Removes a classname from an element. */
        function RemoveClass(element, className) {
            var classNames = element.className.split(' ');
            var classIndex = classNames.indexOf(className);

            if (classIndex >= 0) {
                classNames.splice(classIndex, 1);
                element.className = classNames.join(' ');
            }
        }
        DOMRoutines.RemoveClass = RemoveClass;

        /* Adds a class to an element if it's not there or removes it if it is. */
        function ToggleClass(element, className) {
            var classNames = element.className.split(' ');
            var classIndex = classNames.indexOf(className);

            if (classIndex >= 0) {
                classNames.splice(classIndex, 1);
            } else {
                classNames.push(className);
            }

            element.className = classNames.join(' ');
        }
        DOMRoutines.ToggleClass = ToggleClass;
    })(exports.DOMRoutines || (exports.DOMRoutines = {}));
    var DOMRoutines = exports.DOMRoutines;

    /* Encapsulates the connection to the server websocket */
    var Connection = (function () {
        function Connection(debugPrefix, host, player, onServerTick, onConnect) {
            var _this = this;
            /* Sends an object as JSON via the websocket */
            this.Send = function (object) {
                object.Player = _this.player;
                var msg = JSON.stringify(object);
                console.log(_this.debugPrefix + ": sending message " + msg);
                _this.connection.send(msg);
            };
            /* Sends the server the message that this client is ready to start */
            this.SendStartGame = function () {
                _this.Send({ CommandCode: 0 });
            };
            /* Sends the server the given tick number */
            this.SendTick = function (tick) {
                _this.Send({ CommandCode: 1, TickNum: tick });
            };
            /* Is run when the websocket connection is opened */
            this.OnOpen = function (event) {
                console.log(_this.debugPrefix + ": websocket connection opened ", event);

                if (_this.onConnect) {
                    _this.onConnect();
                }
            };
            /* Is run when the websocket connection is closed */
            this.OnClose = function (event) {
                console.log(_this.debugPrefix + ": websocket connection closed ", event);
            };
            /* Is run when the websocket connection has on error */
            this.OnError = function (event) {
                console.log(_this.debugPrefix + ": websocket connection error ", event);
            };
            /* Is run when the websocket connection has receieved a message */
            this.OnMessage = function (event) {
                var serverState = JSON.parse(event.data);
                console.log(_this.debugPrefix + ": websocket recieved message ", serverState);

                if (_this.onServerTick) {
                    _this.onServerTick(serverState);
                }
            };
            this.player = player;
            this.connection = new WebSocket(host);
            this.connection.onopen = this.OnOpen;
            this.connection.onclose = this.OnClose;
            this.connection.onerror = this.OnError;
            this.connection.onmessage = this.OnMessage;
            this.onServerTick = onServerTick;
            this.onConnect = onConnect;
            this.debugPrefix = debugPrefix;
        }
        return Connection;
    })();
    exports.Connection = Connection;
});
