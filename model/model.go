/**
 * User: Jackong
 * Date: 13-12-3
 * Time: 下午11:02
 */
package model

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/Jackong/db"
	. "github.com/Jackong/Honey/global"
)

var (
	models map[string]*Model
)

type Model struct {
	db.Collection
}

func init() {
	models = make(map[string]*Model)
}

func NewModel(name string) *Model{
	if model, ok := models[name]; ok {
		return model
	}

	collection, err := DB.Collection(name)
	if err != nil {
		Log.Alert("%v|New model:%v", name, err)
		ShutDown()
	}

	model := &Model{Collection: collection}
	models[name] = model
	return model
}

func (this *Model) Get(id interface {}, msg proto.Message) bool {
	item, err := this.Find(db.Cond{
		"id": id,
	})
	if err != nil {
		Log.Errorf("%v|%v is not exists:%v", this.Name(), id, err)
		return false
	}
	err = proto.Unmarshal([]byte(item["data"].(string)), msg)
	if err != nil {
		Log.Errorf("%v|Unmarshal meta:%v", this.Name(), err)
		return false
	}
	return true
}

func (this *Model) Set(id interface {}, msg proto.Message) bool {
	buf, err := proto.Marshal(msg)
	if err != nil {
		Log.Errorf("%v|Marshal meta:%v", this.Name(), err)
		return false
	}
	_, err = this.Append(db.Item{
		"id": id,
		"data": buf,
	})
	if err != nil {
		Log.Errorf("%v|Append meta:%v", this.Name(), err)
		return false
	}
	return true
}

func (this *Model) Exist(id interface {}) bool{
	_, err := this.Find(db.Cond{
		"id": id,
	})
	return err == nil
}
