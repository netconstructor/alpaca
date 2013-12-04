package langs

import (
	"bitbucket.org/pkg/inflect"
)

func WriteNode(data *Data) {
	MakeLibraryDir("node")
	RunTemplate := ChooseTemplate("node")

	name := data.Pkg["name"].(string)

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("package.json", "package.json", data)

	MakeDir("lib")

	RunTemplate("lib/index.js", "index.js", data)

	MakeDir(inflect.CamelizeDownFirst(name))

	MakeDir("error")
	RunTemplate("lib/error/index.js", "index.js", data)
	RunTemplate("lib/error/client.js", "client.js", data)
	MoveDir("..")

	MakeDir("client")
	RunTemplate("lib/client/index.js", "index.js", data)
	RunTemplate("lib/client/auth.js", "auth.js", data)
	RunTemplate("lib/client/error.js", "error.js", data)
	RunTemplate("lib/client/request.js", "request.js", data)
	RunTemplate("lib/client/response.js", "response.js", data)
	MoveDir("..")

	MakeDir("api")

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/api.js", inflect.CamelizeDownFirst(k)+".js", data)
		delete(data.Api, "active")
	}
}

func FunctionsNode(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})

	args["node"] = ArgsFunctionMaker("", ", ")
	path["node"] = PathFunctionMaker("\" + this.", " + \"")
}