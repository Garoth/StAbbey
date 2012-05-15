TICK_NUM = 0;
RESPONSE_TICK_NUM = 0;
TICK_DURATION = 3000;

socketOpened = function() {
    console.log("socket opened");
}

socketClosed = function() {
    console.log("socket closed");
}

socketErrored = function() {
    console.log("socket errored!!!");
}

socketMessaged = function(msg) {
    jsObj = $.parseJSON(msg.data);
    console.log("socket message:");
    console.log(jsObj);
    RESPONSE_TICK_NUM = jsObj.LastTick;
}

sendMessage = function(path, urlvars) {
    urlvars.gamekey = gamekey;
    urlvars.player = me;
    var first = true
    for (var variable in urlvars) {
        if (first === true) {
            path += "?";
            first = false
        } else {
            path += "&";
        }
        path += variable + "=" + urlvars[variable];
    }

    var xhr = new XMLHttpRequest();
    console.log("Sending message: " + path);
    xhr.open("POST", path, true);
    xhr.send();
}

tick = function() {
    if (TICK_NUM === RESPONSE_TICK_NUM) {
      TICK_NUM += 1;
    }
    sendMessage("/update", { ticknum: TICK_NUM });
    setTimeout(tick, TICK_DURATION);
}

$(function() {
    $("#test-send").click(function() {
        console.log("Trying to send message");
        sendMessage("/update", {startGame: true});
        tick();
    });
})
