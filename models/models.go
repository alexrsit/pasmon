package models

type types struct {
	Types []string `json:"types"`
}

type SamlSpMeta struct {
	ID              string `json:"id"`
	Content         string `json:"content"`
	Created         string `json:"created"`
	ContentType     string `json:"content_type"`
	ContentEncoding string `json:"content_encoding"`
}
