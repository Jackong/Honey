/**
 * User: Jackong
 * Date: 13-12-3
 * Time: 下午11:02
 */
package model

import (
	"menteslibres.net/gosexy/db"
	. "github.com/Jackong/Honey/global"
)

var (
	collections map[string]db.Collection
)

func init() {
	collections = make(map[string]db.Collection)
}

func Collection(name string) db.Collection{
	if collection, ok := collections[name]; ok {
		return collection
	}
	collection, err := DB.Collection(name)
	if err != nil {
		Log.Warnf("%v|Get collection failed:%v", name, err)
	}
	collections[name] = collection
	return collection
}
