package handlers

import "net/http"

func HandleIndex(response http.ResponseWriter, request *http.Request) {
	index := `<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Websocket example</title>
    <script>
    var exampleSocket = null;
    function startConnection() {
    	if (exampleSocket == null) {

        exampleSocket = new WebSocket("ws://localhost:3000/ws");
			exampleSocket.onmessage = function(event) {
				var msg = event.data;

				var existingMessages = document.getElementById("messages").innerHTML;
				var newMessage = "<div>"+msg+"</div>";
				document.getElementById("messages").innerHTML = newMessage + existingMessages;
			};
		}
	}

	function sendMessage() {
		if(exampleSocket != null) {
			exampleSocket.send("Hi ! " + new Date());
		}
	};

	function closeConnection() {
		if(exampleSocket != null) {
			console.log("closing ...");
			exampleSocket.close();
			exampleSocket = null;
		}
	};
    </script>
</head>
<body>
<h1>Hello from a static page</h1>
<h2>Messages</h2>
<button id="add" onclick="startConnection()">Start connection</button>
<button id="send" onclick="sendMessage()">Send message</button>
<button id="close" onclick="closeConnection()">Close connection</button>
<div id="messages">
</div>
</body>
</html>
`
	response.Write([]byte(index))
}

