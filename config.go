package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

var defaultJSON = `{
  "google_api": "AIza[0-9A-Za-z-_]{35}",
  "google_oauth": "ya29\\.[0-9A-Za-z\\-_]+",
  "amazon_aws_access_key_id": "AKIA[0-9A-Z]{16}",
  "amazon_mws_auth_toke" : "amzn\\\\.mws\\\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
  "amazon_aws_url" : "s3\\.amazonaws.com[/]+|[a-zA-Z0-9_-]*\\.s3\\.amazonaws.com",
  "facebook_access_token": "EAACEdEose0cBA[0-9A-Za-z]+",
  "authorization_basic": "basic\\s*[a-zA-Z0-9=:_\\+\\/-]+",
  "authorization_bearer": "bearer\\s*[a-zA-Z0-9_\\-\\.=:_\\+\\/]+",
  "authorization_api": "api[key|\\s*]+[a-zA-Z0-9_\\-]+",
  "mailgun_api_key": "key-[0-9a-zA-Z]{32}",
  "twilio_api_key": "SK[0-9a-fA-F]{32}",
  "paypal_braintree_access_token": "access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}",
  "square_oauth_secret": "sq0csp-[ 0-9A-Za-z\\-_]{43}|sq0[a-z]{3}-[0-9A-Za-z\\-_]{22,43}",
  "square_access_token": "sqOatp-[0-9A-Za-z\\-_]{22}|EAAA[a-zA-Z0-9]{60}",
  "stripe_standard_api": "sk_live_[0-9a-zA-Z]{24}",
  "stripe_restricted_api": "rk_live_[0-9a-zA-Z]{24}",
  "github_access_token": "[a-zA-Z0-9_-]*:[a-zA-Z0-9_\\-]+@github\\.com*",
  "rsa_private_key": "-----BEGIN RSA PRIVATE KEY-----",
  "ssh_dsa_private_key": "-----BEGIN DSA PRIVATE KEY-----",
  "ssh_dc_private_key": "-----BEGIN EC PRIVATE KEY-----",
  "pgp_private_block": "-----BEGIN PGP PRIVATE KEY BLOCK-----",
  "json_web_token": "ey[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_\\.+/=]*$"
}
`

type Re struct {
	Name   string
	Regexp *regexp.Regexp
}

func MustLoadConfig(fp string) []Re {
	var err error
	data := []byte(defaultJSON)
	if fp != "" {
		data, err = ioutil.ReadFile(fp)
		if err != nil {
			log.Fatalf("read config error: %v", err)
		}
	}

	mp := make(map[string]string)
	err = json.Unmarshal(data, &mp)
	if err != nil {
		log.Fatalf("parse config error: %v", err)
	}
	return ConfigFromMap(mp)
}

func ConfigFromMap(mp map[string]string) []Re {
	if len(mp) == 0 {
		log.Fatal("config is empty")
	}
	res := []Re{}
	for k, v := range mp {
		res = append(res, Re{
			Name:   k,
			Regexp: regexp.MustCompile(v),
		})
	}

	return res
}
