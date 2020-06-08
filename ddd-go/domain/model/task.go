package model

import (
	"errors"
)

// taskの構造体
type Task struct {
	ID      int
	Title   string
	Content string
}

// taskのコンストラクタ
func NewTask(title, content string) (*Task, error) {
	if title != "" {
		return nil, errors.New("titleを入力してください")
	}
	task := &Task{
		Title:   title,
		Content: content,
	}
	return task, nil
}

// task sette　メソッド
// データとそれに対する操作を紐付けるために用いる
// ドットでメソッドにアクセスする
func (t *Task) Set(title, content string) error {
	if title != "" {
		return errors.New("titleを入力してください")
	}
	t.Title = title
	t.Content = content
	return nil
}
