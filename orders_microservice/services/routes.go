package services
import (
  "net/http"
  "database/sql"
  _ "./mysql"
  "log"
  "fmt"
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
    "GetAllOrders",                                     // Name
    "GET",                                            // HTTP method
    "/orders/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /orders/")
                db, err := sql.Open("mysql","root:test@tcp(ordersdb:3306)/orders")

                rows,err2 := db.Query("select * from order_list")

                log.Println(rows)
                log.Println(err2)

                var status = http.StatusBadRequest

                switch {
                case err2 == sql.ErrNoRows:
                  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                  w.WriteHeader(http.StatusOK)
                  w.Write([]byte("{\"result\":\"OK\"}"))
                  return
                case err2 != nil:
                        log.Fatal(err)
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.WriteHeader(http.StatusOK)
                        w.Write([]byte("{\"result\":\"OK\"}"))
                        return
                default:
                        status = http.StatusOK

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
                }
        },
  },
  Route{
    "GetUserOrders",                                     // Name
    "GET",                                            // HTTP method
    "/orders/{userId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling User /orders/")
                var id = mux.Vars(r)["userId"]

                db, err := sql.Open("mysql","root:test@tcp(ordersdb:3306)/orders")

                rows,err2 := db.Query("select * from order_list where user_id=?",id)

                log.Println(rows)
                log.Println(err2)

                var status = http.StatusBadRequest

                switch {
                case err2 == sql.ErrNoRows:
                  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                  w.WriteHeader(http.StatusOK)
                  w.Write([]byte("{\"result\":\"OK\"}"))
                  return
                case err2 != nil:
                        log.Fatal(err)
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.WriteHeader(http.StatusOK)
                        w.Write([]byte("{\"result\":\"OK\"}"))
                        return
                default:
                        status = http.StatusOK

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
                }
        },
  },
  Route{
    "AddToOrders",                                     // Name
    "POST",                                            // HTTP method
    "/orders/add/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
            log.Println("Adding to /orders/add")

            err := r.ParseForm()
          	if err != nil {
          		panic(err)
          	}
          	v := r.Form

            log.Println(v)

            db, errdb := sql.Open("mysql","root:test@tcp(ordersdb:3306)/orders")
            if errdb != nil {
          		panic(errdb)
          	}

            stmt, err2 := db.Prepare("INSERT order_list SET user_id=?, total=?, qty=?")
            if err2 != nil {
          		panic(err2)
          	}

            res, err3 := stmt.Exec(r.Form.Get("user_id"),r.Form.Get("total"),r.Form.Get("qty"))
            if err3 != nil {
          		panic(err3)
          	}

            id, err4 := res.LastInsertId()
            if err4 != nil {
          		panic(err4)
          	}

            log.Println(id)

            log.Println("Redirect back to shop")

            var next = "http://cartservice:4201/cart/clear/"+r.Form.Get("user_id")

            resp, err := http.Get(next)

            log.Println(resp)

            if err != nil {
          		panic(err)
          	}

            http.Redirect(w, r, "/shop/thankyou", http.StatusSeeOther)
        },
  },
  Route{
    "DeleteOrder",                                     // Name
    "GET",                                            // HTTP method
    "/orders/delete/{orderId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Deleting /orders/")
                var id = mux.Vars(r)["orderId"]

                db, err := sql.Open("mysql","root:test@tcp(ordersdb:3306)/orders")

                rows,err2 := db.Query("delete from order_list where order_id=?",id)

                log.Println(rows)
                log.Println(err2)

                var status = http.StatusBadRequest

                switch {
                case err2 != nil:
                        log.Fatal(err)
                        return
                default:
                        status = http.StatusOK
                        defer rows.Close()
                        defer db.Close()
                }

                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(status)
                w.Write([]byte("{\"result\":\"OK\"}"))
        },
  },
}
