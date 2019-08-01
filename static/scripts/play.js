window.addEventListener("DOMContentLoaded", () => {
    pull();
    setInterval(pull, 2000);
});

let prevState = "";

let mine = 0;
let theirs = 0;


function pull() {
    GET(`/api/pull/${roomId}`, (json) => handlePull(json), (json, status) => {
        //todo handle error
        console.log(json);
        console.log(status);
    });
}

function handlePull(json) {

    let state = json.gameState;
    let chaned = prevState !== state;

    if (json.gameState === STATE_CHOOSE_CARDS) {
        pullInChooseCards(json, chaned);
    }
    prevState = state;
}

function pullInChooseCards(json, changed) {

    let main = document.getElementById("mainCardContainer");
    let players = getSortedPlayers(json);

    if (changed) {
        removeAllChilds(main);

        document.getElementById("speaker").innerText = json.payload.activePlayer;

        for (let i = 0; i < players.length; i++) {
            let div = document.createElement("div");
            div.classList.add("card-to-choose");

            let left = document.createElement("div");
            let center = document.createElement("div");
            let right = document.createElement("div");

            left.classList.add("card-to-choose__left");
            center.classList.add("card-to-choose__center");
            right.classList.add("card-to-choose__right");

            left.id = `left-${i + 1}`;
            right.id = `right-${i + 1}`;

            left.setAttribute("position", i + 1);
            right.setAttribute("position", i + 1);

            left.onclick = onCardClick;
            right.onclick = onCardClick;

            center.innerText = i + 1;

            div.appendChild(left);
            div.appendChild(right);
            div.appendChild(center);

            main.appendChild(div);
        }
    } else {

    }
}

function getSortedPlayers(json) {
    let playersMap = json.payload.players;
    let players = [];
    let i = 0;

    for (let key in playersMap) {
        if (playersMap.hasOwnProperty(key)) {
            players[i++] = playersMap[key];
        }
    }
    return players.sort((a, b) => a.position - b.position);
}

function onCardClick() {
    let parent = this.parentElement;
    let position = this.getAttribute("position");
    if (this.classList.contains("card-to-choose__left")) {

        if (mine !== 0 && mine !== position) {
            let old = document.getElementById(`left-${mine}`).parentElement;
            old.classList.remove("card-to-choose__mine");
        }

        parent.classList.remove("card-to-choose__theirs");
        parent.classList.add("card-to-choose__mine");
        mine = position;
    }
    if (this.classList.contains("card-to-choose__right")) {

        if (theirs !== 0 && theirs !== position) {
            let old = document.getElementById(`right-${theirs}`).parentElement;
            old.classList.remove("card-to-choose__theirs");
        }

        parent.classList.remove("card-to-choose__mine");
        parent.classList.add("card-to-choose__theirs");
        theirs = position;
    }
}
