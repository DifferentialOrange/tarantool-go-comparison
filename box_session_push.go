package main

import (
    "log"
    "context"
    // "github.com/tarantool/go-tarantool"
    "github.com/viciious/go-tarantool"
    // "github.com/FZambia/tarantool"
)
// rawset(_G, 'func', function() box.session.push('msg'); end)

// func TestBoxSessionPushTarantool(uri string, user string) {
//     opts := tarantool.Opts{ User: user }
//     conn, err := tarantool.Connect(uri, opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err)
//     }

//     resp1, err1 := conn.Eval("box.session.push('go')", []interface{}{})

//     log.Println("Error", err1)
//     log.Printf("%T", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)


//     resp2, err2 := conn.Call("func", []interface{}{})

//     log.Println("Error", err2)
//     log.Printf("%T", err2)
//     log.Println("Code", resp2.Code)
//     log.Println("Data", resp2.Data)  
// }

// func TestBoxSessionPushViciious(uri string, user string) {
//     opts := tarantool.Options{User: user}
//     conn, err := tarantool.Connect(uri, &opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err.Error())
//     }

//     expr := "box.session.push(1)"
//     q := &tarantool.Eval{ Expression: expr }
//     resp1 := conn.Exec(context.Background(), q)

//     log.Println("Error", resp1.Error)
//     log.Printf("%T", resp1.Error)
//     log.Println("Data", resp1.Data)

//     q1 := &tarantool.Call{ Name: "func" }
//     resp2 := conn.Exec(context.Background(), q1)

//     log.Println("Error", resp2.Error)
//     log.Printf("%T", resp2.Error)
//     log.Println("Data", resp2.Data)   
// }

func TestBoxSessionPushFZambia(uri string, user string) {
    opts := tarantool.Opts{ User: user }
    conn, errc := tarantool.Connect(uri, opts)
    if errc != nil {
        log.Fatalf("Connection refused: %v", errc)
    }

    expr := "box.session.push('ds')"
    resp1, err1 := conn.Exec(tarantool.Eval(expr, []interface{}{}))

    log.Println("Error", err1)
    log.Printf("%T", err1)
    log.Println("Code", resp1.Code)
    log.Println("Data", resp1.Data)

    resp2, err2 := conn.Exec(tarantool.Call("func", []interface{}{}))

    log.Println("Error", err2)
    log.Printf("%T", err2)
    log.Println("Code", resp2.Code)
    log.Println("Data", resp2.Data)
}