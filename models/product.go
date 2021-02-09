package models



type Product struct {
	UID         int      `json:"uid"`
	Title       string   `json:"title"`
	Description string   `json:"desc"`
	Reference   int      `json:"ref"`
	Images      []string `json:"imgs"`
}
