$(document).ready(function() {

	productGet();
	orderGet();

	// DO GET
	function productGet(){
		$.ajax({
			type : "GET",
			url : "http://"+window.location.hostname+"/products/",
			success: function(result){
				$.each(result, function(i, product){

					var productRow = '<tr>' +
                    '<td>' + product.id + '</td>' +
                    '<td>' + product.name.toUpperCase() + '</td>' +
                    '<td>' + product.type + '</td>' +
                    '<td>' + product.category + '</td>' +
                    '<td>' + product.price + '</td>' +
                    '<td>' + product.qty + '</td>' +
                    '<td><a href="/admin/edit_product?id='+product.id+'">Edit</a></td>' +
                    '<td><a href="http://'+window.location.hostname+'/products/delete/'+product.id+'">Delete</a></td>' +
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
	function orderGet(){
		$.ajax({
			type : "GET",
			url : "http://"+window.location.hostname+"/orders/",
			success: function(result){
				$.each(result, function(i, order){

					var orderRow = '<tr>' +
                    '<td>' + order.id + '</td>' +
                    '<td>' + order.user_id + '</td>' +
                    '<td>' + order.total + '</td>' +
                    '<td>' + order.qty + '</td>' +
                    '<td><a href="/admin/edit_product?id='+order.id+'">Edit</a></td>' +
                    '<td><a href="http://'+window.location.hostname+'/orders/delete/'+order.id+'">Delete</a></td>' +
                    '</tr>';

					$('#orderList tbody').append(orderRow);

		        });

				$( "#orderList tbody tr:odd" ).addClass("info");
				$( "#orderList tbody tr:even" ).addClass("success");
			},
			error : function(e) {
				alert("ERROR: ", e);
				console.log("ERROR: ", e);
			}
		});
	}
})
