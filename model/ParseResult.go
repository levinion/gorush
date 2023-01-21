package model

type ParseResult struct{
	Filename string
	Content string
	MetaData map[string]any
}