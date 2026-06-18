# certificate-cli

腾讯云 SSL 证书管理命令行工具，支持证书申请、查看、下载、部署、吊销和删除。

## 安装

```bash
make darwin-arm64   # macOS Apple Silicon
make darwin-amd64   # macOS Intel
make linux-amd64    # Linux x86_64
make windows-amd64  # Windows x86_64
make                # 全平台
```

编译产物位于 `build/` 目录。

## 配置

在可执行文件同目录下创建 `config.yaml`：

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

## 命令一览

| 命令 | 简写 | 参数 | 说明 |
|---|---|---|---|
| `help` | `h` | — | 显示帮助信息 |
| `apply` | `a` | `<域名...>` | 申请免费证书，支持多个域名（10次/秒限频） |
| `list` | `l` | `[页码]` | 获取证书列表，每页100条 |
| `detail` | `dt` | `<证书ID>` | 获取证书详细信息 |
| `keys` | `k` | `<证书ID>` | 获取证书私钥（PEM格式） |
| `download` | `dl` | `<证书ID> [目标路径]` | 下载证书ZIP |
| `batch-download` | `bd` | `<文件夹> <证书ID...>` | 批量下载证书到指定文件夹 |
| `deploy` | `dp` | `<证书ID> <服务器类型> <文件夹>` | 部署证书文件到指定文件夹 |
| `url` | `u` | `<证书ID> [服务类型]` | 获取证书临时下载链接 |
| `revoke` | `rk` | `<证书ID...>` | 吊销证书，支持多个（10次/秒限频） |
| `delete` | `rm` | `<证书ID...>` | 删除证书，支持多个（批量接口） |
| `verify` | `vf` | `<证书ID>` | 检查域名验证状态 |

## 示例

### 申请证书

```bash
# 单个域名
certificate-cli apply example.com
# 多个域名
certificate-cli a example.com example2.com
```

### 查看证书列表

```bash
certificate-cli list        # 第1页
certificate-cli l 2         # 第2页
```

### 查看证书详情和密钥

```bash
certificate-cli dt abc123
certificate-cli k abc123
```

### 下载证书

```bash
# 下载到当前目录：./abc123.zip
certificate-cli dl abc123

# 下载到指定目录：./ssl/abc123.zip
certificate-cli dl abc123 ./ssl/

# 下载为指定文件名：./mycert.zip
certificate-cli download abc123 ./mycert.zip

# 批量下载
certificate-cli bd ./certs/ abc123 xyz456
```

### 部署证书

将证书 ZIP 内对应服务器类型的文件提取到目标文件夹：

```bash
certificate-cli dp abc123 nginx /etc/nginx/ssl/
certificate-cli dp abc123 tomcat /opt/tomcat/ssl/
```

支持的服务器类型：`Apache` / `IIS` / `Nginx` / `Tomcat`（不区分大小写）。

ZIP 内文件结构如下，`deploy` 匹配父目录名并提取所有文件：

```
<certID>
├── Apache
│   ├── 1_root_bundle.crt
│   ├── 2_<domain>.crt
│   └── 3_<domain>.key
├── IIS
│   ├── <domain>.pfx
│   └── keystorePass.txt
├── Nginx
│   ├── 1_<domain>_bundle.crt
│   └── 2_<domain>.key
├── Tomcat
│   ├── <domain>.jks
│   └── keystorePass.txt
└── ...
```

### 获取下载链接

```bash
certificate-cli u abc123           # 默认为 nginx 类型
certificate-cli u abc123 apache    # 指定服务类型
```

### 吊销和删除证书

```bash
# 吊销
certificate-cli rk abc123
certificate-cli rk abc123 xyz456

# 删除
certificate-cli rm abc123
certificate-cli rm abc123 xyz456
```

> 吊销接口无批量能力，逐个调用（间隔100ms，10次/秒）；删除使用 `DeleteCertificates` 批量接口，一次调用。

### 检查域名验证状态

```bash
certificate-cli vf abc123
```

## 依赖

- Go >= 1.22
- [腾讯云 SSL SDK](https://github.com/tencentcloud/tencentcloud-sdk-go)
