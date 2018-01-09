package main

import "testing"

func TestKpChPasswd(t *testing.T) {
	kp, err := OpenKeepassDb("./test.kdbx")
	testFailIfErrNotNil(err, t)

	defer kp.Close()

	err = kp.UnlockWithPassword("123456")
	testFailIfErrNotNil(err, t)

	err = kp.ChPassword("another-password")
	testFailIfErrNotNil(err, t)

	err = kp.SaveTo("./test2.kdbx")
	testFailIfErrNotNil(err, t)

}

func testFailIfErrNotNil(err interface{}, t *testing.T) {
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}
