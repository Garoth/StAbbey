TICK_NUM = 0;
RESPONSE_TICK_NUM = 0;
QUEUE = [];
TICK_DURATION = 3000;

/**** WebSocket Hooks ****/
socketOpened = function(e) {
    console.log("socket opened");
    console.log(e);
}

socketClosed = function(e) {
    console.log("socket closed");
    console.log(e);
}

socketErrored = function(e) {
    console.log("socket errored!!!");
    console.log(e);
}

socketMessaged = function(e) {
    jsObj = $.parseJSON(e.data);
    console.log("socket message:");
    console.log(jsObj);
    RESPONSE_TICK_NUM = jsObj.LastTick;

    entLayer = jsObj.Boards[0].Layers[0];
    $.each(jsObj.Entities, function(index, player) {
        entLayer[player.Y] = entLayer[player.Y].substr(0, player.X) +
            'X' + entLayer[player.Y].substr(player.X + 1);
    });
    $("#board").html(entLayer.join("<br/>"));

    /* Probably could clobber what the player is trying to do. Beta-code */
    console.log("me is", parseInt(me))
    if (jsObj.Players[parseInt(me)].ActionQueue.length > 0) {
        increaseTick()
    }
    setQueue(jsObj.Players[parseInt(me)].ActionQueue)
}
/**** End WebSocket Hooks ****/

/* Send any message to the server */
sendMessage = function(urlvars) {
    urlvars.Gamekey = gamekey;
    urlvars.Player = me;
    msg = JSON.stringify(urlvars);
    console.log("Sending message: " + msg);
    conn.send(msg);
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

/* Various command that can be sent to the server with sendMessage */
COMMANDS = function() {
    me = {};

    var START_GAME = 0;
    var UPDATE_TICK = 1;
    var SEND_QUEUE = 2;

    /* Tells the server that client is ready to start the game */
    me.startGame = function() {
        return { CommandCode: START_GAME }
    };

    /* Tells the server the client's current tick */
    me.tick = function(tick) {
        return { CommandCode: UPDATE_TICK, TickNum: tick };
    };

    /* Overwrites the player's actions
     * tick: the turn to overwrite actions for (should be current)
     * queue: array of command codes
     */
    me.queueActions = function(queue, tick) {
        return { CommandCode: SEND_QUEUE, Queue: queue, TickNum: tick };
    };

    return me;
}();

setQueue = function(queue) {
      QUEUE = queue;
      $("#queue").text(QUEUE.join(", "));
}

addToQueue = function(element) {
    QUEUE.push(element);
    $("#queue").text(QUEUE.join(", "));
}

clearQueue = function() {
    QUEUE = [];
    $("#queue").text(QUEUE.join(", "));
}

increaseTick = function() {
    TICK_NUM += 1;
    $("#turn").text("" + TICK_NUM);
}

/* Represents each client's tick loop */
tick = function() {
    sendMessage(COMMANDS.tick(TICK_NUM));
    setTimeout(tick, TICK_DURATION);
}

/* Stuff to run after the page has loaded */
$(function() {
    $("#start").click(function() {
        console.log("Sending start game message");
        sendMessage(COMMANDS.startGame())
        increaseTick();
        tick();
        $("#start").toggleClass("disabled").off()
    });

   $("#move-right").click(function() {
        console.log("Sending order to move right");
        addToQueue(ACTIONS.move(ACTIONS.DIRECTIONS.RIGHT))
    });

    $("#move-left").click(function() {
        console.log("Sending order to move left");
        addToQueue(ACTIONS.move(ACTIONS.DIRECTIONS.LEFT))
    });

    $("#move-up").click(function() {
        console.log("Sending order to move up");
        addToQueue(ACTIONS.move(ACTIONS.DIRECTIONS.UP))
    });

    $("#move-down").click(function() {
        console.log("Sending order to move down");
        addToQueue(ACTIONS.move(ACTIONS.DIRECTIONS.DOWN))
    });

    $("#ready").click(function() {
        console.log("Queue ready!");
        sendMessage(COMMANDS.queueActions(QUEUE, TICK_NUM));
        increaseTick();
    })
})
