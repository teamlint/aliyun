# sms

[![Build Status](https://travis-ci.org/northbright/aliyun.svg?branch=master)](https://travis-ci.org/northbright/aliyun)
[![GoDoc](https://godoc.org/github.com/northbright/aliyun/sms?status.svg)](https://godoc.org/github.com/northbright/aliyun/sms)

[Golang](https://golang.org) SDK for [aliyun SMS service](https://www.aliyun.com/product/sms).

#### Example

    c := sms.NewClient(accessKeyID, accessKeySecret)
    numbers := []string{"13800138000"}
    ok, resp, err := c.Send(numbers, "my_product", "SMS_0000", `{"code":"1234","product":"ytx"}`)

#### Documentation
* [API References](https://godoc.org/github.com/northbright/aliyun/sms)

#### License
* [MIT License](../LICENSE)
