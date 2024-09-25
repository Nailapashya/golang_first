package model

type UrusanModel struct {
	IdUrusan   int    `json:"id_urusan" xml:"id_urusan" example:"12"`
	Tahun      int    `json:"tahun" xml:"tahun" example:"2022"`
	IdDaerah   int    `json:"id_daerah" xml:"id_daerah" example:"371"`
	KodeUrusan string `json:"kode_urusan" xml:"kode_urusan" example:"2"`
	NamaUrusan string `json:"nama_urusan" xml:"nama_urusan" example:"URUSAN PEMERINTAHAN WAJIB YANG TIDAK BERKAITAN DENGAN PELAYANAN DASAR"`
	IdUnik     string `json:"id_unik" xml:"id_unik" example:"0076b9dc-02eb-4d05-990a-44b16a9ff76a"` // ID unik urusan
	IsLocked   int    `json:"is_locked" xml:"is_locked" example:"0"`                                // Urusan ini dikunci atau tidak
}
