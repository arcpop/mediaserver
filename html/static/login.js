var currentUser = "";
function doLogin(params) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4) {
            var ok = false;
            var reason = "Connection to server failed";
            if (xmlHttp.status == 200) {
                var response = JSON.parse(xmlHttp.responseText)
                if (response.login == "ok") {
                    ok = true;
                    reason = "";
                    window.location = "https://192.168.56.100/user/" + response.username + "/";
                } else if (response.reason != "") {
                    reason = response.reason
                }
            }
            if (!ok) {
                alert(reason);
            }
        }
    }
    var url = "/login?username=" + document.getElementById("inputEmail").value +
        "&password=" + document.getElementById("inputPassword").value;
    xmlHttp.open("GET", url, true);
    xmlHttp.send();
}
