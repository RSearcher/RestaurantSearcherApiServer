package models

type Review struct {
	Id int `json:"id"`
	Body string `json:"body"`
}

type ParsedText struct {
	Candidates []string `json:"candidates"`
	Morphs [][][]string `json:"morphs"`
}