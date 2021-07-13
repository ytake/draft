package record

type DocumentKey string

// Document temporary document
type Document struct {
	Key     DocumentKey            `json:"key,omitempty" description:"取り出しキー名"`
	Data    map[string]interface{} `json:"document" description:"保管内容"`
	Referer string                 `json:"referer" description:"登録元URL"`
	Expire  int32                  `json:"expire,omitempty" description:"投稿データ保存時間(秒)"`
}

type WriteDocument struct {
	Document
}

type ReadDocument struct {
	Document
	CreatedAt string `json:"created_at,omitempty" description:"データ作成日時"`
}

type Documenter interface {
	RetrieveDocument(key DocumentKey) (*ReadDocument, error)
	SaveDocument(document WriteDocument) error
	// RemoveDocument(key DocumentKey) error
}
