function addButton() {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", backendURL+"/increment", true);
    xhr.send()
    xhr.onload = function () {
        getButtons();
    }
};

function getButtons() {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", backendURL+"/getCount", true);
    xhr.send()
    xhr.onload = function () {
        drawButtons(JSON.stringify(JSON.parse(xhr.responseText),null,2));
    }
};

function drawButtons(count) {
    var html = []
    html.push(`
          <center>`);
    for( let i = 0; i < count; i++) {
        html.push(`
              <input type="button" value="Add Botton" onclick="addButton();" ></input>`)
    }
    html.push(`
          </center>`);
    document.getElementById("buttons").innerHTML = html.join("\n")
};

function clearButtons() {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", backendURL+"/clear", true);
    xhr.send()
    xhr.onload = function () {
        getButtons();
    }
};
