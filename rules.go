package main

import (
	"encoding/json"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var ruleCount int

type RuleConfig struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Type           string     `json:"type"`
	Value          string     `json:"value"`
	Conditions     Conditions `json:"conditions"`
	RejectMessage  string     `json:"rejectMessage"`
	HTTPStatusCode int        `json:"httpStatusCode"`
}

type Conditions struct {
	HourRange []int `json:"hourRange"`
}

func ParseRules() *RuleConfig {
	buf, err := ioutil.ReadFile("./config/rules.json")
	if err != nil {
		log.Fatal("silverplate: cannot read ./config/rules.json. do you have permission or is it missing?")
	}
	var ruleConfig RuleConfig
	err = json.Unmarshal(buf, &ruleConfig)
	if err != nil {
		log.Fatal("silverplate can't parse rules.json")
	}
	return &ruleConfig
}

func (rc *RuleConfig) ApplyRules(proxy *goproxy.ProxyHttpServer) {
	for _, rule := range rc.Rules {
		var reqCond goproxy.ReqCondition
		switch rule.Type {
		case "UrlHasPrefix":
			reqCond = goproxy.UrlHasPrefix(rule.Value)
		case "UrlIs":
			reqCond = goproxy.UrlIs(rule.Value)
		case "ReqHostMatches":
			reqCond = goproxy.ReqHostMatches(regexp.MustCompile(rule.Value))
		case "ReqHostIs":
			reqCond = goproxy.ReqHostIs(rule.Value)
		case "UrlMatches":
			reqCond = goproxy.UrlMatches(regexp.MustCompile(rule.Value))
		case "DstHostIs":
			reqCond = goproxy.DstHostIs(rule.Value)
		case "SrcIpIs":
			reqCond = goproxy.SrcIpIs(rule.Value)
		}
		proxy.OnRequest(reqCond).DoFunc(
			func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				if len(rule.Conditions.HourRange) > 0 {
					if h, _, _ := time.Now().Clock(); h >= rule.Conditions.HourRange[0] && h <= rule.Conditions.HourRange[1] {
						return r, goproxy.NewResponse(r, goproxy.ContentTypeText, rule.HTTPStatusCode, rule.RejectMessage)
					} else {
						return r, nil
					}
				}
				return r, nil
			})
		ruleCount++
	}
}

func RuleCount() int {
	return ruleCount
}

func RuleCountToString() string {
	return strconv.FormatInt(int64(ruleCount), 10)
}
