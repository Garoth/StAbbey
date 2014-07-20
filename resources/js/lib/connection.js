goog.provide('st.Connection');

goog.scope(function() {
    /* Encapsulates the connection to the server websocket */
    st.Connection = function(debugPrefix, host, player, onServerTick, onConnect) {

        var me = {};
        me.player = player;
        me.connection = new WebSocket(host);

        /* Sends an object as JSON via the websocket */
        me.send = function(object) {
            object.Player = me.player;
            var msg = JSON.stringify(object);
            console.log(debugPrefix + ": sending message " + msg);
            me.connection.send(msg);
        };

        /* Sends the server the message that this client is ready to start */
        me.sendStartGame = function() {
            me.send({CommandCode: 0});
        };

        /* Sends the server the given tick number */
        me.sendTick = function(tick) {
            me.send({CommandCode: 1, TickNum: tick});
        };

        /* Is run when the websocket connection is opened */
        me.onOpen = function(event) {
            console.log(debugPrefix + ": websocket opened ", event);

            if (onConnect) {
                onConnect();
            }
        };

        /* Is run when the websocket connection is closed */
        me.onClose = function(event) {
            console.log(debugPrefix + ": websocket closed ", event);
        };

        /* Is run when the websocket connection has on error */
        me.onError = function(event) {
            console.log(debugPrefix + ": websocket error ", event);
        };

        /* Is run when the websocket connection has receieved a message */
        me.onMessage = function(event) {
            // TODO define serverState more precisely
            var serverState = JSON.parse(event.data);
            console.log(debugPrefix + ": websocket message ", serverState);

            if (onServerTick) {
                onServerTick(serverState);
            }
        };

        me.connection.onopen = me.onOpen;
        me.connection.onclose = me.onClose;
        me.connection.onerror = me.onError;
        me.connection.onmessage = me.onMessage;

        return me;
    };
});
