package main

import (
    "log"
    "encoding/binary"
    "github.com/tarantool/go-tarantool"
)

func main() {
    // tarantool
    // box.cfg{ listen = 3401 }
    // box.schema.user.grant('guest', 'read,write,execute,create,drop,alter', 'universe')
    //
    // -- varbinary test
    //
    // box.schema.space.create('testvarbin')
    // box.space.testvarbin:create_index('pk', { type = 'tree', parts = {{ field = 1, type = 'varbinary' }} })
    //
    // buffer = require('buffer')
    // ffi = require('ffi')
    //
    // function encode_bin(bytes)
    //     local tmpbuf = buffer.ibuf()
    //     local p = tmpbuf:alloc(3 + #bytes)
    //     p[0] = 0x91
    //     p[1] = 0xC4
    //     p[2] = #bytes
    //     for i, c in pairs(bytes) do
    //         p[i + 3 - 1] = c
    //     end
    //     return tmpbuf
    // end
    //
    // function bintuple_insert(space, bytes)
    //     local tmpbuf = encode_bin(bytes)
    //     ffi.cdef[[int box_insert(uint32_t space_id, const char *tuple, const char *tuple_end, box_tuple_t **result);]]
    //     ffi.C.box_insert(space.id, tmpbuf.rpos, tmpbuf.wpos, nil)
    // end
    //
    // bintuple_insert(box.space.testvarbin, {0xDE, 0xAD, 0xBE, 0xAF})

    // github.com/tarantool/go-tarantool
    opts := tarantool.Opts{ User: "guest" }
    conn, err := tarantool.Connect("127.0.0.1:3401", opts)

    if err != nil {
        log.Fatalf("Connection refused:", err)
    }

    resp, err := conn.Ping()
    log.Println(resp.Code)
    log.Println(resp.Data)
    log.Println(err)

    resp1, err1 := conn.Eval("return box.space.testvarbin:select{}", []interface{}{})
    
    log.Println("Error", err1)
    log.Println("Code", resp1.Code)
    log.Println("Data", resp1.Data)

    buf := make([]byte, 4)
    binary.BigEndian.PutUint16(buf[0:], 0xa20c)
    binary.BigEndian.PutUint16(buf[2:], 0x04af)

    resp2, err2 := conn.Insert("testvarbin", []interface{}{buf})
    
    log.Println("Error", err2)
    log.Println("Code", resp2.Code)
    log.Println("Data", resp2.Data)
}
