这个文件记录了glide相关信息，它是怎么创建出来的？
-----------------------------------------

```
glide create
```

它在生成`glide.yaml`文件的同时，会搜索`vendor`目录和项目的Go代码中的`import`，添加到`glide.yaml`中。

如果需要引入新的依赖，可以：

```
glide get github.com/golang-demos/go-concat-strings-library
```

它会把相应git仓库中的代码下载到`verdor`目录下，并更新`glide.yaml`文件。

如果该依赖以tags（`git tag`）的形式记录版本号，则glide会使用最新版本，并询问我们采用哪种版本更新策略：

```
Minor (M), Patch (P), or Skip Ranges (S)?
```

（详情可参见输出的日志）

如果我们选了`M`，则`glide.yaml`中会多一条:

```
import:
- package: github.com/golang-demos/go-concat-strings-library
  version: ^0.2.0
```

同时在`vendor`目录下也将会多出`github.com/golang-demos/go-concat-strings-library`目录。

