package repository

import "github.com/tekotechkotech/mahasiswaservice-go1/internal/models"

type Repository interface {
	SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error
	UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error
	GetMahasiswaAlamatByID(id int64) ([]*models.GetMahasiswaAlamatsModels, error)
}
