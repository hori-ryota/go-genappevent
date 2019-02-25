package domain

type ShopID string

//genconstructor
type Shop struct {
	id     ShopID `required:"" getter:""`
	opened bool   `required:"" getter:""`
}

func (m *Shop) Open() Error {
	m.opened = true
	return nil
}

func (m *Shop) Close() Error {
	m.opened = false
	return nil
}
