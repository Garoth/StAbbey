/* Handles updates to the server state */
window.serverStateHandler = function(serverState) {
    /* Set up actions that are currently available */
    $("#dynamic-controls").html("");
    for (var i = 0; i < serverState.myPlayer.AvailableActions.length; i++) {
        window.addAction(serverState.myPlayer.AvailableActions[i]);
    }

    /* Redraw the board */
    window.drawBoard(serverState);

    /* Update the player's current action queue */
    $("#queue").text(MoveQueue.toString());
}
Server.setServerTickHandler(window.serverStateHandler);

/* Handles updates to the client's tick number */
window.clientTickUpdateHandler = function(newTick) {
    $("#tick").text("" + newTick);
}
Server.setClientTickHandler(window.clientTickUpdateHandler);

/* Adds a new ability to the client */
window.addAction = function(action) {
    var cont = $("#dynamic-controls");
    var t = null;

    if (action.AvailableDirections[0]) {
        t = $("<div>" + action.ShortDescription +
            " Left</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued left order: " + action.LongDescription);
            MoveQueue.add(action.ActionString + "l");
            $("#queue").text(MoveQueue.toString());
        });
    }

    if (action.AvailableDirections[1]) {
        t = $("<div>" + action.ShortDescription +
            " Right</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued right order: " + action.LongDescription);
            MoveQueue.add(action.ActionString + "r");
            $("#queue").text(MoveQueue.toString());
        });
    }

    if (action.AvailableDirections[2]) {
        t = $("<div>" + action.ShortDescription +
            " Up</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued up order: " + action.LongDescription);
            MoveQueue.add(action.ActionString + "u");
            $("#queue").text(MoveQueue.toString());
        });
    }

    if (action.AvailableDirections[3]) {
        t = $("<div>" + action.ShortDescription +
                " Down</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued down order: " + action.LongDescription);
            MoveQueue.add(action.ActionString + "d");
            $("#queue").text(MoveQueue.toString());
        });
    }

    if (action.AvailableDirections[4]) {
        t = $("<div>" + action.ShortDescription +
                " Self</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued self order: " + action.LongDescription);
            MoveQueue.add(action.ActionString + "s");
            $("#queue").text(MoveQueue.toString());
        });
    }

    if (action.AvailableDirections[0] === false &&
        action.AvailableDirections[1] === false &&
        action.AvailableDirections[2] === false &&
        action.AvailableDirections[3] === false &&
        action.AvailableDirections[4] === false) {

        t = $("<div>" + action.ShortDescription +
                "</div>").appendTo(cont);
        t.click(function() {
            console.log("Queued order: " + action.LongDescription);
            MoveQueue.add(action.ActionString);
            $("#queue").text(MoveQueue.toString());
        });
    }

    $("<br />").appendTo(cont);
};

