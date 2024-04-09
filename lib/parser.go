package lib

import (
	"encoding/json"
)

type Store struct {
	Mapper map[string]string
	Text   []string
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func (store *Store) Parser(fileBytes []byte) error {
	err := json.Unmarshal(fileBytes, &store.Mapper)
	if err != nil {
		return err
	}

	store.Mapper = reverseMap(store.Mapper)

	for language := range store.Mapper {
		store.Text = append(store.Text, language)
	}

	return nil
}
