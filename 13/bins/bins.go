package bins

type BinsDb struct {
	url string
}

func NewBinsDb(url string) *BinsDb {
	return &BinsDb{
		url: url,
	}
}

func (db *BinsDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *BinsDb) Write(content []byte) {

}
