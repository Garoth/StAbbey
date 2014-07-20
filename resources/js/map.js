goog.require("st.define");
goog.require("st.dom");
goog.require("st.Connection");

goog.provide("st.map");

goog.scope(function() {
    var ENTITY_NAME_TO_IMAGE = {
        'monster': 'monster-shade',
        'player': 'monk',
    };
    var TILE_SIZE = 512.0;
    var TILE_TYPE_TO_IMAGE = {
        '#': 'wall',
        '~': 'fire'
    };
    var TILE_SRC_PREFIX = '/resources/img/tile/';
    // var WALL_TYPE = '#';

    var clientTick = 0;
    // var debugEntities = st.getById('debugEntities');
    var entitiesDiv = st.dom.getById('entities');
    var mapDiv = st.dom.getById('map');
    var mapWidth, mapHeight;

    window.onresize = function() {
        if (mapWidth && mapHeight) {
            document.body.style.zoom = Math.min(
                    window.innerWidth / mapWidth / TILE_SIZE,
                    window.innerHeight / mapHeight / TILE_SIZE);
        }
    };

    window.onkeydown = function(event) {
        if (event.keyCode === 69 /* E */) {
            st.dom.toggleClass(document.body, 'debugEntities');
        }
    };

    var isWall = function(type) {
        return type === '#' || type === '|' || type === null;
    };

    var createEntity = function(entity) {
        var imageName = ENTITY_NAME_TO_IMAGE[entity.Type];
        var container;
        if (imageName) {
            container = st.dom.createImg(null, null,
                    TILE_SRC_PREFIX + imageName + '.png');
        } else {
            container = st.dom.createDiv(null, 'debugEntity');
            container.innerHTML =
                entity.Name + '<br>' + entity.Type + '<br>' + entity.Subtype;
        }
        if (container) {
            container.style.position = 'absolute';
            container.style.left = TILE_SIZE * entity.X + 'px';
            container.style.top = TILE_SIZE * entity.Y + 'px';
            return container;
        }
    };

    var createTile = function(type, upType, rightType, downType, leftType) {
        var img = st.dom.createImg();
        var name = TILE_TYPE_TO_IMAGE[type];
        var rotate = '';

        if (type === '#') {
            // Wall tiles change depending on neighbors.
            var wallCount = 0;
            if (isWall(upType)) {
                wallCount++;
            }

            if (isWall(rightType)) {
                wallCount++;
            }

            if (isWall(downType)) {
                wallCount++;
            }

            if (isWall(leftType)) {
                wallCount++;
            }

            if (wallCount === 4) {
                name = 'wall-cross';
            } else if (wallCount === 3) {
                name = 'wall-t';
                if (!isWall(leftType)) {
                    rotate = -90;
                }

                if (!isWall(rightType)) {
                    rotate = 90;
                }

                if (!isWall(downType)) {
                    rotate = 180;
                }
            } else if (wallCount === 2) {
                if (isWall(leftType)) {
                    if (isWall(upType)) {
                        name = 'wall-corner';
                        rotate = -90;
                    }
                    if (isWall(downType)) {
                        name = 'wall-corner';
                    }
                }
                if (isWall(upType)) {
                    if (isWall(rightType)) {
                        name = 'wall-corner';
                        rotate = 180;
                    }
                    if (isWall(downType)) {
                        rotate = 90;
                    }
                }
                if (isWall(rightType) && isWall(downType)) {
                    name = 'wall-corner';
                    rotate = -90;
                }
            }
        }

        img.src = TILE_SRC_PREFIX + (name || 'ground') + '.png';

        if (rotate) {
            img.style.webkitTransform = 'rotate(' + rotate + 'deg)';
        }

        return img;
    };

    var buildBoard = function(board) {
        mapHeight = board.length;
        mapWidth = board[0].length;
        mapDiv.innerHTML = '';
        for (var y = 0; y < mapHeight; y++) {
            var rowDiv = st.dom.createDiv();
            for (var x = 0; x < mapWidth; x++) {
                var upType = y > 0 ? board[y - 1][x] : null;
                var rightType = x < mapWidth - 1 ? board[y][x + 1] : null;
                var downType = y < mapHeight - 1 ? board[y + 1][x] : null;
                var leftType = x > 0 ? board[y][x - 1] : null;
                rowDiv.appendChild(createTile(
                            board[y][x], upType, rightType, downType, leftType));
            }
            mapDiv.appendChild(rowDiv);
        }
        window.onresize();
    };

    var buildEntities = function(entities) {
        entitiesDiv.innerHTML = '';
        for (var i = 0; i < entities.length; i++) {
            var entity = createEntity(entities[i]);
            if (entity) {
                entitiesDiv.appendChild(entity);
            }
        }
    };

    var handleServerTick = function(serverState) {
        if (serverState.Boards) {
            buildBoard(serverState.Boards[serverState.CurrentBoard].Layers[0]);
        }

        if (serverState.Entities) {
            buildEntities(serverState.Entities);
        }
    };

    // Connect to the server via WebSocket.
    var conn = new st.Connection("Map", STABBEY.HOST,
            null, handleServerTick, function() {
                conn.sendStartGame();
                conn.sendTick(++clientTick);
            });
});
