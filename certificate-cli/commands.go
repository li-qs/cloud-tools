package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	fmt.Printf("%s	%s	%s	%s	%s	%s	%s	%s\n",
		"证书ID",
		"主域名",
		"验证类型",
		"证书类型",
		"有效时间",
		"创建时间",
		"状态",
		"备注",
	)
	for _, cert := range res.Certificates {
		fmt.Printf("%s	%s	%s	%s	%s	%s	%s	%s\n",
			*cert.CertificateId,
			*cert.Domain,
			*cert.VerifyType,
			*cert.PackageTypeName,
			*cert.CertBeginTime+" 至 "+*cert.CertEndTime+"（"+*cert.ValidityPeriod+"个月）",
			*cert.InsertTime,
			*cert.StatusName,
			*cert.Alias,
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
	if len(args) != 2 {
		return errors.New("参数错误！")
	}

	target := args[1]
	if f, _ := os.Stat(target); f != nil {
		p, _ := filepath.Abs(target)
		return fmt.Errorf("file exists: %s", p)
	}

	dir := filepath.Dir(target)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		dir, _ = filepath.Abs(dir)
		return fmt.Errorf("no such directory: %s", dir)
	}

	cID := args[0]
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
