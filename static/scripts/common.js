// alert("common");

function GET(path, successCallback, errorCallback) {
    fetch(path)
        .then((res) => {
            let code = res.status;
            if (res.ok) {
                res.json().then((json) => {
                    successCallback(json, status);
                })
            } else {
                res.json().then((json) => {
                    errorCallback(json, status);
                })
            }
        });
}


function POST(path, body, successCallback, errorCallback) {
    fetch(path, {
        headers: {
            'Content-Type': "application/json"
        },
        method: "POST",
        body: JSON.stringify(body)
    })
        .then((res) => {
            let code = res.status;
            if (res.ok) {
                res.json().then((json) => {
                    successCallback(json, status);
                })
            } else {
                res.json().then((json) => {
                    errorCallback(json, status);
                })
            }
        });

}

function removeAllChilds(listElement) {
    while (listElement.firstChild) {
        listElement.removeChild(listElement.firstChild);
    }
}


function setRoomIdAndName() {
    // alert("setRoomIdAndName");
    document.roomId = "unknownRoomId";
    document.playerName = "unknownPlayerName";
    let params = new URLSearchParams(window.location.search);
    if (params.has("id")) {
        document.roomId = params.get("id");
        localStorage.setItem("roomId", document.roomId);
    } else {
        let roomIdFromStorage = localStorage.getItem("roomId");
        if (roomIdFromStorage !== null) {
            document.roomId = roomIdFromStorage;
        }
    }
    let nameFromStorage = localStorage.getItem("playerName");
    if (nameFromStorage !== null) {
        document.playerName = nameFromStorage;
    }

    // alert("rid: " + document.roomId);
    // alert("name: " + document.playerName);

}

function readyCall(btn) {
    POST("/api/ready", {roomId: document.roomId, playerName: document.playerName}, () => {
        },
        (body, status) => {
            btn.classList.remove("disabled");
            //todo handle error
            console.log(body);
            console.log(status);
        });
}

const STATE_LOBBY = "lobby";
const STATE_WHO_FIRST = "whoFirst";
const STATE_CHOOSE_CARDS = "chooseCards";
//in case 2 or more players claim same card is theirs
const STATE_ROUND_FALSE_RESULT = "roundFalseResult";
const STATE_ROUND_RESULT = "roundResult";
//todo popup
