goog.require("st.define");
goog.require("st.dom");
goog.require("st.Connection");

goog.provide("st.hand");

(function() {
    var BUTTON_PREFIX = '/resources/img/card/btn_';
    var hand = [];
    var clientTick = 0;
    var containerDiv = st.dom.getById('cards');
    var queueNode = document.querySelector('#queue-content');
    var undoNode = document.querySelector('#undo');
    var readyNode = document.querySelector('#ready');

    var buildHand = function() {
        st.dom.removeChildren(containerDiv);
        for (var i = 0; i < hand.length; i++) {
            containerDiv.appendChild(hand[i].root);
        }
        arrangeHand();
    };

    var arrangeHand = function() {
        for (var i = 0; i < hand.length; i++) {
            var style = hand[i].root.style;
            style.left = (4 * i) + 'px';
            style.top = (90 * i) + 'px';
            style.zIndex = i;
        }
    };

    /**
     * Object wrapper for each of the buttons in each card.
     *
     * @param {string} dir
     * @param {function(string)|null} onClick
     */
    st.hand.CardButton = function(dir, onClick) {
        var me = {};
        me.buttonName = BUTTON_PREFIX + dir;
        me.root = null;
        me.root = st.dom.createImg('', 'button', me.buttonName + '.png');
        var pressed;

        if (!onClick) {
            me.root = st.dom.createImg('', 'button', me.buttonName + '_disable.png');
            return me;
        }

        me.root = st.dom.createImg('', 'button', me.buttonName + '.png');

        me.root.onmousedown = function() {
            me.root.src = me.buttonName + '_press.png';
            pressed = true;
        };

        me.root.onmouseout = function() {
            me.root.src = me.buttonName + '.png';
            pressed = false;
        };

        me.root.onmouseup = function() {
            if (pressed) {
                pressed = false;
                me.root.onmouseout();
                onClick(dir);
            }
        };

        return me;
    };

    /**
     * Wrapper object for a card.
     *
     * @param {st.define.Action} action
     */
    st.hand.Card = function(action) {
        var me = {};
        var name = action.ShortDescription;
        var count = '&infin;';

        me.root = st.dom.createDiv('', 'card-wrapper');
        var container = st.dom.createDiv('', 'card');
        // TODO:athorp:2014-10-20 card art should be dynmaic when we have it
        var artImg = st.dom.createImg('', 'art', '/resources/img/card/card_art_001.png');
        var maskImg = st.dom.createImg('', 'mask', '/resources/img/card/card.png');
        var cardTitle = st.dom.createDiv('', 'title');
        var cardCount = st.dom.createDiv('', 'count');
        var buttons = [];

        me.root.appendChild(container);
        cardTitle.innerHTML = name;
        cardCount.innerHTML = count;
        container.appendChild(artImg);
        container.appendChild(maskImg);
        container.appendChild(cardTitle);
        container.appendChild(cardCount);

        // Generates the card buttons based on server data
        var directions = ['up', 'left', 'down', 'right', 'self'];
        var directionMap = {
            'up': 'u',
            'left': 'l',
            'down': 'd',
            'right': 'r',
            'self': 's',
            'channel': 'c'
        };

        /**
         * Handles (custom) click events on card buttons to automatically
         * append the correct action string to the action queue
         *
         * @param {string} dir One of up, down, left, right, self, channel
         */
        var cardButtonClickHandler = function(dir) {
            var currentQueue = [];

            if (queueNode.textContent !== '') {
                currentQueue = queueNode.textContent.split(',');
            }

            currentQueue.push(action.ActionString + directionMap[dir]);
            queueNode.textContent = currentQueue.join(',');
        };

        for (var i = 0; i < directions.length; i++) {
            if (action.AvailableDirections[i] === true) {
                buttons.push(st.hand.CardButton(directions[i],
                            cardButtonClickHandler));

            } else {
                buttons.push(st.hand.CardButton(directions[i], null));
            }
        }

        // Channeling is always available
        buttons.push(st.hand.CardButton('channel', cardButtonClickHandler));

        for (var i = 0; i < buttons.length; i++) {
            var buttonRoot = buttons[i].root;
            buttonRoot.style.top = (178 + 147 * i) + 'px';
            container.appendChild(buttonRoot);
        }

        container.onclick = function() {
            if (hand[hand.length - 1] !== me) {
                var before = [];
                for (var i = 0; hand[i] !== me; i++) {
                    before.push(hand[i]);
                }

                var after = [];
                for (i++; i < hand.length; i++) {
                    after.push(hand[i]);
                }

                hand = before.concat(after);
                hand.push(me);
                arrangeHand();
            }
        };

        return me;
    };

    var windowOnResize = function() {
        containerDiv.style.zoom = "" + window.innerWidth / 780;
    };

    window.addEventListener('resize', windowOnResize);

    st.dom.getById("start-link").onclick = function() {
        windowOnResize();
        st.dom.removeElement(st.dom.getById("start-link"));
        st.dom.removeClass(st.dom.getById("top-level-buttons"), "hidden");
        conn.sendStartGame();
        conn.sendTick(++clientTick);
    };

    readyNode.addEventListener('click', function() {
        if (queueNode.textContent === '') {
            // TODO:athorp:2014-10-23 Make some user-facing UI.
            // Disable/enable button visually?
            console.log('Cant send empty queue');
        } else {
            conn.sendQueue(++clientTick, queueNode.textContent.split(','));
        }
    });

    undoNode.addEventListener('click', function() {
        if (queueNode.textContent !== '') {
            var currentQueue = queueNode.textContent.split(',');
            currentQueue.pop();
            queueNode.textContent = currentQueue.join(',');
        }
    });

    /**
     * Handles the version message.
     *
     * @param {st.define.Version} serverState
     */
    var handleVersionServerMessage = function(serverState) {
        console.log('Game version is', serverState.Version);
    };

    /**
     * Handles the general "world state changed" message and updates
     * everything.
     *
     * @param {st.define.WorldState} serverState
     */
    var handleWorldTick = function(serverState) {
        var myEntity = null;
        var myEntityID = serverState.Players[STABBEY.PLAYER].EntityId;
        var myActions = serverState.Players[STABBEY.PLAYER].AvailableActions;

        for (var i = 0; i < serverState.Entities.length; i++) {
            if (serverState.Entities[i].EntityId === myEntityID) {
                myEntity = serverState.Entities[i];
            }
        }

        // Update available cards / abilities
        hand = [];
        for (var i = 0; i < myActions.length; i++) {
            hand.push(st.hand.Card(myActions[i]));
        }
        buildHand();

        // Update existing (running) queue for this player & do tick
        var myServerQueue = myEntity.ActionQueue;
        queueNode.textContent = myServerQueue.join(',');
        if (myServerQueue.length > 0) {
            conn.sendTick(++clientTick);
            readyNode.classList.add('disabled');
        } else {
            readyNode.classList.remove('disabled');
        }
    };

    /**
     * Figures out the type of message that the server is sending and
     * routes it to the correct function.
     *
     * @param {?} serverState
     */
    var handleServerMessage = function(serverState) {
        if (serverState.Version) {
            handleVersionServerMessage(serverState);
        } else if (serverState.LastTick && serverState.Entities) {
            handleWorldTick(serverState);
        }
    };

    // Has to be defined at the end to have handleServerTick available
    var conn = st.Connection("Hand", STABBEY.HOST, STABBEY.PLAYER,
            handleServerMessage, null);
})();
