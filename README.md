# ZM

> Golang项目管理工具

## 使用文档

### add

> 该命令行工具用于添加任务，将程序路径、名称、端口号以及程序所在目录添加到数据库中。

#### 使用方法

```bash
add [flags] <program>

flags:
  -n, --name string   为任务命名 (必填)
  -p, --port int      任务程序端口 (必填)

args:
  program             要添加的程序名称 (必填)
```
#### 选项
-n, --name string
为任务命名，是必填选项。

-p, --port int
任务程序端口，是必填选项。

#### 参数
program
要添加的程序名称，是必填参数。

```bash
add --name mytask --port 8080 myprogram
```