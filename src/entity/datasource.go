package entity

type DataSource struct {
	Provinsi  map[string]string `json:"provinsi"`
	Kabkot    map[string]string `json:"kabkot"`
	Kecamatan map[string]string `json:"kecamatan"`
}
