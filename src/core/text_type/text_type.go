package text_type

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}
type InputData struct {
	URL  string `json:"site"`
	Tags string `json:"tags"`
}
