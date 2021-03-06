package models

type MahasiswaModels struct {
	ID   int64  `db:"id"`
	Name string `db:"nama"`
	Nim  string `db:"nim"`
}

type MahasiswaAlamatModels struct {
	ID           int64  `db:"id"`
	Jalan        string `db:"jalan"`
	NoRumah      string `db:"no_rumah"`
	IDMahasiswas int64  `db:"id_mahasiswas"`
}

type ShowMahasiswaAlamatModels struct {
	ID      int64  `db:"id"`
	Name    string `db:"nama"`
	Nim     string `db:"nim"`
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
}
