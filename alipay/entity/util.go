package entity

type PublicCertDownloadResponse struct {
	Response *PublicCertDownload `json:"alipay_open_app_alipaycert_download_response"`
}

// =========================================================分割=========================================================

type PublicCertDownload struct {
	ErrorResponse
	AlipayCertContent string `json:"alipay_cert_content"`
}
