var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.redirect('/shop');
});

router.get('/shop', function(req, res, next) {
  res.render('index', { title: 'Buttercup Games' });
});

router.get('/shop/cart', function(req, res, next) {
  var id = 1;
  const http = require('http');
  console.log('http://localhost:4201/cart/'+id)

  http.get('http://localhost:4201/cart/'+id, (resp) => {
    let data = '';

    // A chunk of data has been recieved.
    resp.on('data', (chunk) => {
      data += chunk;
    });

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      var product=JSON.parse(data);
      var total=0;

      var Promise = require('promise');


      console.log(product)

        var prom = new Promise(function(resolve, reject) {
          var price_total=0;
          product.forEach((product)=>{
          // Do async job
          http.get('http://localhost:6767/products/'+product.product_id, (resp) => {
            let data = '';

            // A chunk of data has been recieved.
            resp.on('data', (chunk) => {
              data += chunk;
            });

            // The whole response has been received. Print out the result.
            resp.on('end', () => {
              var product_details=JSON.parse(data);
              price_total+=product.qty*product_details.price;
              console.log(price_total);
            });

          }).on("error", (err) => {
            console.log("Error: " + err.message);
            reject(err);
          });
        });
        resolve(price_total);
  }).then(function(result) {
  total+=result;
  console.log("Final Total: "+total);
  res.render('cart_view', { title: "Cart", total: total });
  // Use user details from here
  console.log(userDetails)
}, function(err) {
  console.log(err);
});
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});

router.get('/shop/:id', function(req, res, next) {
  var id = req.params.id;
  const http = require('http');

  http.get('http://localhost:6767/products/'+id, (resp) => {
    let data = '';

    // A chunk of data has been recieved.
    resp.on('data', (chunk) => {
      data += chunk;
    });

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      var product=JSON.parse(data);
      res.render('product_view', { title: product.name, id: product.id, name: product.name, description: product.description, prodtype: product.prodtype, category: product.category, price: product.price, qty: product.qty });
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});

module.exports = router;
