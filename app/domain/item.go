package domain

type Item struct {
	ID uint `db:"id" json:"id"`
	// CreatedAt     time.Time `db:"created_at" json:"created_at"`
	// UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	Author        string `db:"author" json:"author"`
	Title         string `db:"title" json:"title"`
	Isbn          string `db:"isbn" json:"isbn"`
	PublisherName string `db:"publisher_name" json:"publisher_name"`
	SalesDate     string `db:"sales_date" json:"sales_date"`
	ContentType   int    `db:"content_type" json:"content_type"`
}

// カテゴリ
type ContentType int

const (
	unknown             ContentType = iota
	ContentTypeBook                 // 本
	ContentTypeCD                   // CD/DVD/BD
	ContentTypeSoftware             // ソフトウェア
	ContentTypeGame                 // ゲーム
	ContentTypeOthers               // その他
)

func ContentTypeString(ct int) string {
	switch ContentType(ct) {
	case ContentTypeBook:
		return "本"
	case ContentTypeCD:
		return "CD/DVD/BD"
	case ContentTypeSoftware:
		return "ソフトウェア"
	case ContentTypeGame:
		return "ゲーム"
	case ContentTypeOthers:
		return "その他"
	default:
		return "unknown"
	}
}
