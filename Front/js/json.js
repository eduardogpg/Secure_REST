$(document ).ready(function() {
		
	$("#target").on("submit", function(e){	
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "http://127.0.0.1:8001/message",
			data: {
				"message" : "new message send"
			},
			success: function(data){
				show_result(data)
			}
		})
	})

	function show_result(data){
		alert("Si llego algo")
	}

})