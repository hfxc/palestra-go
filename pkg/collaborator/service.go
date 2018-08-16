package collaborator

import (
	"strings"
	"time"

	"palestra-go/pkg/entity"

	bson "gopkg.in/mgo.v2/bson"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create an bookmark
func (s *Service) Create(b *entity.Collaborator) (bson.ObjectId, error) {
	b.ID = bson.NewObjectId()
	b.CreatedAt = time.Now()
	return s.repo.Create(b)
}

//Find a bookmark
func (s *Service) Find(id bson.ObjectId) (*entity.Collaborator, error) {
	return s.repo.Find(id)
}

//Search bookmarks
func (s *Service) Search(query string) ([]*entity.Collaborator, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll bookmarks
func (s *Service) FindAll() ([]*entity.Collaborator, error) {
	return s.repo.FindAll()
}

//Delete a bookmark
func (s *Service) Delete(id bson.ObjectId) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
