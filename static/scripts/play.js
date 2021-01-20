window.addEventListener("DOMContentLoaded", () => {
    setRoomIdAndName();
    pull();
    setInterval(pull, 2000);
});

let prevState = "";

let mine = 0;
let theirs = 0;
let amISpeaker = false;


function pull() {
    GET(`/api/pull/${document.roomId}`, (json) => handlePull(json), (json, status) => {
        //todo handle error
        console.log(json);
        console.log(status);
    });
}

function handlePull(json) {

    let state = json.gameState;
    let changed = prevState !== state;

    if (json.gameState === STATE_CHOOSE_CARDS) {
        pullInChooseCards(json, changed);
    } else if (json.gameState === STATE_ROUND_RESULT) {
        pullInRoundResult(json, changed);
    } else if (json.gameState === STATE_ROUND_FALSE_RESULT && changed) {
        alert("Ktoś oszukał! (Dwie lub więcej osób wybrały tą smą kartę jako swoją)");
        readyCall(document.createElement("div"));
    }
    prevState = state;
}

function pullInChooseCards(json, changed) {

    let main = document.getElementById("mainCardContainer");
    let players = getSortedPlayers(json);

    if (changed) {
        showColumns();
        showAcceptButton();
        removeAllChilds(main);

        if (document.playerName === json.payload.activePlayer) {
            document.getElementById("speaker").innerText = `Ty wymyślasz`;
            amISpeaker = true;
        } else {
            document.getElementById("speaker").innerText = `Wymyśla ${json.payload.activePlayer}`;
            amISpeaker = false;
        }

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
            if (!amISpeaker) {
                right.onclick = onCardClick;
            }

            center.innerText = i + 1;

            div.appendChild(left);
            div.appendChild(right);
            div.appendChild(center);

            main.appendChild(div);
        }
    }
}

function pullInRoundResult(json, changed) {

    let main = document.getElementById("mainCardContainer");
    let players = getSortedPlayers(json);
    let playersByMyCard = getPlayersByMyCard(json);
    let activePlayer = json.payload.activePlayer;

    if (changed) {
        hideColumns();
        showGoNextButton();
        removeAllChilds(main);

        for (let i = 0; i < players.length; i++) {
            let div = document.createElement("div");

            div.innerText = playersByMyCard[i + 1].playerName;

            div.classList.add("card-to-choose");
            if (activePlayer === playersByMyCard[i + 1].playerName) {
                div.classList.add("gold");
                div.style.backgroundImage = "none";
            }
            div.innerText += ": " + playersByMyCard[i + 1].score;

            if (document.playerName === playersByMyCard[i + 1].playerName) {
                div.classList.add("green");
                div.style.backgroundImage = "none";
            }


            main.appendChild(div);

            players.forEach((player) => {
                if (player.vote === i + 1 && player.playerName !== activePlayer) {
                    let divVoted = document.createElement("div");
                    divVoted.innerText = `${player.playerName}`;
                    if (player.playerName === name) {
                        divVoted.classList.add("green-text");
                    }
                    main.appendChild(divVoted);
                }
            });
        }
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

function getPlayersByMyCard(json) {
    let playersMap = json.payload.players;
    let players = [];

    for (let key in playersMap) {
        if (playersMap.hasOwnProperty(key)) {
            let player = playersMap[key];
            players[player.myCard] = player;
        }
    }
    return players;
}

function onCardClick() {
    let parent = this.parentElement;
    let position = this.getAttribute("position");
    if (this.classList.contains("card-to-choose__left")) {

        if (mine !== 0 && mine !== position) {
            let old = document.getElementById(`left-${mine}`).parentElement;
            old.classList.remove("card-to-choose__mine");
        }

        if (theirs === position) {
            theirs = 0;
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

        if (mine === position) {
            mine = 0;
        }

        parent.classList.remove("card-to-choose__mine");
        parent.classList.add("card-to-choose__theirs");
        theirs = position;
    }

    let acceptButton = document.getElementById("accept");
    if (mine > 0 && (theirs > 0 || amISpeaker)) {
        acceptButton.classList.remove("disabled");
    } else {
        acceptButton.classList.add("disabled");
    }
}

function onAccept() {
    let acceptButton = document.getElementById("accept");

    if (acceptButton.classList.contains("disabled")) {
        return;
    }

    acceptButton.classList.add("disabled");
    POST("/api/choose-cards", {
            playerName: document.playerName,
            roomId: document.roomId,
            myCard: parseInt(mine),
            myType: amISpeaker ? parseInt(mine) : parseInt(theirs)
        },
        (json) => {
            console.log(json);
        },
        (json, status) => {
            //todo
            console.log(json);
            console.log(status);
        })
}

function onGoNext() {
    let button = document.getElementById("goNext");
    if (button.classList.contains("disabled")) {
        return;
    }
    button.classList.add("disabled");
    readyCall(button);
}

function showColumns() {
    document.getElementById("leftColumn").style.display = "flex";
    document.getElementById("rightColumn").style.display = "flex";
}

function hideColumns() {
    document.getElementById("leftColumn").style.display = "none";
    document.getElementById("rightColumn").style.display = "none";
}

function showAcceptButton() {
    document.getElementById("accept").classList.remove("hidden");
    document.getElementById("goNext").classList.add("hidden");
}

function showGoNextButton() {
    document.getElementById("accept").classList.add("hidden");
    document.getElementById("goNext").classList.remove("hidden");
    document.getElementById("goNext").classList.remove("disabled");
}