goog.provide('st.define');

/* Holds commonly used variables and configuration */
/**
 * @typedef {{
 *      ActionString: string,
 *      AvailableDirections: Array.<boolean>,
 *      LongDescription: string,
 *      ShortDescription: string
 * }}
 */
st.define.Action;

/**
 * @typedef {{
 *      AvailableActions: Array.<st.define.Action>,
 *      EntityId: number,
 *      Id: number
 * }}
 */
st.define.Player;

/**
 * @typedef {{
 *      Layers: Array.<string>,
 *      Level: number
 * }}
 */
st.define.Board;

/**
 * @typedef {{
 *      ActionQueue: Array.<string>,
 *      Ardour: number,
 *      BoardId: number,
 *      EntityId: number,
 *      MaxArdour: number,
 *      Name: string,
 *      Type: string,
 *      Subtype: string,
 *      X: number,
 *      Y: number
 * }}
 */
st.define.Entity;

/**
 * @typedef {{
 *      Boards: Array.<st.define.Board>,
 *      CurrentBoard: number,
 *      Entities: Array.<st.define.Entity>,
 *      LastTick: number,
 *      Players: Array.<st.define.Player>
 * }}
 */
st.define.WorldState;

/**
 * @typedef {{
 *      Version: string
 * }}
 */
st.define.Version;
