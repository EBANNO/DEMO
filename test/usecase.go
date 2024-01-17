package test

type EntityRepository interface {
	Create() error
}

type Usecase interface {
	Process()
}

type usecaseImp struct {
	entityRepository EntityRepository
}

func New(entityRepository EntityRepository) Usecase {
	return &usecaseImp{
		entityRepository: entityRepository,
	}
}

func (u *usecaseImp) Process() {
	go u.entityRepository.Create()
}
