let id = localStorage.getItem("id");

fetch("http://127.0.0.1:8080/order/" + id)
    .then(response => response.json())
    .then(data => {
        document.getElementById("myText").innerHTML = JSON.stringify(data, null, 4);
    })


