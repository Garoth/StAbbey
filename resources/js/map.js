goog.require("st.define");
goog.require("st.dom");
goog.require("st.Connection");

goog.provide("st.map");

(function() {
    var clientTick = 0;
    var mapCanvasNode = document.querySelector('#map');
    var entityCanvasNode = document.querySelector('#entities');
    /** Stores a map of image names to image objects */
    var images = {};

    /**
     * Makes the canvas 'fit best' into the body element's size.
     *
     * @param {Node} canvas
     */
    var fitCanvas = function(canvas) {
        var bodyWidth = document.body.clientWidth;
        var bodyHeight = document.body.clientHeight;
        var canvasWidth = parseInt(canvas.getAttribute('data-width'), 10);
        var canvasHeight = parseInt(canvas.getAttribute('data-height'), 10);
        var newWidth = 0;
        var newHeight = 0;

        if (bodyWidth / canvasWidth * canvasHeight > bodyHeight) {
            newHeight = bodyHeight;
            newWidth = canvasWidth / canvasHeight * bodyWidth;
        } else {
            newHeight = canvasHeight / canvasWidth * bodyHeight;
            newWidth = bodyWidth;
        }

        canvas.setAttribute('width', newWidth);
        canvas.setAttribute('height', newHeight);
    };

    /**
     * Renders the board from the server's data
     *
     * @param {Array.<string>} layer
     */
    var drawBoard = function(layer) {
        fitCanvas(mapCanvasNode);
        var canvasWidth = mapCanvasNode.getAttribute('width');
        var canvasHeight = mapCanvasNode.getAttribute('height');
        var tileSize = Math.round(canvasHeight / (layer.length + 0.5));
        var borderX = tileSize * 0.32;
        var borderY = tileSize * 0.25;
        var ctx = mapCanvasNode.getContext("2d");

        ctx.clearRect(0, 0, canvasWidth, canvasHeight);

        /* Draw outer border of walls */
        var wallImage = images['wall-cross.png'];
        for (var i = 0; i < layer[0].length + 2; i++) {
            ctx.drawImage(wallImage, i * tileSize - (tileSize - borderX),
                    - (tileSize - borderY), tileSize, tileSize);
            ctx.drawImage(wallImage, i * tileSize - (tileSize - borderX),
                    borderY + layer.length * tileSize, tileSize, tileSize);
        }

        for (i = 0; i < layer.length; i++) {
            ctx.drawImage(wallImage, - (tileSize - borderX),
                    i * tileSize + borderY, tileSize, tileSize);
            ctx.drawImage(wallImage, borderX + layer[0].length * tileSize,
                    i * tileSize + borderY, tileSize, tileSize);
        }

        /* Draws board */
        for (var x = 0; x < layer.length; x++) {
            var row = layer[x];
            for (i = 0; i < row.length; i++) {
                /* Draw floor everywhere */
                ctx.drawImage(images['ground.png'], i * tileSize + borderX,
                        x * tileSize + borderY, tileSize, tileSize);

                /* Draw water tiles */
                if (row[i] === "~") {
                    ctx.drawImage(images['water.png'], i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }

                /* Draw carpet tiles */
                if (row[i] === "c") {
                    ctx.drawImage(images['carpet.png'], i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }

                /* Draw grass tiles */
                if (row[i] === "g" || row[i] === "f") {
                    ctx.drawImage(images['grass.png'], i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }

                /* Draw flower overlays */
                if (row[i] === "f") {
                    ctx.drawImage(images['flowers.png'], i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }

                /* Draw walls */
                if (row[i] === "#") {
                    ctx.drawImage(images['wall-cross.png'], i * tileSize + borderX,
                            x * tileSize + borderY, tileSize, tileSize);
                }

                /* Draw tile outlines */
                ctx.strokeRect(i * tileSize + borderX, x * tileSize + borderY,
                        tileSize, tileSize);
            }
        }
    };

    /**
     * Renders the entities layer overtop of the board layer
     *
     * @param {st.define.WorldState} serverState
     */
    var drawEntities = function(serverState) {
        fitCanvas(entityCanvasNode);
        var canvasWidth = entityCanvasNode.getAttribute('width');
        var canvasHeight = entityCanvasNode.getAttribute('height');
        var layer = serverState.Boards[serverState.CurrentBoard].Layers[0];
        var tileSize = Math.round(canvasHeight / (layer.length + 0.5));
        var ctx = entityCanvasNode.getContext("2d");
        var borderX = tileSize * 0.32;
        var borderY = tileSize * 0.25;

        ctx.clearRect(0, 0, canvasWidth, canvasHeight);

        for (var index = 0; index < serverState.Entities.length; index++) {
            var entity = serverState.Entities[index];

            if (entity.Ardour === 0 ||
                entity.BoardId !== serverState.CurrentBoard) {

                continue;
            }

            var entityImg = null;
            var entityPlaceholderText = "";
            var drawHealth = true;

            if (entity.Type === "monster") {
                if (entity.Subtype === "gargoyle") {
                    // TODO:athorp:2014-10-24 should be a gargoyle
                    entityImg = images['monster-shade.png'];
                } else if (entity.Subtype === "chest") {
                    entityImg = images['monster-shade.png'];
                } else if (entity.Subtype === "chest") {
                    entityImg = images['chest.png'];
                } else if (entity.Subtype === "boulder") {
                    entityImg = images['boulder.png'];
                }

            } else if (entity.Type === "player") {
                if (entity.Name === "Player 0") {
                    entityImg = images['monk-green.png'];
                } else if (entity.Name === "Player 1") {
                    entityImg = images['monk-red.png'];
                }

            } else if (entity.Type === "trigger") {
                if (entity.Subtype === "ability loot") {
                    entityPlaceholderText = "Loot";
                } else if (entity.Subtype === "teleport trap") {
                    entityPlaceholderText = "Teleport";
                } else if (entity.Subtype === "caltrop trap") {
                    entityPlaceholderText = "Caltrop";
                } else if (entity.Subtype === "stairs up") {
                    entityImg = images['stairs-up.png'];
                } else if (entity.Subtype === "boulder trap") {
                    entityPlaceholderText = "Bldr T.";
                }

            } else if (entity.Type === "inert") {
                if (entity.Subtype === "sprung trap") {
                    entityPlaceholderText = "No Trap";
                    drawHealth = false;
                } else if (entity.Subtype === "tree") {
                    entityImg = images['tree.png'];
                } else if (entity.Subtype === "statue") {
                    entityImg = images['statue.png'];
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
        }
    };

    /**
     * Loads the desired images into JS objects and calls the given callback
     * when all of them are done.
     *
     * @param {string} imgPath path to image folder
     * @param {Array.<string>} sources list of image names (in imgPath)
     * @param {function(?): undefined} callback Function should expect a map
     *      from image paths to loaded JS images
     */
    var loadImages = function(imgPath, sources, callback) {
        var images = {};
        var loadedImages = 0;
        var numImages = sources.length;

        if (numImages === 0) {
            callback(null);
            return;
        }

        for (var i = 0; i < numImages; i++) {
            var src = sources[i];
            images[src] = new Image();

            images[src].onload = function() {
                if(++loadedImages >= numImages) {
                    callback(images); } };

            images[src].onerror = function() {
                console.log('CRITICAL ERROR: COULDNT LOAD IMAGE', src);
            };

            images[src].src = imgPath + src;
        }
    };

    /**
     * Causes the board to render based on server's data
     *
     * @param {(st.define.WorldState|st.define.Version)} serverState
     */
    var handleServerTick = function(serverState) {
        if (serverState.Boards && serverState.Entities) {
            var layer = serverState.Boards[serverState.CurrentBoard].Layers[0];
            drawBoard(layer);
            drawEntities(/** @type {st.define.WorldState} */ (serverState));
        }
    };

    // Connect to the server via WebSocket.
    var conn = null;
    console.log('Loading board images...');
    var imagesToLoad = [
        'fire.png', 'ground.png', 'hole.png',
        'monk-green.png', 'monk-red.png',
        'wall.png', 'wall-corner.png', 'wall-cross.png', 'wall-t.png',
        'monster-shade.png',
        'boulder.png', 'carpet.png', 'stairs-down.png', 'stairs-up.png',
        'wall-lamp.png', 'chest.png', 'statue.png', 'flowers.png',
        'tree.png', 'water.png', 'grass.png',
        'mia.png'
    ];
    loadImages('/resources/img/tile/', imagesToLoad, function(loadedImages) {
        console.log('All images loading, starting up game');
        images = loadedImages;
        conn = st.Connection("Map", STABBEY.HOST,
            null, handleServerTick, function() {
                conn.sendStartGame();
                conn.sendTick(++clientTick);
            });
    });
})();
