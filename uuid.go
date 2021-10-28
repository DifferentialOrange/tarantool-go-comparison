package main

import (
    // "log"
    // "context"
    // "github.com/tarantool/go-tarantool"
    // "github.com/viciious/go-tarantool"
    // "github.com/FZambia/tarantool"
)

// tarantool
// box.cfg{ listen = 3401 }
// box.schema.user.grant('guest', 'read,write,execute,create,drop,alter', 'universe')

// -- uuid test
//
// box.schema.space.create('testuuid')
// box.space.testuuid:create_index('pk', { type = 'tree', parts = {{ field = 1, type = 'uuid' }} })
//
// uuid = require('uuid')
//
// a = uuid.fromstr("c8f0fa1f-da29-438c-a040-393f1126ad39")
//
// box.space.testuuid:insert({a})

// func TestUUIDTarantool(uri string, user string) {
//     opts := tarantool.Opts{ User: user }
//     conn, err := tarantool.Connect(uri, opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err)
//     }

//     resp1, err1 := conn.Eval("return box.space.testuuid:select{}", []interface{}{})

//     log.Println("Error", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }

// func TestUUIDViciious(uri string, user string) {
//     opts := tarantool.Options{User: user}
//     conn, err := tarantool.Connect(uri, &opts)

//     if err != nil {
//         log.Fatalf("Connection refused:", err.Error())
//     }

//     expr := "return box.space.testuuid:select{}"
//     q := &tarantool.Eval{ Expression: expr }
//     resp1 := conn.Exec(context.Background(), q)

//     log.Println("Error", resp1.Error)
//     log.Println("Data", resp1.Data)
// }

// func TestUUIDFZambia(uri string, user string) {
//     opts := tarantool.Opts{ User: user }
//     conn, errc := tarantool.Connect(uri, opts)
//     if errc != nil {
//         log.Fatalf("Connection refused: %v", errc)
//     }

//     resp1, err1 := conn.Exec(tarantool.Eval("return box.space.testuuid:select{}", []interface{}{}))

//     log.Println("Error", err1)
//     log.Println("Code", resp1.Code)
//     log.Println("Data", resp1.Data)
// }
