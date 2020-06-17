var hiddenInputClient = document.getElementById("IDclient");

var elem = document.getElementById("declareinfection");
elem.href = "declareinfection/" + hiddenInputClient.value;

var buttonTracker = document.getElementById("trackposition");
buttonTracker.href = "trackposition/" + hiddenInputClient.value;