package repository

import (
	"awesomeProject/util"
	"sync"
	"time"
)

type Post struct {
	Id         int64     `gorm:"column:id"`
	ParentId   int64     `gorm:"parent_id"`
	UserId     int64     `gorm:"column:user_id"`
	Content    string    `gorm:"column:content"`
	DiggCount  int32     `gorm:"column:digg_count"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Post) TableName() string {
	return "post"
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostByParentId(parentId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("parent_id = ?", parentId).Find(&posts).Error
	if err != nil {
		util.Logger.Error("find posts by parent_id err:" + err.Error())
		return nil, err
	}
	return posts, nil
}

func (*PostDao) CreatePost(post *Post) error {
	if err := db.Create(post).Error; err != nil {
		util.Logger.Error("insert post err:" + err.Error())
		return err
	}
	return nil
}
