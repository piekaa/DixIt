let roomId;
let name;
window.addEventListener("DOMContentLoaded", () => {
    setRoomIdAndName();
    pullInLobby();
    pullInWhoFirst();
});

function setRoomIdAndName() {
    let params = new URLSearchParams(window.location.search);
    if (params.has("id")) {
        roomId = params.get("id");
        localStorage.setItem("roomId", roomId);
    } else {
        let roomIdFromStorage = localStorage.getItem("roomId");
        if (roomIdFromStorage !== null) {
            roomId = roomIdFromStorage;
        }
    }
    let nameFromStorage = localStorage.getItem("name");
    if (nameFromStorage !== null) {
        name = nameFromStorage;
    }
}

function setName() {
    let name = document.getElementById("name").value;

    if (name.length === 0) {
        console.log("Name length is 0");
        //todo pop up message or validation error
        return;
    }

    POST("/api/name", {
            name: name,
            roomId: roomId
        },
        (json) => {
            localStorage.setItem("name", name);
            document.location.href = "/room/lobby.html";
        },
        (json, status) => {
            //todo handle
            console.log(json);
            console.log(status);
        });
}

function pullInLobby() {
    let listElement = document.getElementById("players");
    if (listElement === null) {
        return;
    }
    fillPlayersListOrStartGame(listElement);
    setInterval(() => {
        fillPlayersListOrStartGame(listElement);
    }, 3000);
}

function pullInWhoFirst() {
    let listElement = document.getElementById("meFirst");
    if (listElement === null) {
        return;
    }
    setInterval(() => {
        GET(`/api/pull/${roomId}`, (json) => {
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

function fillPlayersListOrStartGame(listElement) {
    GET(`/api/pull/${roomId}`, (json) => {
        console.log(json);

        let state = json.gameState;

        if (state === STATE_LOBBY) {
            let players = json.payload.players;

            while (listElement.firstChild) {
                listElement.removeChild(listElement.firstChild);
            }

            for (let playersKey in players) {
                let li = document.createElement("li");
                if (players.hasOwnProperty(playersKey)) {
                    li.innerText = players[playersKey].name;
                    if (players[playersKey].readyToStart) {
                        li.classList.add("ready");
                    }
                    listElement.appendChild(li);
                }
            }
        }
        if (state === STATE_WHO_FIRST) {
            window.location.href = "/room/who-first.html";
        }
    }, (json, status) => {
        //todo handle
        console.log(json);
        console.log(status)
    });
}

function ready() {
    let btn = document.getElementById("ready");
    if (btn.classList.contains("disabled")) {
        return;
    }
    btn.classList.add("disabled");

    POST("/api/ready", {roomId, name}, () => {
        },
        (body, status) => {
            btn.classList.remove("disabled");
            //todo handle error
            console.log(body);
            console.log(status);
        });
}

function copyInvitationLink() {
    //todo Popup
    let link = window.location.host + `/room/join.html?id=${roomId}`;
    let element = document.getElementById("copyToClipboard");
    element.value = link;
    element.select();
    document.execCommand("copy");
    window.getSelection().removeAllRanges();
}

function meFirst() {
    let btn = document.getElementById("meFirst");
    if (btn.classList.contains("disabled")) {
        return;
    }
    btn.classList.add("disabled");

    POST("/api/me-first", {roomId, name}, () => {
        },
        (body, status) => {
            btn.classList.remove("disabled");
            //todo handle error
            console.log(body);
            console.log(status);
        });
}