package main

// import (
//     "log"
//     // "context"
//     // "github.com/tarantool/go-tarantool"
//     // "github.com/viciious/go-tarantool"
//     "github.com/FZambia/tarantool"
// )

// func TestBoxErrorTarantool(uri string, user string) {
//     opts := tarantool.Opts{ User: user }
//     conn, err := tarantool.Connect(uri, opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err)
//     }

//     resp1, err1 := conn.Eval("box.error(box.error.new({code = 1000, reason = 'Reason'}))", []interface{}{})

//     log.Println("Error", err1)
//     log.Printf("%T", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }

// func TestBoxErrorViciious(uri string, user string) {
//     opts := tarantool.Options{User: user}
//     conn, err := tarantool.Connect(uri, &opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err.Error())
//     }

//     expr := "box.error(box.error.new({code = 1000, reason = 'Reason'}))"
//     q := &tarantool.Eval{ Expression: expr }
//     resp1 := conn.Exec(context.Background(), q)

//     log.Println("Error", resp1.Error)
//     log.Printf("%T", resp1.Error)
//     log.Println("Data", resp1.Data)
// }

// func TestBoxErrorFZambia(uri string, user string) {
//     opts := tarantool.Opts{ User: user }
//     conn, errc := tarantool.Connect(uri, opts)
//     if errc != nil {
//         log.Fatalf("Connection refused: %v", errc)
//     }

//     expr := "box.error(box.error.new({code = 1000, reason = 'Reason'}))"
//     resp1, err1 := conn.Exec(tarantool.Eval(expr, []interface{}{}))

//     log.Println("Error", err1)
//     log.Printf("%T", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }