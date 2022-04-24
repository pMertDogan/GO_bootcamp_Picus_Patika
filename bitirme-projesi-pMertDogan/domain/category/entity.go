package category

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"mime/multipart"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Categorys []Category

type Category struct {
	gorm.Model
	CategoryName        string `gorm:"unique" binding:"required,min=3,alphanum"` //make sure category name is unique
	CategoryDescription string `binding:"required,min=3,alphanum"`
	// CategoryParentID    int
	CategoryIsActive bool `gorm:"default:true"`
}

// fromJson Category
func UnmarshalBooks(data []byte) (Categorys, error) {
	var r Categorys
	err := json.Unmarshal(data, &r)
	return r, err
}

//Categorys toJson
func (r *Categorys) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Category toJson
func (r *Category) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func CategoryFromCSV(file *multipart.FileHeader) (Categorys, error) {
	var nCategory Category
	var nCategoryS Categorys
	// Creating a connection via csv files
	// csvFile, err := os.Open(path)
	// defer csvFile.Close()
	csvFile, err := file.Open()
	defer csvFile.Close()

	if err != nil {
		zap.L().Error("error opening csv file", zap.Error(err))
		return nil, errors.New("error opening csv file" + err.Error())
	}
	//creating a *csv.reader
	reader := csv.NewReader(csvFile)
	//reading csv files via reader and method of reader

	csvData, err := reader.ReadAll()
	if err != nil {
		zap.L().Error("error reading csv file", zap.Error(err))
		return nil, errors.New("error reading csv file" + err.Error())
	}
	//in csv file, each comma represent a column, with this knowledge every row has 12 column and each column parse where they are belong.

	for i, each := range csvData {
		// fmt.Println(i, each)
		//Skip csv header row

		if i > 0 {
			//we dont need do this . Gorm handle it
			// id, _ := strconv.Atoi(each[0])
			// nCategory.ID =  uint(id)
			// nCategory.CreatedAt = time.Now()
			// nCategory.UpdatedAt = time.Now()
			// nCategory.DeletedAt = time.Now()
			nCategory.CategoryName = each[0]
			nCategory.CategoryDescription = each[1]
			nCategory.CategoryIsActive, err = strconv.ParseBool(each[2])
			if err != nil {
				zap.L().Error("error parsing bool", zap.Error(err))
				return nil, errors.New("error, please validate  values" + err.Error())
			}
			nCategoryS = append(nCategoryS, nCategory)
		} else if i == 0 {
			//thats make less query :)
			// Headers should be CategoryName,CategoryDescription,CategoryIsActive
			//We dont care header row but this add additional check to avoid false file usage
			if each[0] != "CategoryName" || each[1] != "CategoryDescription" || each[2] != "CategoryIsActive" {
				zap.L().Error("error, please validate  headers. Supported csv headers are 	CategoryName,CategoryDescription,CategoryIsActive")
				return nil, errors.New("error, please validate  headers. Supported csv headers are 	CategoryName,CategoryDescription,CategoryIsActive")
			}
		}

	}

	return nCategoryS, nil
	// return nCategoryS,nil
}
