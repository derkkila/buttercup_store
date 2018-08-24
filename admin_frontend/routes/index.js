var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Hello World' });
});

router.get('/admin', function(req, res, next) {
  res.render('admin_index', { title: 'Admin'});
});

router.get('/admin/add_product', function(req, res, next) {
  res.render('admin_product', { title: 'Add Product'});
});

router.get('/admin/edit_product', function(req, res, next) {
  var id = req.query.id;
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
      res.render('admin_product', { title: 'Edit Product', id: product.id, name: product.name, description: product.description, prodtype: product.prodtype, category: product.category, price: product.price, qty: product.qty});
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
});



module.exports = router;
