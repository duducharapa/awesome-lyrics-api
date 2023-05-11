package models

type Lyric struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
