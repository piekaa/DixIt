function createRoom() {
    //todo handle error
    POST("/api/room", {}, (json) => {
        localStorage.setItem("roomId", json.roomId);
        window.location.href = `/room/join.html?id=${json.roomId}`;
    });
}