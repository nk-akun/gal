### Practice for Web Framework Router based golang.

### V1要点：
* Server实现了请求入口ServeHTTP, http.ListenAndServe只需要一个实现了ServeHTTP的interface即可
* 实现了Context，封装了http.ResponseWriter, *http.Request, Method, Path等内容,并据Context实现了若干常用方法
* Router中使用map存储路由表, key为"method-path",例如"GET-/Hello"

### V2要点:
* 实现了前缀树路由,支持类似 /:name或/*file形式
    * GET、POST等方法分别维护一棵树,然后做把path按照"/"进行split，再做trie树
* Router中增加roots来存储路由前缀树的根
* 支持分组路由，构建GroupRouter,并将server封装进GroupRouter，同时把GroupRouter封装进server，这样互相封装使之二者的方法能够相互使用.
在设计上，Server是最上级结构，GroupRouter是其一部分
* 具体实现见代码注释，使用demo见main中testV2

### V3要点:
* 实现了基于分组的中间件应用，可以指定任意一个group创建中间件函数
* 实现的基本思想是将中间件函数存于group，然后根据用户请求URL前缀匹配择出对应中间件函数，并存于Context的handlers中，并且把URL对应的处理函数加入进Context的handlers的最后，这样只需要把handlers中的function依次执行就可以了
* 支持了静态资源服务和HTML模板渲染