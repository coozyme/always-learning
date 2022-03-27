package helper

import (
	"fmt"
	"testing"
)

/*
	Menjalankan unit test
	- go test
	- go test -v								-> Jika kita ingin lihat lebih detail function test apa saja yang sudah di running,
	- go test -v -run TestNamaFunction	->	Jika kita ingin memilih function unit test mana yang ingin di running,
	- go test ./...							-> Menjalankan semua unit test dari top folder

	Menggagalkan unit test menggunakan testing.T , Go-Lang menyediakan function
	terdapat Fail(), FailNow(), Error(), Fatal()
	- Fail()		->	akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai,
						maka unit test tersebut dianggap gagal.
	- FailNow()	-> akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test.
	- Error()	->	function lebih seperti melakukan log (print) error, namun setelah melakukan log error,
						dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal,
						namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
	- Fatal() 	-> mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(),
						sehingga mengakibatkan eksekusi unit test berhenti.

	Error() ==	Fail()
	Fatal() ==	FailNow()
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ary")
	if result != "Hello Ary" {
		//unit test failed
		t.Error("Result is not Ary")
	}
	fmt.Println("TestHelloWorld is DONE")
}
func TestHelloWorldAry(t *testing.T) {
	result := HelloWorld("Ary")
	if result != "Hello Arys" {
		//unit test failed
		t.Error("Result is not Hello Ary")
	}
	fmt.Println("TestHelloWorldAry is DONE")
}
