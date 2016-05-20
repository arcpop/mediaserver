function loadFolder(dir) {
    var xmlHttp = new XMLHttpRequest();
    var url = "/data" + dir;
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4) {
            if (xmlHttp.status != 200) {
                return
            }
            var response = JSON.parse(xmlHttp.responseText);
            renderDirectory(dir, response);
        }
    }
    xmlHttp.open("GET", url, true);
    xmlHttp.send();
}

function getParent(d) {
    return d.slice(0, d.lastIndexOf("/") - 1)
}

function renderDirectory(directoryLink, content) {
    str = "";
    if (directoryLink != "/") {
        upper = getParent(directoryLink)
        str += "<tr>"
        str += "<td><a href='#' onclick='loadFolder(" + 
            upper + ")'>..</a></td>";
        str += "</tr>";
    }
    for (d in content.dirs) {
        str += "<tr>";
        str += "<td><a href='#' onclick='loadFolder(" + 
            directoryLink + "/" + d.name + ")'>" + d.name + "</a></td>";
        str += "<td> - </td>";
        str += "<td>" + d.children + "</td>";
        str += "</tr>";
    }
    for (f in content.files) {
        str += "<tr>";
        str += "<td><a href='#' onclick='loadFolder(" + 
            directoryLink + "/" + f.name + ")'>" + f.name + "</a></td>";
        str += "<td>" + f.length + "</td>";
        str += "<td>" + f.details + "</td>";
        str += "</tr>";
    }
    $('#directory-table-body').innerHTML = str;
}