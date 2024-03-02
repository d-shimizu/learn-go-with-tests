package main

type Directory map[string]string

func (d Directory) Search(word string) string {
	return d[word]
}
