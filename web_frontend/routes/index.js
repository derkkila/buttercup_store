var express = require('express');
var uuid = require('uuid');

var router = express.Router();

router.use(require('cookie-parser')());

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('error', { message: 'Buttercup Games Root' });
});

router.get('/shop', function(req, res, next) {

  let bc = req.cookies.bc_session;

  if (!bc) {

        // crude id gen for now
        let transactionid=uuid.v1();
        res.cookie('bc_session', transactionid);
        req.cookies.bc_session = transactionid;

    }

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

router.post('/shop/checkout', function(req, res, next) {
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
      res.render('checkout_view', { title: "Checkout", cart: cart })
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});

router.get('/shop/thankyou', function(req, res, next) {
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
      res.render('thankyou_view', { title: "Thank You!", cart: cart })
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
      product=product[0];
      product.price = Number(product.price).toFixed(2);
      res.render('product_view', { title: product.name, id: product.id, name: product.name, description: product.description, prodtype: product.prodtype, category: product.category, price: product.price, qty: product.qty, filepath: product.filepath });
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});

module.exports = router;