window.tileImages = {};
window.IMAGES_LOADED = false;
/* Loads the tiles for the map so that they are ready to be drawn */
window.loadTiles = function() {
    var imgPath = "/resources/img/";

    tileImages.floor = new Image();
    var floorDef = $.Deferred();
    tileImages.floor.onload = function() {
        floorDef.resolve();
    };
    tileImages.floor.src = imgPath + "Floor.png";

    tileImages.grass = new Image();
    var grassDef = $.Deferred();
    tileImages.grass.onload = function() {
        grassDef.resolve();
    };
    tileImages.grass.src = imgPath + "Grass.png";

    tileImages.flowers = new Image();
    var flowersDef = $.Deferred();
    tileImages.flowers.onload = function() {
        flowersDef.resolve();
    };
    tileImages.flowers.src = imgPath + "Flowers.png";

    tileImages.water = new Image();
    var waterDef = $.Deferred();
    tileImages.water.onload = function() {
        waterDef.resolve();
    };
    tileImages.water.src = imgPath + "Water.png";

    tileImages.wall = new Image();
    var wallDef = $.Deferred();
    tileImages.wall.onload = function() {
        wallDef.resolve();
    };
    tileImages.wall.src = imgPath + "Wall.png";

    tileImages.bluePlayer = new Image();
    var bluePlayerDef = $.Deferred();
    tileImages.bluePlayer.onload = function() {
        bluePlayerDef.resolve();
    };
    tileImages.bluePlayer.src = imgPath + "Player-Blue.png";

    tileImages.greenPlayer = new Image();
    var greenPlayerDef = $.Deferred();
    tileImages.greenPlayer.onload = function() {
        greenPlayerDef.resolve();
    };
    tileImages.greenPlayer.src = imgPath + "Player-Green.png";

    tileImages.genericMonster = new Image();
    var genericMonsterDef = $.Deferred();
    tileImages.genericMonster.onload = function() {
        genericMonsterDef.resolve();
    };
    tileImages.genericMonster.src = imgPath + "Generic-Monster.png";

    tileImages.chest = new Image();
    var chestDef = $.Deferred();
    tileImages.chest.onload = function() {
        chestDef.resolve();
    };
    tileImages.chest.src = imgPath + "Chest.png";

    tileImages.tree = new Image();
    var treeDef = $.Deferred();
    tileImages.tree.onload = function() {
        treeDef.resolve();
    };
    tileImages.tree.src = imgPath + "Tree.png";

    tileImages.carpet = new Image();
    var carpetDef = $.Deferred();
    tileImages.carpet.onload = function() {
        carpetDef.resolve();
    };
    tileImages.carpet.src = imgPath + "RedCarpet.png";

    tileImages.statue = new Image();
    var statueDef = $.Deferred();
    tileImages.statue.onload = function() {
        statueDef.resolve();
    };
    tileImages.statue.src = imgPath + "Statue.png";

    tileImages.stairsup = new Image();
    var stairsupDef = $.Deferred();
    tileImages.stairsup.onload = function() {
        stairsupDef.resolve();
    };
    tileImages.stairsup.src = imgPath + "Stairs-Up.png";

    tileImages.stairsdown = new Image();
    var stairsdownDef = $.Deferred();
    tileImages.stairsdown.onload = function() {
        stairsdownDef.resolve();
    };
    tileImages.stairsdown.src = imgPath + "Stairs-Down.png";

    tileImages.boulder = new Image();
    var boulderDef = $.Deferred();
    tileImages.boulder.onload = function() {
        boulderDef.resolve();
    };
    tileImages.boulder.src = imgPath + "Boulder.png";

    $.when(floorDef, chestDef, wallDef, bluePlayerDef, treeDef, waterDef,
            carpetDef, greenPlayerDef, genericMonsterDef, grassDef,
            statueDef, flowersDef, stairsupDef, stairsdownDef,
            boulderDef).then(function() {
        IMAGES_LOADED = true;
        console.log("ALL IMAGES LOADED!");
    });
};
window.loadTiles();

