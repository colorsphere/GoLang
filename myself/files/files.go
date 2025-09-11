package files

import (
	"encoding/json"
	"work/data"
)

type FileName struct {
	fileName string
}

func NewBase(fileName string) *FileName {
	return &FileName{
		fileName: fileName,
	}
}

func ToByte(myvar *data.JsonWork) ([]byte, error) {
	tobyte, err := json.Marshal(myvar)
	if err != nil {
		return nil, err
	} else {
		return tobyte, nil
	}
}

func ToList(myvar *[]byte) (data.JsonWork, error) {
	var tolist data.JsonWork
	err := json.Unmarshal(*myvar, &tolist)
	return tolist, err
}
