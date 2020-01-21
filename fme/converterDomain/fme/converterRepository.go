package fme

type Repository interface {
	Save(fme *Fme) error
	Get() (*Fme, error)
}