/* Draws the Canvas based map */
window.drawBoard = function(serverState) {
    var layers = serverState.Boards[serverState.CurrentBoard].Layers[0];
    var canvas = document.getElementById("canvas-board");
    var canvasWidth = $("#canvas-board").width();
    var canvasHeight = $("#canvas-board").height();
    var tileSize = Math.round(canvasHeight / (layers.length + 0.5));
    var borderX = tileSize * 0.32;
    var borderY = tileSize * 0.25;
    var ctx = canvas.getContext("2d");

    ctx.clearRect(0, 0, canvasWidth, canvasHeight);

    /* Draw outer border of walls */
    for (var i = 0; i < layers[0].length + 2; i++) {
        ctx.drawImage(tileImages.wall, i * tileSize - (tileSize - borderX),
                - (tileSize - borderY), tileSize, tileSize);
        ctx.drawImage(tileImages.wall, i * tileSize - (tileSize - borderX),
                borderY + layers.length * tileSize, tileSize, tileSize);
    }

    for (i = 0; i < layers.length; i++) {
        ctx.drawImage(tileImages.wall, - (tileSize - borderX),
                i * tileSize + borderY, tileSize, tileSize);
        ctx.drawImage(tileImages.wall, borderX + layers[0].length * tileSize,
                i * tileSize + borderY, tileSize, tileSize);
    }

    /* Draws main board under everything */
    for (var x = 0; x < layers.length; x++) {
        var layer = layers[x];
        for (i = 0; i < layer.length; i++) {
            /* Draw floor everywhere */
            if (IMAGES_LOADED) {
                ctx.drawImage(tileImages.floor, i * tileSize + borderX,
                        x * tileSize + borderY, tileSize, tileSize);
            }

            /* Draw water tiles */
            if (layer[i] === "~") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.water, i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }
            }

            /* Draw carpet tiles */
            if (layer[i] === "c") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.carpet, i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }
            }

            /* Draw grass tiles */
            if (layer[i] === "g" || layer[i] === "f") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.grass, i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }
            }

            /* Draw flower overlays */
            if (layer[i] === "f") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.flowers, i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }
            }

            /* Draw walls */
            if (layer[i] === "#") {
                if (IMAGES_LOADED) {
                    ctx.drawImage(tileImages.wall, i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }
            }

            /* Draw tile outlines */
            ctx.strokeRect(i * tileSize + borderX, x * tileSize + borderY,
                    tileSize, tileSize);
        }
    }

    $.each(serverState.Entities, function(index, entity) {
        if (IMAGES_LOADED === false ||
                entity.Ardour === 0 ||
                entity.BoardId !== serverState.CurrentBoard) {
            return;
        }

        var entityImg = null;
        var entityPlaceholderText = "";
        var drawHealth = true;

        if (entity.Type === "monster") {
            if (entity.Subtype === "gargoyle") {
                entityImg = tileImages.genericMonster;
            } else if (entity.Subtype === "chest") {
                entityImg = tileImages.chest;
            } else if (entity.Subtype === "boulder") {
                entityImg = tileImages.boulder;
            }

        } else if (entity.Type === "player") {
            if (entity.Name === "Player 0") {
                entityImg = tileImages.bluePlayer;
            } else if (entity.Name === "Player 1") {
                entityImg = tileImages.greenPlayer;
            }

        } else if (entity.Type === "trigger") {
            if (entity.Subtype === "ability loot") {
                entityPlaceholderText = "Loot";
            } else if (entity.Subtype === "teleport trap") {
                entityPlaceholderText = "Teleport";
            } else if (entity.Subtype === "caltrop trap") {
                entityPlaceholderText = "Caltrop";
            } else if (entity.Subtype === "stairs up") {
                entityImg = tileImages.stairsup;
            } else if (entity.Subtype === "boulder trap") {
                entityPlaceholderText = "Bldr T.";
            }

        } else if (entity.Type === "inert") {
            if (entity.Subtype === "sprung trap") {
                entityPlaceholderText = "No Trap";
                drawHealth = false;
            } else if (entity.Subtype === "tree") {
                entityImg = tileImages.tree;
            } else if (entity.Subtype === "inert statue") {
                entityImg = tileImages.statue;
            }
        }

        /* Entity icon */
        if (entityImg != null) {
            ctx.drawImage(entityImg, entity.X * tileSize + borderX,
                entity.Y * tileSize + borderY, tileSize - 1, tileSize - 1);
        } else {
            ctx.fillStyle="#FFFFFF";
            ctx.font="10px Arial";
            ctx.fillText(entityPlaceholderText,
                    entity.X * tileSize + borderX + 4,
                    (entity.Y + 1) * tileSize + borderY - 10);
        }

        /* Health Bar */
        if (drawHealth === true && entity.Ardour !== entity.MaxArdour) {
            ctx.fillStyle = "rgb(200,0,0)";
            ctx.fillRect(entity.X * tileSize + borderX,
                    (entity.Y + 1) * tileSize - 5 + borderY,
                    tileSize - 1,
                    4);
            var ardourPercent = entity.Ardour / entity.MaxArdour;
            ctx.fillStyle = "rgb(0,200,0)";
            ctx.fillRect(entity.X * tileSize + borderX,
                    (entity.Y + 1) * tileSize - 5 + borderY,
                    (tileSize - 1) * ardourPercent,
                    4);
        }
    });

};

/* Stuff to run after the page has loaded */
$(function() {
    $("#start").click(function() {
        Server.startGame();
        $("#start").toggleClass("disabled").off();
    });

    $("#ready").click(function() {
        if (MoveQueue.length > 0) {
            Server.sendQueue();
        } else {
            console.log("Ignoring ready click with no queue");
        }
    });

    $(document).keydown(function(e){
        if (e.keyCode == 37) {
            $("div:contains('Move Left')").click();
            return;
        }
        if (e.keyCode == 38) {
            $("div:contains('Move Up')").click();
            return;
        }
        if (e.keyCode == 39) {
            $("div:contains('Move Right')").click();
            return;
        }
        if (e.keyCode == 40) {
            $("div:contains('Move Down')").click();
            return;
        }
        if (e.keyCode == 13) {
            $("#ready").click();
            return;
        }
    });
});
