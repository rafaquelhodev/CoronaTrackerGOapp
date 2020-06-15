var pathArray = window.location.pathname.split('/');
userID = pathArray[pathArray.length - 1];
console.log(userID)

var elem = document.getElementById("declareinfectionhome");

elem.href = "declareinfectionhome/" + userID;