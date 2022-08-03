package main

import (
	"encoding/json"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here
type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}
type ruangTamu struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Ruang struct {
	ruangTamu []*ruangTamu
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	encode, err := json.Marshal(r.ruangTamu)
	if err != nil {
		return "JSON Marshal error: "
	}

	return string(encode)
}

func NewRuang(r Ruang) Ruang {
	return r
}
