# gonacos
```
封装nacos 打造 go nacos sdk、go nacos client
```

## 使用限制
```
支持Go>=v1.14版本
支持Nacos>2.x版本
```

## 下载包
```
    go get -u github.com/jellycheng/gonacos
或者
    GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gonacos

直接获取master分支代码：
    go get -u github.com/jellycheng/gonacos@master

```

## 注册服务使用示例
```
package main

import (
	"context"
	"fmt"
	"github.com/jellycheng/gonacos"
	"runtime"
	"time"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	svcDone := make(chan struct{})
	ctx, _ := context.WithCancel(context.Background())
	bTime := 600
	//设置定时任务--心跳
	timeBeat := time.NewTimer(time.Second * time.Duration(bTime))
	go func() {
		defer timeBeat.Stop()
		for {
			select {
			case <-timeBeat.C:
				// 到了重新触发时间
				fmt.Println("到了重新触发时间", time.Now())

				// 重置计时器
				timeBeat.Reset(time.Second * time.Duration(bTime))
			case <-ctx.Done():
				fmt.Println("停止计数")
				return
			}
		}
	}()

	go func() {
		// 注册服务
		nacosServerConfig := gonacos.NacosServerConfig{ApiUrl: "http://console.nacos.io", BeatInterval: 5 * 1000}
		RegisterMobileapi(nacosServerConfig)
		
	}()


	<-svcDone
	fmt.Println("finish")
}

func RegisterMobileapi(nacosServerConfig gonacos.NacosServerConfig)  {
	nacosServiceManage := gonacos.NewNacosServiceManage(nacosServerConfig)
	dto := gonacos.NacosInstanceDto{
		ServiceName: gonacos.StringPtr("mobile-api"),
		Ip:          gonacos.StringPtr("10.1.20.2"),
		Port:        gonacos.Int64Ptr(80),
		NamespaceId: gonacos.StringPtr("dde761c3-96be-4a98-b349-3c2289033322"),
		GroupName:   gonacos.StringPtr("mall"),
	}
	if ok, err := nacosServiceManage.RegisterInstance(dto, true); ok {
		fmt.Println("mobile-api注册成功")
	} else {
		fmt.Println("mobile-api注册失败：", err.Error())
	}
}

```