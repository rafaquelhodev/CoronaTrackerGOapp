function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.watchPosition(function (position) {
            console.log(position);
            showPosition(position);
        }, function (error) {
            console.log(error);
        }, { enableHighAccuracy: true, maximumAge: 3000, timeout: 3000 });
    }
}

function showPosition(position) {
    document.getElementById("latitude").value = position.coords.latitude;
    document.getElementById("longitude").value = position.coords.longitude;
}

getLocation();