/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午9:27
 */
package net

var (
	mapper map[string]Module
)

func init() {
	mapper = make(map[string]Module)
}

func Attach(name string, module Module) bool {
	if _, ok := mapper[name]; ok {
		return false
	}
	mapper[name] = module
	return true
}

func GetModule(name string) Module {
	return mapper[name]
}
