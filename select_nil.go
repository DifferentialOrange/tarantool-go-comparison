package main

import (
    "fmt"
    "log"
    "context"
    // "github.com/tarantool/go-tarantool"
    "github.com/viciious/go-tarantool"
    // "github.com/FZambia/tarantool"
)

// func TestSelectNilTarantool(uri string, user string, space string) {
//     opts := tarantool.Opts{ User: user }
//     conn, err := tarantool.Connect(uri, opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err)
//     }

//     expr := fmt.Sprintf("return box.space.%s:select{}", space)
//     resp1, err1 := conn.Eval(expr, []interface{}{})

//     log.Println("Error", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }

// func TestSelectNilViciious(uri string, user string, space string) {
//     opts := tarantool.Options{User: user}
//     conn, err := tarantool.Connect(uri, &opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err.Error())
//     }

//     expr := fmt.Sprintf("return box.space.%s:select{}", space)
//     q := &tarantool.Eval{ Expression: expr }
//     resp1 := conn.Exec(context.Background(), q)

//     log.Println("Error", resp1.Error)
//     log.Println("Data", resp1.Data)
// }

// func TestSelectNilFZambia(uri string, user string, space string) {
//     opts := tarantool.Opts{ User: user }
//     conn, errc := tarantool.Connect(uri, opts)
//     if errc != nil {
//         log.Fatalf("Connection refused: %v", errc)
//     }

//     expr := fmt.Sprintf("return box.space.%s:select{}", space)
//     resp1, err1 := conn.Exec(tarantool.Eval(expr, []interface{}{}))

//     log.Println("Error", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }