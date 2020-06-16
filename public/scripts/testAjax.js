var person = {
    name: "rafael",
    address: "sjc"
}
$.ajax({
    url: "http://localhost:8090/usercoordinates",
    type: "POST",
    dataType: "json",
    data: JSON.stringify(person)
});