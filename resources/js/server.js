/* Server Library
 *
 * A library that abstracts away communication with the server. It takes
 * care of wrangling websockets and providing data in a convenient way.
 */

/* Data structure for formatting commands to send to the server */
var ServerCommands = function() {
    me = {};

    var START_GAME = 0;
    var UPDATE_TICK = 1;
    var SEND_QUEUE = 2;

    /* Tells the server that client is ready to start the game */
    me.startGame = function() {
        return { CommandCode: START_GAME };
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

/* Data structure for organizing your action queue to send to the server */
var MoveQueue = function() {
    var me = {};
    me.length = 0;

    var ActionQueue = [];

    me.set = function(queue) {
        ActionQueue = queue;
        me.length = ActionQueue.length;
    };

    me.get = function() {
        return ActionQueue;
    };

    me.add = function(element) {
        ActionQueue.push(element);
        me.length = ActionQueue.length;
    };

    me.clear = function() {
        ActionQueue = [];
        me.length = 0;
    };

    me.toString = function() {
        return ActionQueue.join(", ");
    };

    return me;
}();

/* Handles the interaction with the websocket / server API.  */
var Server = function() {
    var me = {};

    me.MyTickNumber = 0;
    me.ServerTickNum = 0;
    me.TickFrequency = 3000;
    me.serverTickHandler = null;
    me.clientTickHandler = null;

    /* Sets the function to call when the server's sent an update
     *
     * @param {function(Object newServerState)} serverTickHandler
     *      Sends newly acquired server state to this function
     */
    me.setServerTickHandler = function(handler) {
        me.serverTickHandler = handler;
    };

    /* Sets the function to call when the client's tick is updated
     *
     * @param {function(number newTick)} clientTickHandler
     *      Sends updated tick number to this function
     */
    me.setClientTickHandler = function(handler) {
        me.clientTickHandler = handler;
    };

    if (window.conn == null) {
        console.log("Err: connection wasn't set up for server library");
        return null;
    }

    window.conn.onopen = function(e) {
        console.log("Server: websocket connection opened\n", e);
    };

    window.conn.onclose = function(e) {
        console.log("Server: websocket connection closed\n", e);
    };

    window.conn.onerror = function(e) {
        console.log("Server: websocket connection error\n", e);
    };

    /* Send any message to the server */
    me.sendMessage = function(vars) {
        vars.Gamekey = gamekey;
        vars.Player = window.player;
        var msg = JSON.stringify(vars);
        console.log("Sending message: " + msg);
        window.conn.send(msg);
    };

    me.setTick = function(tickNum) {
        if (tickNum == null) {
            tickNum = me.TickNum + 1;
        }

        me.TickNum = tickNum;

        if (me.clientTickHandler != null) {
            me.clientTickHandler(me.TickNum);
        }
    };

    var firstMessage = true;
    window.conn.onmessage = function(e) {
        serverState = $.parseJSON(e.data);
        console.log("Server: websocket received message");
        console.log(serverState);

        /* This is a message that describes start state */
        if (firstMessage === true) {
            firstMessage = false;
            console.log("Server: version is " + serverState.Version);
            return;
        }

        me.ServerTickNum = serverState.LastTick;

        serverState.myPlayer = serverState.Players[parseInt(window.player)];
        var entityId = serverState.myPlayer.EntityId;
        serverState.myEntityId = entityId;

        MoveQueue.set(serverState.Entities[entityId].ActionQueue);

        me.serverTickHandler(serverState);

        if (serverState.Entities[entityId].ActionQueue.length > 0) {
            me.setTick(me.ServerTickNum + 1);
            me.tick();
        } else {
            me.setTick(me.ServerTickNum);
            me.tick();
        }
    };

    /* Sends the client's current tick to the server */
    me.tick = function() {
        me.sendMessage(ServerCommands.tick(me.TickNum));
    };

    /* Tells the server that this client is ready to join the game */
    me.startGame = function() {
        me.sendMessage(ServerCommands.startGame());
        me.setTick(me.ServerTickNum + 1);
        me.tick();
        /* Send the current tick in regular intervals as a heartbeat */
        setInterval(me.tick, me.TickFrequency);
    };

    /* Sends the current action queue to the server */
    me.sendQueue = function() {
        me.sendMessage(ServerCommands.queueActions(MoveQueue.get(), me.TickNum));
        me.setTick(me.ServerTickNum + 1);
        me.tick();
    };

    return me;
}();
