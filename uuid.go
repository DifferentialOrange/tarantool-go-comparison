package main

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
//     TestSelectNilTarantool(uri, user, "testuuid")
// }

// func TestUUIDViciious(uri string, user string) {
//     TestSelectNilViciious(uri, user, "testuuid")
// }

// func TestUUIDFZambia(uri string, user string) {
//     TestSelectNilFZambia(uri, user, "testuuid")
// }
