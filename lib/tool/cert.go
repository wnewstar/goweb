package tool

import (
    "io/ioutil"
    "crypto/x509"
    "encoding/pem"
    "golang.org/x/crypto/pkcs12"
)

/**
 * @desc    解析公钥
 * @name    DecodePublicCert
 * @date    2020-03-06
 * @author  wnewstar
 * @param   filepath        string                          证书路径
 * @return  cert            *x509.Certificate               公钥信息
 */
func DecodePublicCert(
    filepath string,
) (
    cert *x509.Certificate,
    err error,
) {
    var data []byte
    data, err = ioutil.ReadFile(filepath)
    if err == nil {
        temp, _ := pem.Decode(data)
        cert, err = x509.ParseCertificate(temp.Bytes)
    }

    return
}

/**
 * @desc    解析私钥
 * @name    DecodePrivateCert
 * @date    2020-03-06
 * @author  wnewstar
 * @param   filepath        string                          证书路径
 * @param   password        string                          证书密码
 * @return  cert            *x509.Certificate               公钥信息
 * @return  pkey            interface{}                     私钥信息
 */
func DecodePrivateCert(
    filepath string,
    password string,
) (
    cert *x509.Certificate,
    pkey interface{},
    err error,
) {
    var data []byte
    data, err = ioutil.ReadFile(filepath)
    if err == nil { pkey, cert, err = pkcs12.Decode(data, password) }

    return
}
