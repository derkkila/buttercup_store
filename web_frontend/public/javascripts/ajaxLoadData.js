$(document).ready(function() {

	ajaxGet();

	// DO GET
	function ajaxGet(){
		$.ajax({
			type : "GET",
			url : "http://"+window.location.hostname+":6767/products/",
			success: function(result){
				$.each(result, function(i, product){

					var productRow = '<tr>' +
                    '<td><a href="/shop/'+product.id+'">'+ product.name + '</a></td>' +
                    '<td>' + product.type + '</td>' +
                    '<td>' + product.category + '</td>' +
                    '<td>' + product.price + '</td>' +
                    '<td>' + product.qty + '</td> ' +
                    '</tr>';

					$('#productList tbody').append(productRow);

		        });

				$( "#productList tbody tr:odd" ).addClass("info");
				$( "#productList tbody tr:even" ).addClass("success");
			},
			error : function(e) {
				alert("ERROR: ", e);
				console.log("ERROR: ", e);
			}
		});
	}
})
