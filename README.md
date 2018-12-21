# todolist-go-webassembly

## 目的
一直想学习一下 Go 语言，刚好 Go 1.11 支持 WebAssembly，WebAssembly 也是经常听过但是自己没有去了解。

抱着学习的目的，就写了这样一个 Demo 来作为练习。

## 运行
wasm 文件的加载使用 fetch 必须使用，本地需要启动一个 http 服务，某些服务程序不能正确的将 wasm 文件的 Content-Type 设置正确，所以抄了一个 go 写的 HTTP Service（因为 Go 技艺不精湛，还没有看到网络部分）。

运行 `go run server.go`，访问 [http://localhost:8080](http://localhost:8080) 即可运行。

## 实现的功能

* Go 对 DOM 的操作
* JS 中加载 Go 编译的 wasm
* 在 GO 中绑定 DOM 事件

## 碰到的问题
如果 Go 源文件分成多个文件，使用同样的 package name，编译出来的 wasm 在浏览器中运行，会报一个很奇怪的错误，把所有代码合到一个文件里面编译就不会出现这个问题。

错误内容多为内存中的错误，我也是一头雾水，只能暂且记录下来，待往后深入的学习能有所解答。
