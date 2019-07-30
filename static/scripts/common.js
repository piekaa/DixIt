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

const STATE_LOBBY = "lobby";
const STATE_WHO_FIRST = "whoFirst";
const STATE_CHOOSE_CARDS = "chooseCards";
//in case 2 or more players claim same card is theirs
const STATE_ROUND_FALSE_RESULT = "roundFalseResult";
const STATE_ROUND_RESULT = "roundResult";
//todo popup
