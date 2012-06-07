TICK_NUM = 0;
RESPONSE_TICK_NUM = 0;
TICK_DURATION = 3000;

/**** GAE Socket Hooks ****/
socketOpened = function() {
  console.log("socket opened");
}

socketClosed = function() {
  console.log("socket closed");
}

socketErrored = function() {
  console.log("socket errored!!!");
}

socketMessaged = function(msg) {
  jsObj = $.parseJSON(msg.data);
  console.log("socket message:");
  console.log(jsObj);
  RESPONSE_TICK_NUM = jsObj.LastTick;

  entLayer = jsObj.Boards[0].Layers[0];
  $.each(jsObj.Players, function(index, player) {
    console.log("Player with " + player.Y + " " + player.X)
    entLayer[player.Y] = entLayer[player.Y].substr(0, player.X) + 'X' + entLayer[player.Y].substr(player.X + 1)
  });
  $("#board").html(entLayer.join("<br/>"));
}
/**** End GAE Socket Hooks ****/

/* Send any message to the server */
sendMessage = function(path, urlvars) {
  urlvars.gamekey = gamekey;
  urlvars.player = me;
  var first = true
  for (var variable in urlvars) {
    if (first === true) {
      path += "?";
      first = false
    } else {
      path += "&";
    }
    path += variable + "=" + urlvars[variable];
  }

  var xhr = new XMLHttpRequest();
  console.log("Sending message: " + path);
  xhr.open("POST", path, true);
  xhr.send();
}

/* Send an order to the server */
sendCommand = function(urlvars) {
  sendMessage("/update", urlvars);
}

/* Various actions that could be sent to the server */
ACTIONS = function() {
  var me = {}

  /* List of commands the server understands. These will be validated by the
   * server to ensure that the client is actually able to do what they're
   * asking to do.
   *
   * The base codes below are used to create full codes.
   */
  var base_code_verbs = {
    move: "m"
  }

  me.DIRECTIONS = {
    LEFT  : "l",
    RIGHT : "r",
    UP    : "u",
    DOWN  : "d"
  }

  me.move = function(direction) {
    return base_code_verbs.move + direction
  }

  return me;
}();

/* Various command that can be sent to the server with sendCommand */
COMMANDS = function() {
  me = {};

  var START_GAME = 0;
  var UPDATE_TICK = 1;
  var SEND_QUEUE = 2;

  /* Tells the server that client is ready to start the game */
  me.startGame = function() {
    return { commandcode: START_GAME }
  };

  /* Tells the server the client's current tick */
  me.tick = function(tick) {
    return { commandcode: UPDATE_TICK, ticknum: tick };
  };

  /* Overwrites the player's actions
   * tick: the turn to overwrite actions for (should be current)
   * queue: array of command codes
   */
  me.queueActions = function(queue, tick) {
    return { commandcode: SEND_QUEUE, queue: queue.join("-"), ticknum: tick };
  };

  return me;
}();

/* Represents each client's tick loop */
tick = function() {
  if (TICK_NUM === RESPONSE_TICK_NUM) {
    TICK_NUM += 1;
  }

  sendCommand(COMMANDS.tick(TICK_NUM));
  setTimeout(tick, TICK_DURATION);
}

/* Stuff to run after the page has loaded */
$(function() {
  $("#start").click(function() {
    console.log("Trying to send message");
    sendCommand(COMMANDS.startGame())
    tick();
  });

  $("#move-right").click(function() {
    console.log("Sending order to move right");
    sendCommand(COMMANDS.queueActions(
        [ACTIONS.move(ACTIONS.DIRECTIONS.RIGHT)], TICK_NUM));
  });

  $("#move-left").click(function() {
    console.log("Sending order to move left");
    sendCommand(COMMANDS.queueActions(
        [ACTIONS.move(ACTIONS.DIRECTIONS.LEFT)], TICK_NUM));
  });

  $("#move-up").click(function() {
    console.log("Sending order to move up");
    sendCommand(COMMANDS.queueActions(
        [ACTIONS.move(ACTIONS.DIRECTIONS.UP)], TICK_NUM));
  });

  $("#move-down").click(function() {
    console.log("Sending order to move down");
    sendCommand(COMMANDS.queueActions(
        [ACTIONS.move(ACTIONS.DIRECTIONS.DOWN)], TICK_NUM));
  });
})
