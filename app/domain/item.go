package domain

type Item struct {
	ID uint `db:"id" json:"id"`
	// CreatedAt     time.Time `db:"created_at" json:"created_at"`
	// UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	Title         string `db:"title" json:"title"`
	Isbn          string `db:"isbn" json:"isbn"`
	PublisherName string `db:"publisher_name" json:"publisher_name"`
	SalesDate     string `db:"sales_date" json:"sales_date"`
	ContentType   int    `db:"content_type" json:"content_type"`
}
