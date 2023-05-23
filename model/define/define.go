package define

type Fellow struct {
	Name   string
	Year   string
	URL    string
	Region string
}

type NameData struct {
	Text   string     `json:"text"`
	Detail NameDetail `json:"detail"`
}

type NameDetail struct {
	Title    string `json:"title"`
	First    string `json:"first"`
	Middle   string `json:"middle"`
	Last     string `json:"last"`
	Suffix   string `json:"suffix"`
	Nickname string `json:"nickname"`
}
