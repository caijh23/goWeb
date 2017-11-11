package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)


func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		IndentJSON : true,
	})) //注入中间件（渲染JSON和HTML模板的处理器中间件

	//普通的GET方式路由
	m.Get("/", func() string {
		return "hello world!"
	})
	//实现实例中的hello id
	m.Get("/hello/:name", func(params martini.Params, r render.Render){
		r.JSON(200, map[string]interface{}{"Test": "Hello " + params["name"]})
	})

	m.RunOnAddr(":8888") //运行程序监听端口
}
