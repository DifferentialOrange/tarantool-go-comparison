package main

func main() {
    var uri string = "127.0.0.1:3401"
    var user string = "guest"

    // TestVarbinaryTarantool(uri, user)
    // TestVarbinaryViciious(uri, user)
    // TestVarbinaryFZambia(uri, user)

    // TestUUIDTarantool(uri, user)
    TestUUIDViciious(uri, user)
    // TestUUIDFZambia(uri, user)
}
