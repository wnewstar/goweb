package tool

import (
    "bytes"
    "strings"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
)

func replaceEn(src []byte) ([]byte) {
    str := string(src)
    str = strings.Replace(str, "+", "-", -1)
    str = strings.Replace(str, "/", "_", -1)

    return []byte(str)
}

func replaceDe(src []byte) ([]byte) {
    str := string(src)
    str = strings.Replace(str, "-", "+", -1)
    str = strings.Replace(str, "_", "/", -1)

    return []byte(str)
}

func unpadding(src []byte) ([]byte) {

    return src[:len(src) - int(src[len(src) - 1])]
}

func dopadding(src []byte, size int) ([]byte) {
    num := size - len(src) % size
    val := byte(num)

    return append(src, bytes.Repeat([]byte{val}, num)...)
}

func EncryptAes(src []byte, key []byte) ([]byte) {
    blk, _ := aes.NewCipher(key)
    liv := blk.BlockSize()
    src = dopadding(src, liv)
    cipher.NewCBCEncrypter(blk, key[:liv]).CryptBlocks(src, src)

    return replaceEn([]byte(base64.StdEncoding.EncodeToString([]byte(src))))
}

func DecryptAes(src []byte, key []byte) ([]byte) {
    src, _ = base64.StdEncoding.DecodeString(string(replaceDe(src)))
    blk, _ := aes.NewCipher(key)
    liv := blk.BlockSize()
    les := len(src)
    lep := -1
    if les >= liv && (les % liv == 0) {
        cipher.NewCBCDecrypter(blk, key[:liv]).CryptBlocks(src, src)
        les = len(src)
        lep = int(src[les - 1])
    }

    if lep == -1 || les < liv || lep > liv { return []byte{} } else { return []byte(unpadding(src)) }
}
