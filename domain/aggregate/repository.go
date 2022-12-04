package aggregate

type Repository interface {
	Save(shape Shape) error
}
