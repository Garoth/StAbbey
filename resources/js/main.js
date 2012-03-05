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
    console.log("socket message:");
    console.log(msg);
    $("body").append("got msg: " + msg);
}

/* TODO urlvars not implemented */
sendMessage = function(path, urlvars) {
    path += "?gamekey=" + gamekey;

    var xhr = new XMLHttpRequest();
    xhr.open("POST", path, true);
    xhr.send();
}

$(function() {
    $("#test-send").click(function() {
        console.log("Trying to send message");
        sendMessage("/update", "foo");
    });
})
