package main

// tarantool
// box.cfg{ listen = 3401 }
// box.schema.user.grant('guest', 'read,write,execute,create,drop,alter', 'universe')

// -- decimal test
//
// box.schema.space.create('testdecimal')
// box.space.testdecimal:create_index('pk', { type = 'tree', parts = {{ field = 1, type = 'decimal' }} })
//
// decimal = require('decimal')
//
// a = decimal.new('1234.5678')
//
// box.space.testdecimal:insert({a})

// func TestDecimalTarantool(uri string, user string) {
//     TestSelectNilTarantool(uri, user, "testdecimal")
// }

// func TestDecimalViciious(uri string, user string) {
//     TestSelectNilViciious(uri, user, "testdecimal")
// }

// func TestDecimalFZambia(uri string, user string) {
//     TestSelectNilFZambia(uri, user, "testdecimal")
// }
