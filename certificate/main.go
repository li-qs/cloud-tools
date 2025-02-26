package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"liqs.net/cloud-tools/certificate/client"
)

var (
	secretId      string
	secretKey     string
	certificateId string
	domainName    string
	contactEmail  string
	contactPhone  string
)

func init() {
	// 读取命令行参数
	flag.StringVar(&secretId, "secretId", "", "腾讯云 secret id")
	flag.StringVar(&secretKey, "secretKey", "", "腾讯云 secret key")
	flag.StringVar(&certificateId, "certificateId", "", "证书 ID")
	flag.StringVar(&domainName, "domain", "", "证书绑定域名")
	flag.StringVar(&contactEmail, "email", "", "联系邮箱")
	flag.StringVar(&contactPhone, "phone", "", "联系电话")
	flag.Parse()
}

func main() {
	var err error

	// 初始化客户端
	var client client.Client
	err = client.Init(secretId, secretKey)
	if err != nil {
		fmt.Printf("客户端初始化失败: %v\n", err)
		return
	}

	// 申请免费证书
	if certificateId == "" {
		fmt.Printf("域名：%v\n开始申请证书...\n", domainName)

		certificateId, err = client.ApplyCertificate(domainName, contactEmail, contactPhone)
		if err != nil {
			fmt.Printf("申请失败：%v\n", err)
			return
		}

		fmt.Printf("申请已提交，certificateId：%v\n", certificateId)
	}

	// 轮询审核状态
	waitingForReview(client)

	// 下载证书
	fmt.Println("开始下载...")
	filename := domainName + "_" + certificateId + ".zip"
	err = client.DownloadCertificate(certificateId, filename)
	if err != nil {
		fmt.Printf("下载失败：%v\n", err)
	} else {
		fmt.Printf("下载成功 --> %v\n", filename)
	}
}

func waitingForReview(c client.Client) error {
	fmt.Println("正在查询证书状态，请勿退出...")

	for {
		certDetail, err := c.GetCertificateDetail(certificateId)
		if err != nil {
			fmt.Printf("查询证书状态失败：%v\n", err)
			if strings.Contains(err.Error(), "Code=FailedOperation.CertificateNotFound") {
				os.Exit(1)
			} else {
				time.Sleep(20 * time.Second)
				continue
			}
		}

		status := *certDetail.Status
		statusName := *certDetail.StatusName
		if domainName == "" {
			domainName = *certDetail.Domain
			fmt.Printf("域名：%v\n", domainName)
		}

		if status == client.StatusPending ||
			status == client.StatusAddedDNSRecord ||
			status == client.StatusEnterprisePending ||
			status == client.StatusSubmittedPending {
			fmt.Printf("审核中，当前状态：%v\n", statusName)
		} else {
			fmt.Printf("审核结束，当前状态：%v\n", statusName)
			break
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}
