package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//content := "abc123"
	keys := "abc123"
	if len(os.Args) < 3 {
		fmt.Println("使用说明:\n" +
			"     -e:\t表示加密\n" +
			"     -d:\t表示解密\n" +
			"srcfile:\t表示要加密的文件\n" +
			"desfile：\t表示加密后的输出文件\n\n" +
			"栗子:\n   des_cbc -e src.txt des.txt\t将src.txt中的数据加密后写入到des.txt文件中\n" +
			"   des_cbc -d src.txt des.txt\t将src.txt中的数据解密后写入到des.txt文件中\r\n")
		os.Exit(0)
	}
	var buffer bytes.Buffer
	head := []byte(len([]byte(keys)))

	fmt.Println("head len: ", len(head))
	buffer.Write(head)
	buffer.Write([]byte(keys))
	fmt.Println("====bytes w: " , buffer.Bytes())

	//用于生成密钥
	key := GetKeyDES_CBC(keys)

	fileName := os.Args[2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("未找到待处理文件")
		os.Exit(0)
	}
	defer file.Close()
	//读取文件内容
	plain, _ := ioutil.ReadAll(file)

	switch os.Args[1] {
	case "-d":
		DecryptDES_CBC_WriteFile(plain, key, os.Args[3])
	case "-e":
		EncryptDES_CBC_WriteFile(plain, key, os.Args[3])
	}
}

func GetKeyDES_CBC(keys string) []byte {
	tmp := sha256.Sum256([]byte(keys))
	key := tmp[:24]
	fmt.Println("=========arg1:" + hex.EncodeToString(key))

	return key
}

//CBC加密
func EncryptDES_CBC(content, key []byte) string {
	fmt.Println("加密前hex： " + hex.EncodeToString(content))
	//创建block
	block, _ := des.NewTripleDESCipher(key)

	EncryptMode := cipher.NewCBCEncrypter(block, key[:8])
	//明文补足PKCS5Padding
	cryptbyte := PKCS5Padding(content, block.BlockSize())

	crypted := make([]byte, len(cryptbyte))
	EncryptMode.CryptBlocks(crypted, cryptbyte)

	fmt.Println("加密后hex： " + hex.EncodeToString(crypted))

	cryptedStr := base64.StdEncoding.EncodeToString(crypted)

	return cryptedStr
}

//将加密后的数据保存在文件中
func EncryptDES_CBC_WriteFile(content, key []byte, fileName string) bool {
	cryptedStr := EncryptDES_CBC(content, key)

	err := ioutil.WriteFile(fileName, []byte(cryptedStr), 0777)
	if err != nil {
		fmt.Println("保存加密后文件失败!")
		return false
	}
	fmt.Println("文件已加密,务必记住加密key!")
	return true
}

//CBC解密
func DecryptDES_CBC(content, key []byte) []byte {
	//创建block
	block, _ := des.NewTripleDESCipher(key)

	DecryptMode := cipher.NewCBCDecrypter(block, key[:8])
	plain, _ := base64.StdEncoding.DecodeString(string(content))
	DecryptMode.CryptBlocks(plain, plain)
	plain = PKCS5UnPadding(plain)
	fmt.Println("解密成功： " + string(plain))

	return plain
}

//将加密后的数据保存在文件中
func DecryptDES_CBC_WriteFile(content, key []byte, fileName string) bool {
	decrypyed := DecryptDES_CBC(content, key)
	err := ioutil.WriteFile(fileName, decrypyed, 0777)
	if err != nil {
		fmt.Println("保存解密后文件失败!")
		return false
	} else {
		fmt.Println("文件已解密!")
	}
	fmt.Println("解密后：" + string(decrypyed))
	return true
}

//明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
