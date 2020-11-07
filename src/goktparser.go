package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/barokurniawan/goktparser/src/entity"
)

type GoKtParser struct {
	SourcePath    string
	RawDataDaerah string
	DataDaerah    *entity.DataSource
}

func (gkp *GoKtParser) ReadResource() (string, error) {
	dat, err := ioutil.ReadFile(gkp.SourcePath)
	if err != nil {
		return "", err
	}

	gkp.RawDataDaerah = string(dat)
	return gkp.RawDataDaerah, nil
}

func (gkp *GoKtParser) ParseResource() *entity.DataSource {
	mDataSource := entity.DataSource{}
	err := json.Unmarshal([]byte(gkp.RawDataDaerah), &mDataSource)
	if err != nil {
		panic(err)
	}

	gkp.DataDaerah = &mDataSource
	return gkp.DataDaerah
}

func (gkp *GoKtParser) GetProvince(Code string) string {
	if val, ok := gkp.DataDaerah.Provinsi[Code]; ok {
		return val
	}

	return ""
}

func (gkp *GoKtParser) GetCity(Code string) string {
	if val, ok := gkp.DataDaerah.Kabkot[Code]; ok {
		return val
	}

	return ""
}

func (gkp *GoKtParser) GetDistrict(Code string) string {
	if val, ok := gkp.DataDaerah.Kecamatan[Code]; ok {
		return val
	}

	return ""
}

func (gkp *GoKtParser) ParseNIK(nik string) *entity.Output {
	r, _ := regexp.Compile(`(\d\d)(\d\d)(\d\d)(\d\d)(\d\d)(\d\d)(\d\d\d\d)`)
	province, _ := regexp.Compile(`\d\d`)
	city, _ := regexp.Compile(`\d\d`)
	district, _ := regexp.Compile(`\d\d`)
	date, _ := regexp.Compile(`\d\d`)
	month, _ := regexp.Compile(`\d\d`)
	year, _ := regexp.Compile(`\d\d`)
	uniqueNum, _ := regexp.Compile(`\d\d\d\d`)

	var output *entity.Output
	matched := r.FindAllStringSubmatch(nik, -1)
	for i := range matched {
		province := province.FindString(matched[i][1])
		city := province + city.FindString(matched[i][2])
		district := city + district.FindString(matched[i][3])
		date := date.FindString(matched[i][4])
		month := month.FindString(matched[i][5])
		year := year.FindString(matched[i][6])
		uniqueNum := uniqueNum.FindString(matched[i][7])

		fmt.Println(province, city, district, date, month, year, uniqueNum)
		output = &entity.Output{
			Province:  gkp.GetProvince(province),
			City:      gkp.GetCity(city),
			District:  gkp.GetDistrict(district),
			BirthDate: fmt.Sprintf("%s-%s-%s", year, month, date),
			UniqueID:  uniqueNum,
		}

	}

	return output
}
