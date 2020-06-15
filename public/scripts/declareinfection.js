var pathArray = window.location.pathname.split('/');
userID = pathArray[pathArray.length - 1];

document.postInfectionDate.action = "postInfectionDate/" + userID;