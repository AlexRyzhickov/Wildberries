let orders = document.getElementById('rating');

fetch("http://127.0.0.1:8080/orders")
    .then(response => response.json())
    .then(data => {
        for (var i = 0; i < data.length; i++) {
            var id = data[i];
            var li = document.createElement('li');
            li.appendChild(document.createTextNode(id));
            li.onclick = function () {
                localStorage.setItem("id", id)
                window.location.href = "order.html";
            }
            orders.appendChild(li);
        }
    })


