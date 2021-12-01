package dto

import (
	"github.com/tekotechkotech/mahasiswaservice-go1/pkg/common/validator"
)

type MahasiswaReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nim     string         `json:"nim" valid:"required" validname:"nim"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

type AlamatReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

func (dto *MahasiswaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type UpadeMahasiswaNamaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	ID   int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *UpadeMahasiswaNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type GetMahasiswaAlamatByIDRespDTO struct {
	ID      int64            `json:"id"`
	Nama    string           `json:"nama"`
	Nim     string           `json:"nim"`
	Alamats []*AlamatRespDTO `json:"alamat"`
}

type AlamatRespDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

// tugas 1
// http://localhost:9000/api/v1/latihan/alamat
// Method Post
// Body :
// {
// 		"jalan": "foo",
// 		"no_rumah": "234",
// 		"mahasiswa_id : 1
// }

type MahasiswaAlamatReqDTO struct {
	Jalan         string `json:"jalan" valid:"required" validname:"jalan"`
	NoRumah       string `json:"no_rumah" valid:"required" validname:"no_rumah"`
	id_mahasiswas int64  `json:"id_mahasiswas" valid:"required" validname:"id_mahasiswas"`
}

func (dto *MahasiswaAlamatReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

// ///////////
// {
// 		succes : true,
// 		message : get data succes,
// 		data : {
// 			id : 1,
// 			nama : foo,
// 			nim : bar,
// 			alamat : [
// 				{
// 					jalan : foo,
// 					no_rumah : bar,
// 				}
// 			]
// 		}
// 	}
