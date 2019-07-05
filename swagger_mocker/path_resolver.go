package swagger_mocker

import "log"

const (
	PATH_NORMAL = iota
	PATH_VARIABLE
)

type RouteNode struct {
	Path string
	Api *ApiGroup
	Type int
	Children map[string]*RouteNode
}

const VAR_NAME = "{VAR}"

func (r *RouteNode) AddChild(paths []string, api *ApiGroup) *RouteNode {
	root := paths[0]
	rawRoot := root
	leftPaths := paths[1:]
	if isVariable(root) {
		root = VAR_NAME
	}
	var child *RouteNode
	if router, ok := r.Children[root]; ok {
		if len(leftPaths) == 0 && router.Api != nil {
			log.Print("Overwriting Existing Router.", router.Api)
		}
		child = router
	} else {
		newRoute := NewRouter()
		newRoute.Path = rawRoot
		r.Children[root] = newRoute
		if isVariable(root) {
			newRoute.Type = PATH_VARIABLE
		} else {
			newRoute.Type = PATH_NORMAL
		}
		child = newRoute
	}
	if len(leftPaths) != 0 {
		return child.AddChild(leftPaths, api)
	} else {
		child.Api = api
		return child
	}
}

func (r *RouteNode) FindApi(paths []string) *ApiGroup {
	root := paths[0]
	leftPaths := paths[1:]
	if isVariable(root) {
		root = VAR_NAME
	}
	if router, ok := r.Children[root]; ok {
		if len(leftPaths) == 0 {
			return router.Api
		} else {
			return router.FindApi(leftPaths)
		}
	} else {
		return nil
	}
}

func NewRouter() *RouteNode {
	router := &RouteNode{}
	router.Children = map[string]*RouteNode{}
	return router
}

func isVariable(path string) bool {
	return path[0] == '{' && path[len(path) - 1] == '}'
}
