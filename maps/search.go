package maps

type Dictionary map[string]string

type DictionaryErr string

var (
	ErrNotFound     = DictionaryErr("not found")
	ErrWordExists   = DictionaryErr("world exist")
	ErrWordNotExist = DictionaryErr("world not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(text string) (string, error) {
	result, ok := d[text]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(world string, definition string) error {
	_, err := d.Search(world)

	switch err {
	case ErrNotFound:
		d[world] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
