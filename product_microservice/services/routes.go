package services
import (
  "net/http"
  "database/sql"
  _ "./mysql"
  "../model"
  "log"
  "fmt"
  "strconv"
  "encoding/json"
  "./mux-master"
)

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
Name        string
Method      string
Pattern     string
HandlerFunc http.HandlerFunc
}
// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route
// Initialize our routes
var routes = Routes{
  Route{
    "GetProduct",                                     // Name
    "GET",                                            // HTTP method
    "/products/{productId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /products/{productId}")
                db, err := sql.Open("mysql","root:test@tcp(127.0.0.1:3306)/products")

                var productId = mux.Vars(r)["productId"]

                var (
                  id int
                  name string
                  description string
                  prodtype string
                  category string
                  price float64
                  qty int
                )
                err2 := db.QueryRow("select *from product_list where id = ?", productId).Scan(&id, &name, &description, &prodtype, &category, &price, &qty)

                var status = http.StatusBadRequest

                switch {
                case err2 == sql.ErrNoRows:
                        log.Printf("No Product with that ID.")
                        status = http.StatusNoContent
                case err2 != nil:
                        log.Fatal(err)
                        status = http.StatusInternalServerError
                default:
                        fmt.Printf("Product Name is %s\n", name)
                        status = http.StatusOK
                }

                product := model.Product{
                  Id: strconv.Itoa(id),
                  Name: name,
                  Description: description,
                  ProdType: prodtype,
                  Category: category,
                  Price: strconv.FormatFloat(price, 'E', 2, 32),
                  Qty: strconv.Itoa(qty),
                }
                jsonBytes, _ := json.Marshal(product)
                //var output = fmt.Sprintf("%v",jsonBytes)
                log.Println(product.Id, product.Name)

                defer db.Close()
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(status)
                w.Write([]byte(jsonBytes))
        },
  },
  Route{
  "UpdateProduct",                                     // Name
  "POST",                                            // HTTP method
  "/products/{productId}",                          // Route pattern
  func(w http.ResponseWriter, r *http.Request) {
    log.Println("Updating /products/{productId}")

    var id = mux.Vars(r)["productId"]

    err := r.ParseForm()
  	if err != nil {
  		panic(err)
  	}
  	v := r.Form
  	product := r.Form.Get("name")

    log.Println(v)
    log.Println(product)

    db, errdb := sql.Open("mysql","root:test@tcp(127.0.0.1:3306)/products")
    if errdb != nil {
  		panic(errdb)
  	}

    stmt, err2 := db.Prepare("update product_list set name=?,description=?,type=?,category=?,price=?,qty=? where id=?")
    if err2 != nil {
  		panic(err2)
  	}

    res, err3 := stmt.Exec(r.Form.Get("name"),r.Form.Get("description"),r.Form.Get("prodtype"),r.Form.Get("category"),r.Form.Get("price"),r.Form.Get("qty"), id)
    if err3 != nil {
  		panic(err3)
  	}

    affect, err4 := res.RowsAffected()
    if err4 != nil {
  		panic(err4)
  	}

    fmt.Println(affect)


    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"result\":\"OK\"}"))
  },
  },
  Route{
  "DecreaseQty",                                     // Name
  "GET",                                            // HTTP method
  "/products/{productId}/{qty}",                          // Route pattern
  func(w http.ResponseWriter, r *http.Request) {
    log.Println("Updating Qty /products/{productId}")

    var id = mux.Vars(r)["productId"]
    var qty = mux.Vars(r)["qty"]

    db, errdb := sql.Open("mysql","root:test@tcp(127.0.0.1:3306)/products")
    if errdb != nil {
  		panic(errdb)
  	}

    stmt, err2 := db.Prepare("update product_list set qty=qty-? where id=?")
    if err2 != nil {
  		panic(err2)
  	}

    res, err3 := stmt.Exec(qty,id)
    if err3 != nil {
  		panic(err3)
  	}

    affect, err4 := res.RowsAffected()
    if err4 != nil {
  		panic(err4)
  	}

    fmt.Println(affect)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"result\":\"OK\"}"))
  },
  },
  Route{
    "AddProduct",                                     // Name
    "POST",                                            // HTTP method
    "/products/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
      log.Println("Posting /products/")

      err := r.ParseForm()
    	if err != nil {
    		panic(err)
    	}
    	v := r.Form
    	product := r.Form.Get("name")

      log.Println(v)
      log.Println(product)

      db, errdb := sql.Open("mysql","root:test@tcp(127.0.0.1:3306)/products")
      if errdb != nil {
    		panic(errdb)
    	}

      stmt, err2 := db.Prepare("INSERT product_list SET name=?,description=?,type=?,category=?,price=?,qty=?")
      if err2 != nil {
    		panic(err2)
    	}

      res, err3 := stmt.Exec(r.Form.Get("name"),r.Form.Get("description"),r.Form.Get("prodtype"),r.Form.Get("category"),r.Form.Get("price"),r.Form.Get("qty"))
      if err3 != nil {
    		panic(err3)
    	}

      id, err4 := res.LastInsertId()
      if err4 != nil {
    		panic(err4)
    	}

      log.Println(id)

      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("{\"result\":\"OK\"}"))
    },
  },
}
