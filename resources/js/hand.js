goog.require("st.define");
goog.require("st.dom");
goog.require("st.Connection");

goog.provide("st.hand");

goog.scope(function() {
    var BUTTON_PREFIX = '/resources/img/card/btn_';
    var hand = [];

    st.hand.CardButton = function(type, onClick) {
        var me = {};
        me.buttonName = BUTTON_PREFIX + type;
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
                onClick();
            }
        };

        return me;
    };

    st.hand.Card = function(name, actions, image, count) {

        var me = {};
        me.root = st.dom.createDiv('', 'card-wrapper');
        var container = st.dom.createDiv('', 'card');
        var artImg = st.dom.createImg('', 'art', image);
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

        me.clickLeft = function() {
        };

        me.clickDown = function() {
        };

        me.clickUp = function() {
        };

        me.clickRight = function() {
        };

        me.clickSelf = function() {
        };

        me.clickChannel = function() {
        };

        buttons.push(st.hand.CardButton('up',
            actions & st.define.CARD_ACTIONS.UP ? me.clickUp : null));
        buttons.push(st.hand.CardButton('left',
            actions & st.define.CARD_ACTIONS.LEFT ? me.clickLeft : null));
        buttons.push(st.hand.CardButton('right',
            actions & st.define.CARD_ACTIONS.RIGHT ? me.clickRight : null));
        buttons.push(st.hand.CardButton('down',
            actions & st.define.CARD_ACTIONS.DOWN ? me.clickDown : null));
        buttons.push(st.hand.CardButton('self',
            actions & st.define.CARD_ACTIONS.SELF ? me.clickSelf : null));
        buttons.push(st.hand.CardButton('channel',
            actions & st.define.CARD_ACTIONS.CHANNEL ? me.clickChannel : null));

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

    var handleServerTick = function(serverState) {
        if (serverState.Players) {
            hand = [];
            var actions = serverState.Players[STABBEY.PLAYER].AvailableActions;

            for (var i = 0; i < actions.length; i++) {
                var action = actions[i];
                var dirs = st.define.CARD_ACTIONS.CHANNEL;

                if (action.AvailableDirections[0]) {
                    dirs |= st.define.CARD_ACTIONS.LEFT;
                }

                if (action.AvailableDirections[1]) {
                    dirs |= st.define.CARD_ACTIONS.RIGHT;
                }

                if (action.AvailableDirections[2]) {
                    dirs |= st.define.CARD_ACTIONS.UP;
                }

                if (action.AvailableDirections[3]) {
                    dirs |= st.define.CARD_ACTIONS.DOWN;
                }

                if (action.AvailableDirections[4]) {
                    dirs |= st.define.CARD_ACTIONS.SELF;
                }

                hand.push(st.hand.Card(action.ShortDescription, dirs,
                    '/resources/img/card/card_art_001.png', '&infin;'));
            }
            buildHand();
        }
    };

    var clientTick = 0;
    var containerDiv = st.dom.getById('cards');
    var conn = st.Connection("Hand",
        STABBEY.HOST, STABBEY.PLAYER, handleServerTick, null);

    window.onresize = windowOnResize;

    st.dom.getById("start-link").onclick = function() {
        windowOnResize();
        st.dom.removeElement(st.dom.getById("start-link"));
        conn.sendStartGame();
        conn.sendTick(++clientTick);
    };

    var windowOnResize = function() {
        containerDiv.style.zoom = "" + window.innerWidth / 780;
    };

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

});
