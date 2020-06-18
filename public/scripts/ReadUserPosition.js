var positionJson = {
    user: -1,
    lati: 1.0,
    long: 1.0
}

var pathArray = window.location.pathname.split('/');
userID = pathArray[pathArray.length - 1];

var positionJson = {
    user: parseInt(userID),
    lati: 1.0,
    long: 1.0
}

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.watchPosition(function (position) {
            showPosition(position);
            postPositionWithAjax(positionJson, position);
        }, function (error) {
            console.log(error);
        }, { enableHighAccuracy: true, maximumAge: 3000, timeout: 3000 });
    }
}

function showPosition(position) {
    document.getElementById("latitude").value = position.coords.latitude;
    document.getElementById("longitude").value = position.coords.longitude;
}

function postPositionWithAjax(positionJson, position) {
    positionJson.lati = position.coords.latitude;
    positionJson.long = position.coords.longitude;

    $.ajax({
        url: "http://localhost:8090/usercoordinates",
        type: "POST",
        dataType: "json",
        data: JSON.stringify(positionJson)
    });
}

getLocation();