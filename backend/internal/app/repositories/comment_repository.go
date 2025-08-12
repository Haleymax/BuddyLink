type CommentRepository struct {
	CreateTable() error
	Insert(comment models.Comment) error
	Delete(comment models.Comment) error
	Update(oldData, newData models.Comment) error
	FindById(id uint64) (models.Comment, error)
	FindByCardId(cardId uint64) ([]models.Comment, error)
	FindByUserId(userId uint64) ([]models.Comment, error)
	FindByParentId(parentId uint64) ([]models.Comment, error)
	FindAll() ([]models.Comment, error)
	Exists(comment models.Comment) (bool, error)
}
