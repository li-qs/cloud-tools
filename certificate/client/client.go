package client

import (
	"encoding/base64"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

type Client struct {
	Client *ssl.Client // 客户端
}

// 定义状态码常量
const (
	StatusPending            = 0  // 审核中
	StatusApproved           = 1  // 已通过
	StatusFailed             = 2  // 审核失败
	StatusExpired            = 3  // 已过期
	StatusAddedDNSRecord     = 4  // 已添加DNS记录
	StatusEnterprisePending  = 5  // 企业证书，待提交
	StatusCanceling          = 6  // 订单取消中
	StatusCanceled           = 7  // 已取消
	StatusSubmittedPending   = 8  // 已提交资料，待上传确认函
	StatusRevoking           = 9  // 证书吊销中
	StatusRevoked            = 10 // 已吊销
	StatusReissuing          = 11 // 重颁发中
	StatusPendingRevokeProof = 12 // 待上传吊销确认函
)

// 初始化客户端
func (c *Client) Init(secretId string, secretKey string) error {
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = "ssl.tencentcloudapi.com"

	credential := common.NewCredential(secretId, secretKey)
	client, err := ssl.NewClient(credential, "", clientProfile)
	if err != nil {
		return err
	}

	c.Client = client
	return nil
}

// 申请免费证书
func (c *Client) ApplyCertificate(domain string, email string, phone string) (certificateId string, err error) {
	request := ssl.NewApplyCertificateRequest()
	request.DomainName = &domain
	request.ContactEmail = &email
	request.ContactPhone = &phone
	request.CsrEncryptAlgo = common.StringPtr("RSA")
	request.DvAuthMethod = common.StringPtr("DNS_AUTO")
	request.DeleteDnsAutoRecord = common.BoolPtr(true)

	response, err := c.Client.ApplyCertificate(request)
	if err != nil {
		return "", err
	}

	return *response.Response.CertificateId, nil
}

// 查询证书详细状态信息
func (c *Client) GetCertificateDetail(certificateId string) (*ssl.DescribeCertificateResponseParams, error) {
	request := ssl.NewDescribeCertificateRequest()
	request.CertificateId = &certificateId

	response, err := c.Client.DescribeCertificate(request)
	if err != nil {
		return nil, err
	}

	return response.Response, nil
}

// 下载证书
func (c *Client) DownloadCertificate(certificateId string, targetPath string) error {
	request := ssl.NewDownloadCertificateRequest()
	request.CertificateId = &certificateId

	response, err := c.Client.DownloadCertificate(request)
	if err != nil {
		return err
	}

	content, err := base64.StdEncoding.DecodeString(*response.Response.Content)
	if err != nil {
		return err
	}

	err = os.WriteFile(targetPath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
