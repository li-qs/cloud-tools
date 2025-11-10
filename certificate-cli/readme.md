# certificate-cli

腾讯云 SSL 证书管理工具，功能逐步完善中...

## 使用

- 获取帮助信息：`certificate-cli help`
- 申请免费证书：`certificate-cli apply [www.example.com]`
- 获取证书列表：`certificate-cli list` 或 `certificate-cli list [page number]`
- 获取证书信息：`certificate-cli detail [ID]`
- 获取证书密钥：`certificate-cli keys [ID]`
- 下载证书：`certificate-cli download [ID] [/target_dir/target_file.zip]`
- 获取证书下载链接：`certificate-cli url [ID]`
- 吊销证书：`certificate-cli revoke [ID]`
- 删除证书：`certificate-cli delete [ID]`

## 配置

配置文件为 `config.yaml`，内容实例：

```yaml
# 服务地址，详见腾讯云文档：https://cloud.tencent.com/document/api/400/41659
endpoint: ssl.tencentcloudapi.com

# 安全凭证 SecretId
secretID: xxxxxx

# 安全凭证 SecretKey
secretKey: xxxxxx

# 联系邮箱（证书状态更新时发送邮件通知）
contactEmail: abc@example.com

# 联系电话（证书状态更新时发送短信通知）
contactPhone: 13900000000
```
