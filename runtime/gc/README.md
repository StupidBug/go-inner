# GC

本模块旨在测试项目的 GC 情况,方便了解 GO 的 GC 运行肌理

## GC Assist

测试需要

```shell
go test ./runtime/gc -run TestNoAssitSweep -trace=trace.out
go tool trace trace.out
```
