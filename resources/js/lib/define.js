goog.provide('st.define');

/* Holds commonly used variables and configuration */
goog.scope(function() {
    /** const */
    st.define.CARD_ACTIONS = {
        LEFT: 1,
        DOWN: 2,
        UP: 4,
        RIGHT: 8,
        SELF: 16,
        CHANNEL: 32,
        ALL: 63,
        DIRS: 15
    };
});
