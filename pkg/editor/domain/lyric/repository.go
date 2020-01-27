package lyric

type Repository interface {
	Save(*Lyric) error
	Update(lyric *Lyric, index int) error
	ByIndex(ID int) (*Lyric, error)
}
