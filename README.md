# ZM

> Golang任务管理工具

## 使用文档

### 安装

```bash
go install github.com/ZTaboo/ZM@latest
```

### init

> 初始化数据库和web的基本信息，首次使用需此命令

#### 使用示例

```bash
zm init
```

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

### run

> 可直接运行任务，也可后台运行任务

#### 使用示例

```bash
zm run [task name]
zm run backend [task name]
```
### status

> 查看当前任务运行状态：Running/Stop

#### 使用示例

```bash
zm status
```
> output

```bash
+--------+-----------------+------+---------+
|  Name  |      File       | Port | Status  |
+--------+-----------------+------+---------+
| ZM Web |     zm.exe      | 1999 |  Stop   |
|  main  | /root/code/main | 8080 | Running |
+--------+-----------------+------+---------+
```

### delete

> 删除命令会优先停止已运行任务，然后删除此任务

#### 使用示例

```bash
zm delete [task name]
```

### stop

> 停止任务

```bash
zm stop [task name]
```