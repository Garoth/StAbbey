var st = {};

// Flags used to indicate available card actions.
st.CardActions = {
  LEFT: 1,
  DOWN: 2,
  UP: 4,
  RIGHT: 8,
  SELF: 16,
  CHANNEL: 32,
  ALL: 63,
  DIRS: 15
};

// Shorthands.
st.getById = function(id) {
  return document.getElementById(id);
};
st.createElement = function(type, id, classname) {
  var element = document.createElement(type);
  if (id) {
    element.id = id;
  }
  if (classname) {
    element.className = classname;
  }
  return element;
};
st.createDiv = function(id, classname) {
  return st.createElement('div', id, classname);
};
st.createImg = function(id, classname, src) {
  var img = st.createElement('img', id, classname);
  img.src = src;
  return img;
};

// Removes all of the child elements of an element.
st.removeChildren = function(element) {
  while (element.firstChild) {
    element.removeChild(element.firstChild);
  }
};

// Removes an element from the DOM.
st.removeElement = function(element) {
  if (element.parentNode) {
    element.parentNode.removeChild(element);
  }
};

// Class to encapsulate a WebSocket connection to the server.
st.Connection = function(path, player, gamekey, onClientTick, onServerTick) {
  this.connection = new WebSocket(path);
  this.connection.onopen = this.onOpen.bind(this);
  this.connection.onclose = this.onClose.bind(this);
  this.connection.onerror = this.onError.bind(this);
  this.connection.onmessage = this.onMessage.bind(this);
  this.gamekey = gamekey;
  this.player = player;
  this.onClientTick = onClientTick;
  this.onServerTick = onServerTick;
  this.serverTick;
};
st.Connection.prototype.send = function(object) {
  object.Gamekey = this.gamekey;
  object.Player = this.player;
  var msg = JSON.stringify(object);
  console.log("Sending message: " + msg);
  this.connection.send(msg);
};

// Tells the server that client is ready to start the game.
st.Connection.prototype.sendStartGame = function() {
  this.send({CommandCode: 0});
};

// Tells the server the client's current tick.
st.Connection.prototype.sendTick = function(tick) {
  this.send({CommandCode: 1, TickNum: tick});
};

st.Connection.prototype.onOpen = function(event) {
  console.log("Server: websocket connection opened\n", event);
};
st.Connection.prototype.onClose = function(event) {
  console.log("Server: websocket connection closed\n", event);
};
st.Connection.prototype.onError = function(event) {
  console.log("Server: websocket connection error\n", event);
};
st.Connection.prototype.onMessage = function(event) {
  serverState = JSON.parse(event.data);
  console.log("Server: websocket recieved message");
  console.log(serverState);
  this.onServerTick && this.onServerTick(serverState);
};
