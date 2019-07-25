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

//todo popup
