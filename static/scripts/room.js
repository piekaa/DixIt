window.addEventListener("DOMContentLoaded", () => {
    setRoomIdAndName();
    fillCopyToClipboardAndQrCode();
    pullInLobby();
});

function fillCopyToClipboardAndQrCode() {
    let protocolAndHost = window.location.protocol + "//" + window.location.host;
    let link = protocolAndHost + `/room/join.html?id=${document.roomId}`;
    let qrLink = protocolAndHost + `/api/qr?url=${protocolAndHost}/room/join.html?id=${document.roomId}`;
    let element = document.getElementById("copyToClipboard");
    element.value = link;
    let qrImage = document.getElementById("qr");
    qrImage.src = qrLink;
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


function fillPlayersListOrStartGame(listElement) {
    GET(`/api/pull/${document.roomId}`, (json) => {
        console.log(json);

        let state = json.gameState;

        if (state === STATE_LOBBY) {
            let players = json.payload.players;

            removeAllChilds(listElement);

            for (let playersKey in players) {
                let li = document.createElement("li");
                if (players.hasOwnProperty(playersKey)) {
                    li.innerText = players[playersKey].playerName;
                    if (players[playersKey].readyToStart) {
                        li.classList.add("ready");
                    }
                    listElement.appendChild(li);
                }
            }
        }
        if (state === STATE_WHO_FIRST) {
            window.location.href = "/room/who-first.html";
        } else if (state === STATE_CHOOSE_CARDS) {
            window.location.href = "/room/play.html";
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
    readyCall(btn);
}

function copyInvitationLink() {
    //todo Popup
    let link = window.location.host + `/room/join.html?id=${document.roomId}`;
    let element = document.getElementById("copyToClipboard");
    element.value = link;
    element.select();
    document.execCommand("copy");
    window.getSelection().removeAllRanges();
}