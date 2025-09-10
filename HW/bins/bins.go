package bins

type Tojson []struct {
	Name    string `json:"name"`
	CerName string `json:"cername"`
}

// func (tojson Tojson) WriteFile() ([]byte, error) { // методы прописываются сразу после объявления структуры
// 	jsonData, err := json.Marshal(tojson)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if files.Exists("test.test") {
// 		fmt.Println("Файла не существует")
// 	}

// 	return file, nil
// }
