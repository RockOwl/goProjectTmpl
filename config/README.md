# Go配置库

config是一个组件化的配置库，提供了一种简单方式读取多种内容源、多种文件类型的配置。

## 功能

- 插件化：根据配置需要可从多个内容源（本地文件、others等）加载配置。

- 热加载：变更时自动载入新配置

- 默认设置：配置由于某些原因丢失时，可以回退使用默认值。

## 相关结构

- Config： 配置的通用接口，提供了相关的配置读取操作。

- ConfigLoader： 配置加载器，通过实现ConfigLoader相关接口以支持加载策略。

- Codec： 配置编解码接口，通过实现Codec相关接口以支持多种类型配置。

- DataProvider: 内容源接口，通过实现DataProvider相关接口以支持多种内容源。目前支持`file`等配置内容源。

## 如何使用

```go

// 加载本地配置文件: config.WithProvider("file")
config.Load("../app.yaml", config.WithCodec("yaml"), config.WithProvider("file"))

// 默认的DataProvider是使用本地文件
config.Load("../app.yaml", config.WithCodec("yaml"))

// 读取bool类型配置
c.GetBool("server.debug", false)

// 读取String类型配置
c.GetString("server.app", "default")

```

### 并发安全的监听远程配置变化

```go
import (
	"sync/atomic"
    ...
)

type yamlFile struct {
    Server struct {
        App string
    }
}

// 参考：https://golang.org/pkg/sync/atomic/#Value
var cfg atomic.Value // 并发安全的Value

// 使用config中Watch接口监听 配置中心 远程配置变化
c, _ := config.Get("remote").Watch(context.TODO(), "test.yaml")

go func() {
    for r := range c {
        yf := &yamlFile{}
        fmt.Printf("event: %d, value: %s", r.Event(), r.Value())

        if err := yaml.Unmarshal([]byte(r.Value()), yf); err == nil {
            cfg.Store(yf)
        }

    }
}()

// 当配置初始化完成后，可以通过 atomic.Value 的 Load 方法获得最新的配置对象
cfg.Load().(*yamlFile)

```

