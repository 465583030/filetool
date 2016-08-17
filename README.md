# filetool

```
> filetool
[D]   42K drwxr-xr-x /filetool
|-- [D]   27K drwxr-xr-x .git
|-- [F]  266B -rw-r--r-- .gitignore
|-- [F]   11K -rw-r--r-- LICENSE
|-- [F]   10B -rw-r--r-- README.md
|-- [F]  966B -rw-r--r-- filetool.go
|-- [F]    1K -rw-r--r-- tool.go
|-- [F]  470B -rw-r--r-- tool_test.go
`-- [F]  924B -rw-r--r-- tree.go
```

做了一个查看文件大小的小工具,以树形显示目录结构,每行显示的内容包括

```
[文件F或目录D] 大小 权限 文件或目录名
```

目录的大小是该目录下所有文件大小的和。

**注意: 这个工具会遍历当前目录下所有的文件,如果文件较多,慎用!**


## 设置树深度

```
filetool -d 3
```

默认值是2, 第1层是当前目录。若为-1则显示全部。


## 过滤文件大小

```
filetool -size 10M
```

只显示大小超过 10M 的文件或目录。

支持的标记如: `12`, `12B`, `12K`, `12M`, `12G`, `12T`, 大小写不敏感。

