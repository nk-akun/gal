### Practice for Web Framework Router based golang.

### 要点：
* Server实现了请求入口ServeHTTP, http.ListenAndServe只需要一个实现了ServeHTTP的interface即可.
* 实现了Context，封装了http.ResponseWriter, *http.Request, Method, Path等内容,并据Context实现了若干常用方法.
* Router中使用map存储路由表, key为"method-path".