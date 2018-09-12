package services
import (
  "net/http"
  "database/sql"
  _ "./mysql"
  "log"
  "fmt"
  "encoding/json"
  "./mux-master"
  "strings"
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
    "GetAllProducts",                                     // Name
    "GET",                                            // HTTP method
    "/products/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /products/")
                db, err := sql.Open("mysql","root:test@tcp(productdb:3306)/products")

                rows,err2 := db.Query("select *from product_list")

                var status = http.StatusBadRequest

                switch {
                case err2 == sql.ErrNoRows:
                        log.Printf("No Product with that ID.")
                        status = http.StatusNoContent
                case err2 != nil:
                        log.Fatal(err)
                        status = http.StatusInternalServerError
                default:
                        status = http.StatusOK
                }

                defer rows.Close()

                columns, err := rows.Columns()
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }

                count := len(columns)
                tableData := make([]map[string]interface{}, 0)
                values := make([]interface{}, count)
                valuePtrs := make([]interface{}, count)
                for rows.Next() {
                    for i := 0; i < count; i++ {
                        valuePtrs[i] = &values[i]
                    }
                    rows.Scan(valuePtrs...)
                    entry := make(map[string]interface{})
                    for i, col := range columns {
                        var v interface{}
                        val := values[i]
                        b, ok := val.([]byte)
                        if ok {
                            v = string(b)
                        } else {
                            v = val
                        }
                        entry[col] = v
                    }
                    tableData = append(tableData, entry)
                }
                jsonData, err := json.Marshal(tableData)
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }
                fmt.Println(string(jsonData))

                defer db.Close()

                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.WriteHeader(status)
                w.Write([]byte(jsonData))
        },
  },
  Route{
    "DeleteProduct",                                     // Name
    "GET",                                            // HTTP method
    "/products/delete/{productId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Deleting /products/{productId}")
                var productId = mux.Vars(r)["productId"]

                db, err := sql.Open("mysql","root:test@tcp(productdb:3306)/products")

                stmt, err := db.Prepare("delete from product_list where id = ?")

                res, err := stmt.Exec(productId)

                affect, err := res.RowsAffected()

                log.Println(affect)

                var status = http.StatusBadRequest

                switch {
                case err == sql.ErrNoRows:
                        log.Printf("No Product with that ID.")
                        status = http.StatusNoContent
                case err != nil:
                        log.Fatal(err)
                        status = http.StatusInternalServerError
                default:
                        status = http.StatusOK
                }

                log.Println(status)

                defer db.Close()

                log.Println("Redirect back to admin")
                http.Redirect(w, r, "http://localhost:3000/admin/", http.StatusSeeOther)
        },
  },
  Route{
    "GetProduct",                                     // Name
    "GET",                                            // HTTP method
    "/products/{productId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /products/{productId}")
                db, err := sql.Open("mysql","root:test@tcp(productdb:3306)/products")

                var productId = mux.Vars(r)["productId"]
                var status = http.StatusOK

                rows, err2 := db.Query("select * from product_list l left join product_images i on l.id=i.id where l.id = ?", productId)

                if err2 != nil {
                  log.Fatal(err2)
                  status = http.StatusInternalServerError
                }

                defer rows.Close()

                columns, err := rows.Columns()
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }

                count := len(columns)
                tableData := make([]map[string]interface{}, 0)
                values := make([]interface{}, count)
                valuePtrs := make([]interface{}, count)
                for rows.Next() {
                    for i := 0; i < count; i++ {
                        valuePtrs[i] = &values[i]
                    }
                    rows.Scan(valuePtrs...)
                    entry := make(map[string]interface{})
                    for i, col := range columns {
                        var v interface{}
                        val := values[i]
                        b, ok := val.([]byte)
                        if ok {
                            v = string(b)
                        } else {
                            v = val
                        }
                        entry[col] = v
                    }
                    tableData = append(tableData, entry)
                }
                jsonData, err := json.Marshal(tableData)
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }
                fmt.Println(string(jsonData))

                defer db.Close()

                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.WriteHeader(status)
                w.Write([]byte(jsonData))
        },
  },
  Route{
    "GetProductImages",                                     // Name
    "GET",                                            // HTTP method
    "/products/images/{productId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /products/images/{productId}")
                db, err := sql.Open("mysql","root:test@tcp(productdb:3306)/products")

                var productId = mux.Vars(r)["productId"]
                var status = http.StatusOK

                rows, err2 := db.Query("select * from product_images where id = ?", productId)
                if err2 != nil {
                  log.Fatal(err2)
                  status = http.StatusInternalServerError
                }

                defer rows.Close()

                columns, err := rows.Columns()
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }

                count := len(columns)
                tableData := make([]map[string]interface{}, 0)
                values := make([]interface{}, count)
                valuePtrs := make([]interface{}, count)
                for rows.Next() {
                    for i := 0; i < count; i++ {
                        valuePtrs[i] = &values[i]
                    }
                    rows.Scan(valuePtrs...)
                    entry := make(map[string]interface{})
                    for i, col := range columns {
                        var v interface{}
                        val := values[i]
                        b, ok := val.([]byte)
                        if ok {
                            v = string(b)
                        } else {
                            v = val
                        }
                        entry[col] = v
                    }
                    tableData = append(tableData, entry)
                }
                jsonData, err := json.Marshal(tableData)
                if err != nil {
                  log.Fatal(err)
                  status = http.StatusInternalServerError
                }
                fmt.Println(string(jsonData))

                defer db.Close()

                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.WriteHeader(status)
                w.Write([]byte(jsonData))
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

    db, errdb := sql.Open("mysql","root:test@tcp(productdb:3306)/products")
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


    if strings.Contains(r.Referer(), "http://localhost:3000/admin/edit_product"){
      log.Println("Redirect back to admin")
      http.Redirect(w, r, "http://localhost:3000/admin/edit_product?id="+id, http.StatusSeeOther)
    } else {
      log.Println(r.Referer())
      log.Println("Return JSON")
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("{\"result\":\"OK\"}"))
    }
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

    db, errdb := sql.Open("mysql","root:test@tcp(productdb:3306)/products")
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

      db, errdb := sql.Open("mysql","root:test@tcp(productdb:3306)/products")
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

      log.Println("Return JSON")
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("{\"result\":\"OK\"}"))
    },
  },
}
