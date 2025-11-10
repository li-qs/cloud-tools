package main

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

type Api struct {
	SecretID string
	Endpoint string
	client   *ssl.Client
}

func New(endpoint, secretID, secretKey string) (*Api, error) {
	p := profile.NewClientProfile()
	p.HttpProfile.Endpoint = endpoint

	credential := common.NewCredential(secretID, secretKey)
	client, err := ssl.NewClient(credential, "", p)
	if err != nil {
		return nil, err
	}

	return &Api{
		SecretID: secretID,
		Endpoint: endpoint,
		client:   client,
	}, nil
}

// 免费证书申请
func (a *Api) ApplyCertificate(domain, email, phone string) (*ssl.ApplyCertificateResponseParams, error) {
	r := ssl.NewApplyCertificateRequest()
	r.DomainName = &domain
	r.ContactEmail = &email
	r.ContactPhone = &phone
	r.CsrEncryptAlgo = common.StringPtr("RSA")
	r.DvAuthMethod = common.StringPtr("DNS_AUTO")
	r.DeleteDnsAutoRecord = common.BoolPtr(true)
	res, err := a.client.ApplyCertificate(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 获取证书列表
func (a *Api) ListCertificates(offset, limit int) (*ssl.DescribeCertificatesResponseParams, error) {
	o := uint64(offset)
	l := uint64(limit)

	r := ssl.NewDescribeCertificatesRequest()
	r.Offset = &o
	r.Limit = &l
	res, err := a.client.DescribeCertificates(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 获取证书信息
func (a *Api) DescribeCertificate(certificateID string) (*ssl.DescribeCertificateResponseParams, error) {
	r := ssl.NewDescribeCertificateRequest()
	r.CertificateId = &certificateID
	res, err := a.client.DescribeCertificate(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 获取证书详情
func (a *Api) DescribeCertificateDetail(certificateID string) (*ssl.DescribeCertificateDetailResponseParams, error) {
	r := ssl.NewDescribeCertificateDetailRequest()
	r.CertificateId = &certificateID
	res, err := a.client.DescribeCertificateDetail(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 获取证书下载链接
func (a *Api) DescribeDownloadCertificateUrl(certificateID, serviceType string) (*ssl.DescribeDownloadCertificateUrlResponseParams, error) {
	r := ssl.NewDescribeDownloadCertificateUrlRequest()
	r.CertificateId = &certificateID
	r.ServiceType = &serviceType
	res, err := a.client.DescribeDownloadCertificateUrl(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 下载证书
func (a *Api) DownloadCertificate(certificateID string) (*ssl.DownloadCertificateResponseParams, error) {
	r := ssl.NewDownloadCertificateRequest()
	r.CertificateId = &certificateID
	res, err := a.client.DownloadCertificate(r)
	if err != nil {
		return nil,err
	}

	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 吊销证书
func (a *Api) RevokeCertificate(certificateID string) (*ssl.RevokeCertificateResponseParams, error) {
	r := ssl.NewRevokeCertificateRequest()
	r.CertificateId = &certificateID
	res, err := a.client.RevokeCertificate(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 删除证书
func (a *Api) DeleteCertificate(certificateID string) (*ssl.DeleteCertificateResponseParams, error) {
	r := ssl.NewDeleteCertificateRequest()
	r.CertificateId = &certificateID
	res, err := a.client.DeleteCertificate(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

// 批量删除证书
func (a *Api) DeleteMultiCertificates(certificateIDs []*string) (*ssl.DeleteCertificatesResponseParams, error) {
	r := ssl.NewDeleteCertificatesRequest()
	r.CertificateIds = certificateIDs
	res, err := a.client.DeleteCertificates(r)
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}
