window.addEventListener("DOMContentLoaded", () => {
    setRoomIdAndName();
});

function setName() {
    // alert("działam!");
    let name = document.getElementById("name").value;

    if (name.length === 0) {
        console.log("Name length is 0");
        //todo pop up message or validation error
        return;
    }

    // alert("Roomid : " + document.roomId);

    POST("/api/playerName", {
            playerName: name,
            roomId: document.roomId
        },
        (json) => {
            // alert(json);
            localStorage.setItem("playerName", name);
            document.location.href = "/room/lobby.html";
        },
        (json, status) => {
            // alert(status);
            //todo handle
            console.log(json);
            console.log(status);
        });
}
