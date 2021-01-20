window.addEventListener("DOMContentLoaded", () => {
    setRoomIdAndName();
    pullInWhoFirst();
});


function pullInWhoFirst() {
    let listElement = document.getElementById("meFirst");
    if (listElement === null) {
        return;
    }
    setInterval(() => {
        GET(`/api/pull/${document.roomId}`, (json) => {
            console.log(json);
            let state = json.gameState;
            if (state === STATE_CHOOSE_CARDS) {
                window.location.href = "/room/play.html";
            }
        }, (json, status) => {
            //todo handle
            console.log(json);
            console.log(status)
        });
    }, 1000);
}


function meFirst() {
    let btn = document.getElementById("meFirst");
    if (btn.classList.contains("disabled")) {
        return;
    }
    btn.classList.add("disabled");

    POST("/api/me-first", {roomId: document.roomId, playerName: document.playerName}, () => {
        },
        (body, status) => {
            btn.classList.remove("disabled");
            //todo handle error
            console.log(body);
            console.log(status);
        });
}