Glide Demo
==========

使用[glide](https://github.com/Masterminds/glide)进行go项目的依赖管理。

```
brew install glide
```

```
glide create
```

将会生成`glide.yaml`文件

```
glide get github.com/golang-demos/go-concat-strings-library
```

将会下载`https://github.com/Masterminds/cookoo`到本地`vendor`目录。
如果该项目有`tags`，glide将会拿到最新版本号，并且询问对于此库未来的版本更新，我们选择哪种策略：

```
Minor (M), Patch (P), or Skip Ranges (S)?
```

- `Minor (M)`: 在主版本不变的前提下，更新到最新版
- `Patch (P)`: 在主次版本不变的前提下，更新到最新版
- `Skip Ranges (S)`: 固定版本，不再变化

我们可以根据自己的情况选择。比如我的Demo项目，一旦完成可能很久都不需要再更新，为了保证以后随时可以正常运行，所以通常选`S`

则`glide.yaml`中会多一条:

```
import:
- package: github.com/Masterminds/cookoo
  version: v1.3.0
```

同时在`vendor`目录下，也会多出相应的目录：`vendor/github.com/Masterminds`

总结
---

看到glide，就要想到两点：

- `glide.yaml`：记录了依赖
- `verndor`：保存了下载的库
