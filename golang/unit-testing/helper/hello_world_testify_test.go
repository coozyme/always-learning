package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	assert	==	Error()	==	Fail()
	require	== Fatal() ==	FailNow()
*/

/*
	Before After Test can use TestMain.
	Functionnya harus bernama TestMain dan mempunyai param dari pointer testing.M .
	Function TestMain ini hanya dilakukan dalam satu package itu saja.

*/
func TestMain(t *testing.M) {
	fmt.Println("BEFORE RUN TES MAIN")
	t.Run()
	fmt.Println("AFTER RUN TES MAIN")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ary")
	assert.Equal(t, "Hello Ary", result, "Result mush be Hello Ary")
	fmt.Println("TestHelloWorldAssert with assert is DONE")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ary")
	require.Equal(t, "Hello Ary", result, "Result mush be Hello Ary")
	fmt.Println("TestHelloWorldRequire with require is DONE")
}
