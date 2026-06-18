package main

import "fmt"

func HelpInfo() {
	fmt.Println(`用法：certificate-cli <命令> [参数...]

可用命令：
  help                              显示本帮助信息
  apply    <域名>                    申请免费证书（单域名）
  list     [页码]                    获取证书列表（可选页码，默认第1页）
  detail   <证书ID>                 获取证书的详细信息
  keys     <证书ID>                 获取证书的私钥（PEM格式）
  download <证书ID> [目标路径]       下载证书为ZIP压缩包。
                                      • 不指定路径时：保存为 ./<证书ID>.zip
                                      • 路径以斜杠结尾（如 ./ssl/）：保存为 <路径><证书ID>.zip
                                      • 否则：将路径视为完整文件路径（如 ./mycert.zip）
  url      <证书ID>                 获取证书的临时下载链接
  revoke   <证书ID>                 吊销证书（立即失效）
  delete   <证书ID>                 删除证书（从系统中移除）
  verify   <证书ID>                 检查域名验证状态（用于待审核申请）

示例：
  certificate-cli apply example.com
  certificate-cli list 2
  certificate-cli download abc123 ./ssl/          # 保存为 ./ssl/abc123.zip
  certificate-cli download abc123 ./mycert.zip    # 保存为 ./mycert.zip
  certificate-cli download abc123                 # 保存为 ./abc123.zip

更多帮助，请参考文档。`)
}
