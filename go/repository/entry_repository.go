package repository

type EntryRepository interface {
}

type EntryRepositoryImpl struct {
}

func NewEntryRepository() EntryRepository {
	return &EntryRepositoryImpl{}
}
