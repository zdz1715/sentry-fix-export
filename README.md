# sentry-fix-export
更改sentry导出文件字段类型，使其能带入成功

## 发生场景
> [sentry自托管项目](https://github.com/getsentry/self-hosted)22.7.0 -> 22.10.0


根据[sentry官方文档](sentry官方文档)备份之后导入失败，原因是导出的字段类型不一致，所以才有了此脚本转换下字段类型，经测试导入成功，以此记录

![image](https://user-images.githubusercontent.com/24859074/200988626-5f03d8e1-cbe3-448d-a9f8-a6223bf5dccc.png)

### 使用

git clone 
