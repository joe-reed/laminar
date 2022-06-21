package store

type InMemoryStore []string

func (s *InMemoryStore) Add(item string) error {
	*s = append(*s, item)

	return nil
}

func (s *InMemoryStore) Next() (string, error) {
	if len(*s) == 0 {
		return "", nil
	}

	return (*s)[0], nil
}

func (s *InMemoryStore) Pop() (string, error) {
	item, err := s.Next()
	if err != nil {
		return "", err
	}

	if len(*s) == 0 {
		return "", nil
	}

	*s = (*s)[1:]

	return item, nil
}
