$(document).ready(function() {

	ajaxGet();

	var target = 0;
	switch( window.location.pathname )
	{
	    case "/shop":
	        target = 0;
	        break;

	    case "/shop/cart":
	        target = 1;
	        break;
	    /* add other cases */
	}

	$($("#navmenu a")[target]).addClass("active");

	// DO GET
	function ajaxGet(){
		$.ajax({
			type : "GET",
			url : "/products/",
			success: function(result){
				$.each(result, function(i, product){

					var productRow = '<tr>' +
                    '<td><a href="/shop/'+product.id+'">'+ product.name + '</a></td>' +
                    '<td>' + product.type + '</td>' +
                    '<td>' + product.category + '</td>' +
                    '<td>' + product.price + '</td>' +
                    '<td>' + product.qty + '</td>' +
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
