package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type commenterRepository struct {
	*dataStoreRepository
}

type commenterEntity struct {
	Name string
}

func NewCommenterRepository(ctx context.Context) usecase.CommenterRepository {
	return &commenterRepository{
		dataStoreRepository: newDataStoreRepository(ctx, "Commenter"),
	}
}

func (r *commenterRepository) NextCommenterId() domain.CommenterId {
	return domain.CommenterId(r.nextID())
}

func (r *commenterRepository) Add(commenter *domain.Commenter) {
	key, entity := r.toDataStoreEntity(commenter)
	r.put(key, entity)
}

func (r *commenterRepository) Delete(id domain.CommenterId) {
	r.delete(r.newKey(int64(id), ""))
}

func (r *commenterRepository) Get(commenterId domain.CommenterId) *domain.Commenter {
	var entity commenterEntity
	key := r.newKey(int64(commenterId), "")
	r.get(key, &entity)
	return r.build(key, &entity)
}

func (r *commenterRepository) toDataStoreEntity(commenter *domain.Commenter) (*datastore.Key, *commenterEntity) {
	key := r.newKey(int64(commenter.CommenterId()), "")
	entity := &commenterEntity{
		Name: commenter.Name(),
	}
	return key, entity
}

func (r *commenterRepository) build(key *datastore.Key, entity *commenterEntity) *domain.Commenter {
	return domain.NewCommenter(domain.CommenterId(key.IntID()), entity.Name)
}
