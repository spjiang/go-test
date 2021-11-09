package util

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"time"
)

type CreatorAuth struct {
	Body string
	Sign string
}

// createExpireSign 创建token
func (c *CreatorAuth) CreateExpireSign() (r string, err error) {
	var h hash.Hash
	secret := "504e3bdec1d8ec364fed5e8df410a560" //os.Getenv("AUTH_SECRET")
	//if len(secret) <= 0 {
	//	return "", errors.New("The auth secret key is not set. ")
	//}
	if len(c.Body) <= 0 {
		return "", errors.New("The body is empty. ")
	}
	// 有效期
	ttl := "6000" //os.Getenv("AUTH_TTL")
	//if len(ttl) <= 0 {
	//	return "", errors.New("The ttl is empty. ")
	//}

	ttlInt, err := strconv.ParseInt(ttl, 10, 64)
	if ttlInt <= 0 {
		return "", errors.New("The ttl configuration error. ")
	}
	expire := time.Now().Unix() + ttlInt // 单位s,打印结果:1491888244
	// 加密数据
	src := fmt.Sprintf("%s+%s+%d", secret, c.Body, expire)
	h = sha256.New()
	h.Write([]byte(src))
	// 一次加密后的结果，可以进行二次加密
	r = fmt.Sprintf("%x", h.Sum(nil)) + strconv.FormatInt(expire, 10)
	c.Sign = r
	fmt.Println(c.Sign)
	return
}

// VerifyExpireSign 验证
func (c *CreatorAuth) VerifyExpireSign() (bool, error) {
	var h hash.Hash
	if len(c.Body) <= 0 {
		return false, errors.New("The body is empty. ")
	}
	secret := "504e3bdec1d8ec364fed5e8df410a560" //os.Getenv("AUTH_SECRET")
	//if len(secret) <= 0 {
	//	return false, errors.New("The auth secret key is not set. ")
	//}
	if len(c.Sign) <= 10 {
		return false, errors.New("The sign is illegal. ")
	}
	// 签名对比
	// 截取sign后10位
	expire := c.Sign[len(c.Sign)-10:]
	expireInt, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		return false, errors.New("The sign is illegal. ")
	}
	// 加密数据
	src := fmt.Sprintf("%s+%s+%d", secret, c.Body, expireInt)
	h = sha256.New()
	h.Write([]byte(src))
	sign := fmt.Sprintf("%x", h.Sum(nil)) + expire
	if c.Sign != sign {
		return false, errors.New("The sign is illegal. ")
	}
	fmt.Println(sign)
	// 验证有效期
	if expireInt < time.Now().Unix() {
		return false, errors.New("The sign has expired. ")
	}
	return true, nil
}
