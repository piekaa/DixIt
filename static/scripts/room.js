let roomId;
window.addEventListener("DOMContentLoaded", () => {
    setRoomId();
    fillPlayersList();
});

function setRoomId() {
    let params = new URLSearchParams(window.location.search);
    if (params.has("id")) {
        roomId = params.get("id");
    } else {
        let roomIdFromStorage = localStorage.getItem("roomId");
        if (roomIdFromStorage !== null) {
            roomId = roomIdFromStorage;
        }
    }
}

function setName() {
    let name = document.getElementById("name").value;

    if (name.length === 0) {
        console.log("Name length is 0")
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

function fillPlayersList() {
    let listElement = document.getElementById("players");
    if (listElement === undefined) {
        return;
    }
    GET(`/api/pull/${roomId}`, (json) => {
        console.log(json);
        let players = json.payload.players;

        for (let playersKey in players) {
            let li = document.createElement("li");
            li.innerText = players[playersKey].name;
            listElement.appendChild(li);
        }

    }, (json, status) => {
        //todo handle
        console.log(json);
        console.log(status)
    })
}