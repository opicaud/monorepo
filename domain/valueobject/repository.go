package valueobject

type Repository interface {
	Save(shape Shape) error
}
