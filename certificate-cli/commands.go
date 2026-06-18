package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ApplyCertificate(args []string) error {
	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	domain := args[0]
	res, err := api.ApplyCertificate(domain, config.ContactEmail, config.ContactPhone)
	if err != nil {
		return err
	}

	str, _ := json.MarshalIndent(*res, "", "  ")
	fmt.Printf("申请成功！\n%s\n", str)

	return nil
}

func ListCertificates(args []string) error {
	var err error
	page := 1

	if len(args) > 1 {
		return errors.New("参数错误！")
	}

	if len(args) == 1 {
		page, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if page <= 0 {
			page = 1
		}
	}

	pageSize := 100
	offset := (page - 1) * pageSize

	res, err := api.ListCertificates(offset, pageSize)
	if err != nil {
		return err
	}

	fmt.Printf("证书总数：%d，本页：%d，页码：%d，每页：%d\n", *res.TotalCount, len(res.Certificates), page, pageSize)
	fmt.Printf("%s	%s	%s	%s	%s	%s\n",
		"证书ID",
		"绑定域名",
		"到期时间",
		"创建时间",
		"证书类型",
		"状态",
	)
	for _, cert := range res.Certificates {
		fmt.Printf("%s	%s	%s	%s	%s	%s\n",
			*cert.CertificateId,
			*cert.Domain,
			*cert.CertEndTime,
			*cert.InsertTime,
			*cert.PackageTypeName,
			*cert.StatusName,
		)
	}

	return nil
}

func DescribeCertificate(args []string) error {
	var err error

	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	res, err := api.DescribeCertificate(cID)
	if err != nil {
		return err
	}

	str, _ := json.MarshalIndent(*res, "", "  ")
	fmt.Println(string(str))

	return nil
}

func CertificateKeyPair(args []string) error {
	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	res, err := api.DescribeCertificateDetail(cID)
	if err != nil {
		return err
	}

	if res.EncryptCert != nil {
		fmt.Printf("SHA1指纹：%s\n国密加密算法：%s\n国密加密证书公钥：%s\n国密加密证书私钥：%s\n",
			*res.EncryptCertFingerprint,
			*res.EncryptAlgorithm,
			*res.EncryptCert,
			*res.EncryptPrivateKey,
		)
	} else {
		fmt.Printf("SHA1指纹：%s\n加密算法：%s\n证书公钥：%s\n证书私钥：%s\n",
			*res.CertFingerprint,
			*res.EncryptAlgorithm,
			*res.CertificatePublicKey,
			*res.CertificatePrivateKey,
		)
	}

	return nil
}

func DescribeDownloadCertificateUrl(args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	sType := "nginx"
	if len(args) == 2 {
		sType = args[1]
	}

	res, err := api.DescribeDownloadCertificateUrl(cID, sType)
	if err != nil {
		return err
	}

	fmt.Printf("下载文件名：%s\n下载链接：%s\n",
		*res.DownloadFilename,
		*res.DownloadCertificateUrl,
	)

	return nil
}

func DownloadCertificate(args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.New("参数错误！需要1个或2个参数")
	}

	cID := args[0]
	var target string

	if len(args) == 1 {
		// 默认：当前目录下的 <cID>.zip
		target = fmt.Sprintf("./%s.zip", cID)
	} else {
		argPath := args[1]
		if strings.HasSuffix(argPath, string(os.PathSeparator)) {
			// 是目录，拼接文件名
			target = filepath.Join(argPath, cID+".zip")
		} else {
			// 不是目录，直接作为文件路径
			target = argPath
		}
	}

	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, 0755); err != nil {
		absDir, _ := filepath.Abs(dir)
		return fmt.Errorf("无法创建目录 %s: %v", absDir, err)
	}

	if _, err := os.Stat(target); err == nil {
		absPath, _ := filepath.Abs(target)
		return fmt.Errorf("file exists: %s", absPath)
	}

	res, err := api.DownloadCertificate(cID)
	if err != nil {
		return err
	}

	content, err := base64.StdEncoding.DecodeString(*res.Content)
	if err != nil {
		return err
	}
	if err := os.WriteFile(target, content, 0644); err != nil {
		return err
	}

	fmt.Printf("已下载至：%s\n", target)
	return nil
}

func RevokeCertificate(args []string) error {
	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	res, err := api.RevokeCertificate(cID)
	if err != nil {
		return err
	}

	str, _ := json.MarshalIndent(*res, "", "  ")
	fmt.Printf("吊销成功！\n%s\n", str)

	return nil
}

func DeleteCertificate(args []string) error {
	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	res, err := api.DeleteCertificate(cID)
	if err != nil {
		return err
	}

	str, _ := json.MarshalIndent(*res, "", "  ")
	fmt.Printf("删除成功！\n%s\n", str)

	return nil
}

func CheckCertificateDomainVerification(args []string) error {
	if len(args) != 1 {
		return errors.New("参数错误！")
	}

	cID := args[0]
	res, err := api.CheckCertificateDomainVerification(cID)
	if err != nil {
		return err
	}

	localCheckResult := ""
	switch *res.VerificationResults[0].LocalCheck {
	case 1:
		localCheckResult = "验证通过"
	case -1:
		localCheckResult = "被限频 或 找不到 txt 记录"
	case -2:
		localCheckResult = "找不到 txt 记录"
	case -3:
	case -8:
		localCheckResult = "找不到 ns 记录"
	case -4:
	case -9:
		localCheckResult = "找不到文件"
	case -5:
	case -10:
		localCheckResult = "文件不匹配"
	case -6:
		localCheckResult = "找不到 cname 记录"
	case -7:
		localCheckResult = "cname 记录不匹配"
	}

	caCheckResult := ""
	switch *res.VerificationResults[0].CaCheck {
	case -1:
		caCheckResult = "未检测通过"
	case 2:
		caCheckResult = "检测通过"
	}

	fmt.Printf("绑定域名：%s\n证书是否已签发：%t\n验证类型：%s\n腾讯云检测结果：%s\nCA 检测结果：%s\n是否被限频拦截：%t\n失败原因：%s\n",
		*res.VerificationResults[0].Domain,
		*res.VerificationResults[0].Issued,
		*res.VerificationResults[0].VerifyType,
		localCheckResult,
		caCheckResult,
		*res.VerificationResults[0].Frequently,
		*res.VerificationResults[0].LocalCheckFailReason,
	)

	return nil
}
