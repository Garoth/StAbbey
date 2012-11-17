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
    drawBoard(jsObj);

    /* Probably could clobber what the player is trying to do. Beta-code */
    /* TODO EntityId might not be PlayerId sometime. Player should prob.
     *      report its EntityId as well as its Id */
    console.log("me is", parseInt(me));
    var myEntityId = jsObj.Players[parseInt(me)].EntityId;
    if (jsObj.Entities[myEntityId].ActionQueue.length > 0) {
        increaseTick();
        tick(true);
    }
    setQueue(jsObj.Entities[myEntityId].ActionQueue)
}
/**** End WebSocket Hooks ****/

tileImages = {}
IMAGES_LOADED = false;
/* Loads the tiles for the map so that they are ready to be drawn */
loadTiles = function() {
    var imgPath = "/resources/img/"

    tileImages.floor = new Image();
    var floorDef = $.Deferred();
    tileImages.floor.onload = function() {
        console.log("Floor image loaded");
        floorDef.resolve();
    };
    tileImages.floor.src = imgPath + "Floor.png";

    tileImages.column = new Image();
    var columnDef = $.Deferred();
    tileImages.column.onload = function() {
        console.log("Column image loaded");
        columnDef.resolve();
    };
    tileImages.column.src = imgPath + "Column.png";

    tileImages.bluePlayer = new Image();
    var bluePlayerDef = $.Deferred();
    tileImages.bluePlayer.onload = function() {
        console.log("Blue Player image loaded");
        bluePlayerDef.resolve();
    };
    tileImages.bluePlayer.src = imgPath + "Player-blue.png";

    tileImages.greenPlayer = new Image();
    var greenPlayerDef = $.Deferred();
    tileImages.greenPlayer.onload = function() {
        console.log("Green Player image loaded");
        greenPlayerDef.resolve();
    };
    tileImages.greenPlayer.src = imgPath + "Player-green.png";

    tileImages.genericMonster = new Image();
    var genericMonsterDef = $.Deferred();
    tileImages.genericMonster.onload = function() {
        console.log("Generic Monster image loaded");
        genericMonsterDef.resolve();
    };
    tileImages.genericMonster.src = imgPath + "Generic-Monster.png";

    $.when(floorDef, columnDef, bluePlayerDef,
            greenPlayerDef, genericMonsterDef).then(function() {
        IMAGES_LOADED = true;
        console.log("ALL IMAGES LOADED!");
    });
};
loadTiles();

/* Draws the Canvas based map */
drawBoard = function(serverState) {
    var layers = serverState.Boards[0].Layers[0];
    var canvas = document.getElementById("canvas-board");
    var canvasWidth = $("#canvas-board").width();
    var canvasHeight = $("#canvas-board").height();
    var tileSize = Math.round(canvasHeight / layers.length);
    var ctx = canvas.getContext("2d");

    ctx.clearRect(0, 0, canvasWidth, canvasHeight);

    for (var x = 0; x < layers.length; x++) {
        var layer = layers[x];
        for (var i = 0; i < layer.length; i++) {
            /* Draw tile outlines */
            ctx.strokeRect(i * tileSize, x * tileSize, tileSize, tileSize);

            /* Draw floor everywhere */
            if (IMAGES_LOADED) {
                ctx.drawImage(tileImages.floor, i * tileSize, x * tileSize,
                        tileSize, tileSize);
            }

            /* Draw walls */
            if (layer[i] === "-" || layer[i] === "L" || layer[i] === "|") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.column, i * tileSize,
                            x * tileSize, tileSize, tileSize);
                }
            }
        }
    }

    $.each(serverState.Entities, function(index, entity) {
        if (IMAGES_LOADED === false) {
            return;
        }

        var entityImg = null;

        if (entity.Type === "monster") {
            entityImg = tileImages.genericMonster;
        } else if (entity.Type === "player") {
            if (entity.Name === "Player 0") {
                entityImg = tileImages.bluePlayer;
            } else if (entity.Name === "Player 1") {
                entityImg = tileImages.greenPlayer;
            }
        }

        ctx.drawImage(entityImg, entity.X * tileSize,
            entity.Y * tileSize, tileSize - 2, tileSize - 2);
    });

};

/* Send any message to the server */
sendMessage = function(vars) {
    vars.Gamekey = gamekey;
    vars.Player = me;
    msg = JSON.stringify(vars);
    console.log("Sending message: " + msg);
    conn.send(msg);
};

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
    $("#tick").text("" + TICK_NUM);
}

/* Represents each client's tick loop */
tick = function(once) {
    if (once == null) {
        once = false
    }

    sendMessage(COMMANDS.tick(TICK_NUM));

    if (once != true) {
      setTimeout(tick, TICK_DURATION);
    }
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
        tick(true);
    })

    $(document).keydown(function(e){
        if (e.keyCode == 37) {
            $("#move-left").click();
            return false;
        }
        if (e.keyCode == 38) {
            $("#move-up").click();
            return false;
        }
        if (e.keyCode == 39) {
            $("#move-right").click();
            return false;
        }
        if (e.keyCode == 40) {
            $("#move-down").click();
            return false;
        }
        if (e.keyCode == 13) {
            $("#ready").click();
            return false;
        }
    });
})
