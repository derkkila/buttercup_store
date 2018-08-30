var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('error', { message: 'Buttercup Games Root' });
});

router.get('/shop', function(req, res, next) {
  res.render('index', { title: 'Buttercup Games' });
});

router.get('/shop/cart', function(req, res, next) {
  var id = 1;
  const http = require('http');
  console.log('http://cartservice:4201/cart/'+id)

  http.get('http://cartservice:4201/cart/'+id, (resp) => {
    let data = '';

    // A chunk of data has been recieved.
    resp.on('data', (chunk) => {
      data += chunk;
    });

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      var cart=JSON.parse(data);
      res.render('cart_view', { title: "Cart", cart: cart })
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});

router.get('/shop/:id', function(req, res, next) {
  var id = req.params.id;
  const http = require('http');

  http.get('http://productservice:6767/products/'+id, (resp) => {
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
