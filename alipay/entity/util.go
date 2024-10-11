package entity

type PublicCertDownloadResponse struct {
	Response *PublicCertDownload `json:"alipay_open_app_alipaycert_download_response"`
}

type PublicCertDownload struct {
	ErrorResponse
	AlipayCertContent string `json:"alipay_cert_content"`
}
