package services

import "github.com/tekotechkotech/mahasiswaservice-go1/pkg/dto"

type Services interface {
	SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error
	UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error
	SaveAlamatId(req *dto.AlamatIdReqDTO) error
	ShowAllMahasiswaAlamat() ([]*dto.MahasiswaAlamatResDTO, error)
}
