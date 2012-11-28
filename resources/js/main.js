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

firstMessage = true;
socketMessaged = function(e) {
    jsObj = $.parseJSON(e.data);
    console.log("socket message:");
    console.log(jsObj);

    /* This is a message that describes start state */
    if (firstMessage === true) {
        firstMessage = false;
        console.log("Version is " + jsObj.Version);
        return;
    }

    RESPONSE_TICK_NUM = jsObj.LastTick;

    console.log("me is", parseInt(me));
    var myPlayer = jsObj.Players[parseInt(me)]
    var myEntityId = myPlayer.EntityId;

    $("#dynamic-controls").html("");
    $.each(myPlayer.AvailableActions, function(index, action) {
        addAction(action);
    });

    entLayer = jsObj.Boards[0].Layers[0];
    $.each(jsObj.Entities, function(index, player) {
        entLayer[player.Y] = entLayer[player.Y].substr(0, player.X) +
            'X' + entLayer[player.Y].substr(player.X + 1);
    });
    $("#board").html(entLayer.join("<br/>"));
    drawBoard(jsObj);

    if (jsObj.Entities[myEntityId].ActionQueue.length > 0) {
        increaseTick();
        tick(true);
    }
    setQueue(jsObj.Entities[myEntityId].ActionQueue)
}
/**** End WebSocket Hooks ****/

/* Adds a new ability to the client */
addAction = function(action) {
    var cont = $("#dynamic-controls");

    if (action.AvailableDirections[0]) {
        var t = $("<div>" + action.ShortDescription +
            " Left</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued left order: " + action.LongDescription);
            addToQueue(action.ActionString + "l");
        });
    }

    if (action.AvailableDirections[1]) {
        var t = $("<div>" + action.ShortDescription +
            " Right</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued right order: " + action.LongDescription);
            addToQueue(action.ActionString + "r");
        });
    }

    if (action.AvailableDirections[2]) {
        var t = $("<div>" + action.ShortDescription +
            " Up</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued up order: " + action.LongDescription);
            addToQueue(action.ActionString + "u");
        });
    }

    if (action.AvailableDirections[3]) {
        var t = $("<div>" + action.ShortDescription +
                " Down</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued down order: " + action.LongDescription);
            addToQueue(action.ActionString + "d");
        });
    }

    if (action.AvailableDirections[4]) {
        var t = $("<div>" + action.ShortDescription +
                " Self</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued self order: " + action.LongDescription);
            addToQueue(action.ActionString + "s");
        });
    }

    if (action.AvailableDirections[0] == false &&
        action.AvailableDirections[1] == false &&
        action.AvailableDirections[2] == false &&
        action.AvailableDirections[3] == false &&
        action.AvailableDirections[4] == false) {

        var t = $("<div>" + action.ShortDescription +
                "</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued order: " + action.LongDescription);
            addToQueue(action.ActionString);
        });
    }

    $("<br />").appendTo(cont);
};

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
            if (layer[i] === "#") {
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

        /* Entity icon */
        ctx.drawImage(entityImg, entity.X * tileSize,
            entity.Y * tileSize, tileSize - 2, tileSize - 2);

        /* Health Bar */
        ctx.fillStyle = "rgb(200,0,0)";
        ctx.fillRect(entity.X * tileSize,
                (entity.Y + 1) * tileSize - 5,
                tileSize - 1,
                4);
        var ardourPercent = entity.Ardour / entity.MaxArdour
        ctx.fillStyle = "rgb(0,200,0)";
        ctx.fillRect(entity.X * tileSize,
                (entity.Y + 1) * tileSize - 5,
                (tileSize - 1) * ardourPercent,
                4);
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
        console.log("Sent start game message");
        sendMessage(COMMANDS.startGame())
        increaseTick();
        tick();
        $("#start").toggleClass("disabled").off()
    });

    $("#ready").click(function() {
        console.log("Queue ready!");
        sendMessage(COMMANDS.queueActions(QUEUE, TICK_NUM));
        increaseTick();
        tick(true);
    })

    $(document).keydown(function(e){
        if (e.keyCode == 37) {
            $("div:contains('Move Left')").click();
            return false;
        }
        if (e.keyCode == 38) {
            $("div:contains('Move Up')").click();
            return false;
        }
        if (e.keyCode == 39) {
            $("div:contains('Move Right')").click();
            return false;
        }
        if (e.keyCode == 40) {
            $("div:contains('Move Down')").click();
            return false;
        }
        if (e.keyCode == 13) {
            $("#ready").click();
            return false;
        }
    });
})
