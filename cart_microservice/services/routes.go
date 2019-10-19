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
    "GetAllCarts",                                     // Name
    "GET",                                            // HTTP method
    "/cart/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling /cart/")
                db, err := sql.Open("mysql","root:test@tcp(cartdb:3306)/cart")

                rows,err2 := db.Query("select * from cart_list")

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
    "GetUserCarts",                                     // Name
    "GET",                                            // HTTP method
    "/cart/{userId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
                log.Println("Calling User /cart/")
                var id = mux.Vars(r)["userId"]

                db, err := sql.Open("mysql","root:test@tcp(cartdb:3306)/cart")

                rows,err2 := db.Query("select * from cart_list where user_id=?",id)

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
    "AddToCarts",                                     // Name
    "POST",                                            // HTTP method
    "/cart/add/",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
            log.Println("Adding to /cart/add")

            err := r.ParseForm()
          	if err != nil {
          		panic(err)
          	}
          	v := r.Form

            userId:=r.Form.Get("user_id")
            productId:=r.Form.Get("product_id")
            qty:=r.Form.Get("qty")

            log.Println(v)

            db, errdb := sql.Open("mysql","root:test@tcp(cartdb:3306)/cart")
            if errdb != nil {
          		panic(errdb)
          	}

            stmt, err2 := db.Prepare("INSERT cart_list SET user_id=?, product_id=?, qty=?")
            if err2 != nil {
          		panic(err2)
          	}

            res, err3 := stmt.Exec(userId,productId,qty)
            if err3 != nil {
          		panic(err3)
          	}

            id, err4 := res.LastInsertId()
            if err4 != nil {
          		panic(err4)
          	}

            log.Println(id)

            log.Println("Redirect back to shop")
            http.Redirect(w, r, r.Referer()+"?added=t", http.StatusSeeOther)
        },
  },
  Route{
    "ClearCart",                                     // Name
    "GET",                                            // HTTP method
    "/cart/clear/{userId}",                          // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
            log.Println("Clearing /cart for User")

            var id = mux.Vars(r)["userId"]

            db, err := sql.Open("mysql","root:test@tcp(cartdb:3306)/cart")
            if err != nil {
          		panic(err)
          	}
            rows,err2 := db.Query("delete from cart_list where user_id=?",id)
            if err2 != nil {
          		panic(err2)
          	}

            log.Println(rows)
            log.Println(err2)

            defer db.Close()

            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("{\"result\":\"OK\"}"))
        },
  },
  Route{
    "DeleteItem",                                     // Name
    "GET",                                            // HTTP method
    "/cart/delete/{userId}/{productId}",                // Route pattern
    func(w http.ResponseWriter, r *http.Request) {
            log.Println("Delete Item from /cart for User")

            var id = mux.Vars(r)["userId"]
            var productId = mux.Vars(r)["productId"]

            db, err := sql.Open("mysql","root:test@tcp(cartdb:3306)/cart")
            if err != nil {
          		panic(err)
          	}
            rows,err2 := db.Query("delete from cart_list where user_id=? AND product_id=?",id,productId)
            if err2 != nil {
          		panic(err2)
          	}

            log.Println(rows)
            log.Println(err2)

            defer db.Close()

            log.Println("Redirect back to shop")
            http.Redirect(w, r, r.Referer()+"?added=t", http.StatusSeeOther)
        },
  },
}
