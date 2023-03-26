package repositories

import (
	"github.com/zekhoi/golang-todo/app/entities"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	FindAll() []entities.Activity
	FindByID(id int) entities.Activity
	Save(activity entities.Activity) entities.Activity
	Update(activity entities.Activity) entities.Activity
	Delete(id int) error
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db: db}
}

func (r *activityRepository) FindAll() []entities.Activity {
	var activities []entities.Activity
	r.db.Find(&activities)

	return activities
}

func (r *activityRepository) FindByID(id int) entities.Activity {
	var activity entities.Activity
	r.db.Where("id = ?", id).First(&activity)

	if activity.ID < 1 {
		return entities.Activity{}
	}

	return activity
}

func (r *activityRepository) Save(activity entities.Activity) entities.Activity {
	var newActivity entities.Activity

	newActivity = entities.Activity{
		Title: activity.Title,
	}
	r.db.Create(&newActivity)

	return newActivity
}

func (r *activityRepository) Update(activity entities.Activity) entities.Activity {
	var updatedActivity entities.Activity

	r.db.Model(&updatedActivity).Where("id = ?", activity.ID).Updates(entities.Activity{
		Title: activity.Title,
	})

	return updatedActivity
}

func (r *activityRepository) Delete(id int) error {
	var activity entities.Activity
	r.db.Where("id = ?", id).Delete(&activity)
	return nil
}
