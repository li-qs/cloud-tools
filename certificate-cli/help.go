package main

import "fmt"

func HelpInfo() {
	fmt.Println(`用法：certificate-cli <命令> [参数...]

可用命令（括号内为简写）：
  help               (h)    显示本帮助信息
  apply    <域名...> (a)    申请免费证书（支持多个域名，频率限制10次/秒）
  list     [页码]   (l)    获取证书列表（可选页码，默认第1页）
  detail   <证书ID>  (dt)  获取证书的详细信息
  keys     <证书ID>  (k)   获取证书的私钥（PEM格式）
  download <证书ID> [目标路径] (dl) 下载证书为ZIP压缩包。
                                       • 不指定路径时：保存为 ./<证书ID>.zip
                                       • 路径以斜杠结尾（如 ./ssl/）：保存为 <路径><证书ID>.zip
                                       • 否则：将路径视为完整文件路径（如 ./mycert.zip）
  batch-download <文件夹> <证书ID...> (bd) 批量下载证书到指定文件夹，文件名 <证书ID>.zip
  deploy   <证书ID> <服务器类型> <文件夹> (dp) 将证书文件部署到指定文件夹。
                                      服务器类型: Apache, IIS, Nginx, Tomcat
                                      从证书ZIP中提取对应服务器类型的文件并写入目标文件夹
  url      <证书ID>  (u)   获取证书的临时下载链接
  revoke   <证书ID...> (rk) 吊销证书（支持多个证书ID，频率限制10次/秒）
  delete   <证书ID...> (rm) 删除证书（支持多个证书ID，使用批量删除接口）
  verify   <证书ID>  (vf)  检查域名验证状态（用于待审核申请）

示例：
  certificate-cli a example.com example2.com          # 简写申请
  certificate-cli l 2                                 # 简写列表
  certificate-cli dl abc123 ./ssl/                    # 保存为 ./ssl/abc123.zip
  certificate-cli download abc123 ./mycert.zip        # 保存为 ./mycert.zip
  certificate-cli download abc123                     # 保存为 ./abc123.zip
  certificate-cli bd ./certs/ abc123 xyz456           # 简写批量下载
  certificate-cli rk abc123 xyz456                    # 简写吊销
  certificate-cli rm abc123 xyz456                    # 简写删除
  certificate-cli dp abc123 nginx /etc/nginx/ssl/     # 部署Nginx证书文件

更多帮助，请参考文档。`)
}
