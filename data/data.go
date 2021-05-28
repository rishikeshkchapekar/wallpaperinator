package data

type Response struct{
	Results []Result `json:"results"`
}

type Result struct{
	Id string `json:"id"`
	Width int `json:"width"`
	Height int `json:"Height"`
	Colour string `json:"color"`
	Description string `json:"description"`
	Urls Urls `json:"urls"`
}

type Urls struct{
	Raw string `json:"raw"`
	Full string `json:"full"`
	Regular string `json:"regular"`
	Small string `json:"small"`
	Thumbnail string `json:"thumb"`
}