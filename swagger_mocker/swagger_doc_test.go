package swagger_mocker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRouter(t *testing.T) {
	router := NewRouter()
	router.AddChild([]string{"p1", "p2", "p3"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World",
		},
	})
	assert.Equal(t, router.Children["p1"].Path, "p1")
	assert.Equal(t, router.Children["p1"].Children["p2"].Path, "p2")
	assert.Equal(t, router.Children["p1"].Children["p2"].Children["p3"].Path, "p3")
	assert.Equal(t, router.Children["p1"].Children["p2"].Children["p3"].Api.Post.Description, "Hello World")
}

func TestInsertVariableRouter(t *testing.T) {
	router := NewRouter()
	router.AddChild([]string{"p1", "{p2}", "p3"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World",
		},
	})
	assert.Equal(t, router.Children["p1"].Type, PATH_NORMAL)
	assert.Equal(t, router.Children["p1"].Children["{VAR}"].Type, PATH_VARIABLE)
}

func TestAddApiToExistingRouter(t *testing.T) {
	router := NewRouter()
	router.AddChild([]string{"p1", "{p2}", "p3", "p4"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World",
		},
	})
	router.AddChild([]string{"p1", "{p2}", "p3"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World P3",
		},
	})
	assert.Equal(t, "Hello World P3", router.Children["p1"].Children["{VAR}"].Children["p3"].Api.Post.Description)
}

func TestAddApiToExistingApi(t *testing.T)  {
	router := NewRouter()
	router.AddChild([]string{"p1", "{p2}", "p3"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World",
		},
	})
	assert.Equal(
		t,
		"Hello World",
		router.Children["p1"].Children["{VAR}"].Children["p3"].Api.Post.Description,
	)
	router.AddChild([]string{"p1", "{p2}", "p3"}, &ApiGroup{
		Post: &Api{
			Description: "Hello World 2",
		},
	})
	assert.Equal(
		t,
		"Hello World 2",
		router.Children["p1"].Children["{VAR}"].Children["p3"].Api.Post.Description,
	)
}

func TestFindingRouterByPaths(t *testing.T) {
	router := NewRouter()
	router.AddChild([]string{ "p1" }, &ApiGroup{ Post: &Api{ Description: "/p1" }, })
	router.AddChild([]string{ "p1", "p2" }, &ApiGroup{ Post: &Api{ Description: "/p1/p2" }, })
	router.AddChild([]string{ "p1", "p3" }, &ApiGroup{ Post: &Api{ Description: "/p1/p3" }, })
	router.AddChild([]string{ "p1", "p2", "p4", "p5" }, &ApiGroup{ Post: &Api{ Description: "/p1/p2/p4/p5" }, })
	router.AddChild([]string{ "p1", "p2", "p4", "p6" }, &ApiGroup{ Post: &Api{ Description: "/p1/p2/p4/p6" }, })

	assert.Equal(t, router.FindApi([]string{ "p1", "p2", "p4", "p6" }).Post.Description, "/p1/p2/p4/p6")
	assert.Equal(t, router.FindApi([]string{ "p1", "p3" }).Post.Description, "/p1/p3")
	assert.Equal(t, router.FindApi([]string{ "p1" }).Post.Description, "/p1")
}
