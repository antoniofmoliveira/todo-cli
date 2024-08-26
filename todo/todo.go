package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
	DueDate  string
	Created  string
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}

type ByPri []Item

func (s ByPri) Len() int { return len(s) }

func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}
