package democrypt

import (
	"math/rand"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(length int) string {
	return stringWithCharset(length, charset)
}

var data = []string{
	"th1s1smyp@ssw0rd",
	randomString(4),
	randomString(16),
	randomString(100),
}
var aes128Hash = "1234567890123456"
var aes192Hash = "123456789012345612345678"
var aes256Hash = "12345678901234561234567890123456"

// C#

// AES Encrypted Values:
var CS_AES_CFB8_128 = "vZYJl180MSjG4C2FLhqkJ6cN6GI7/YJ8viZfe+Q6tcA="
var CS_AES_CFB8_192 = "07z3M8zgEtv/IgyDEbr86oKMESgRXxfTqIffedWV3Bw="
var CS_AES_CFB8_256 = "Gg5IAfmBmTB/DFIfJJ1Wts+e8l0nKNVa6XAXFVfEVrw="
var CS_AES_CBC_128 = "LM9SeQ+gsDRfyf0jnNpjA83D/1b/KSMpM+/VxtPiCKNnIpCTQcMURkHNFYTl3fzk"
var CS_AES_CBC_192 = "dUI9ySK67BE4QZspNNaGRea7bAwJBUtiPb5d21r5pDLjBPCHfUl0lgGaTma9IXkx"
var CS_AES_CBC_256 = "w8UbBtYvXBq4wXzAHVu/0+Ftr+zPTukDnFCQ0rqANGCfsotUH3akq0IVKl3jgSEw"

// RSA Encrypted Values:
var CS_RSA_PKCS1V1_5 = "QLAPKvBmzUFyN5bv0K4VCrhaGe6lh5Zfcd/6fGFb576ybRRWXdgYrpM4GOAAQ4fgsSRm6C9xLuJn5iRXyPZ8BJhFEqvdFpV2j+jac21IIkxG7sqTyLQn4FF6FdUf1bC3fPrC8fNS+J59mpUE5p4dDia3im/YgA4LYaOlY0VGLR5sukdrt2EsWtmHN41a9XDRtklSdfP3ReU9K0NlX+hL3SvJ35TZxkEZMvMKQXYKkfWvhvJARbrZY6JBae1p2n8kRipK0hXDbVgy2kraL1lBYJLGFstsexG9DpL7ydIgxYO3mckS+YoGT5irSM6uoEs5Ffi2YJ2JNWaxwHqTs4fz2Q=="
var CS_RSA_OAEP = "LjgIXwqQYp+Rx4knbuj22e8q0+ixtjIxsdbFdhweXTmqtaudNZLqI0HsaiwIvw3r8/OokVHR9zZtbpGPQinPX4jQf0uQ29FUufd3oYdhC1cV7hjvfQDuYELXl5kKqPUtdYPqgD/tapZAXW8F3RLroxe36I96wjbhrcLZhXu8n0FLvgbbIDZkQ8AgNxVpthqb7LIrVMjIPq3gqlImLZMyDbPhsMnzW708eFBTHlR2dh1H93nj6Tsh+971Mbfmo4iboLRtJzUg2CvrfEHv3SOARKkKFconINPkaTHyZYUW/UIbeTF/Ui1aLpjMkhcUUqKFMVGmH7ZlmT6gHGUZVEy13w=="

// X509 Encrypted Values:
var CS_X509_OAEP = "3/Z5MG7SmIrEpPDmHeB8lzGVLNA96gYPacv4nF2vD5u2RuDidkgAsbxR9171H7B07gFyQklis/9m7RCrT4b0V6KMocwkBVXoBwxwIbYi9ErTYvez+gj7UjUQvE7sHpNWei15UdF+78tqrh37ygRiG2zE6+pWyE82OaEPROVoXwSDMPrxJtVW1P3JsjKybA1kEB/iPn/gdhZixDxoWScUKZHsxD613qs1jsPQ23paJARwWlwmvR5zBCYv1uyYc1kPIXgstvOOl4+rXA/PpRfVDdM4WFSveaJp4zT5mWgQyo4PVQpboXs0KjPxfEN+ott7ZeOxMNrZ9EAl0RERYNVpig=="
var CS_X509_PKCS1V1_5 = "mCxiVk2QcIO6/JQOR/pEhYa33hralrMLEZI+AtOKsbPbm5VTvMU32cA71iHaMOJyXYzFzDfvhFjS+ksZLfue7S3C24rwge0GY5PP4yxUhvA2yZWWrjmi7TF4b/BUICV2GF13tT9Iy34hbTGQcUgDaozmF1/332kInUM2Tc3pPTbMyB5YferSKnsLRQooV2kmnNZGHb0Gnt0GRxXm4H/Tm+rDsmbz7TbOJjzJefxaKgdVD2/gt5h7aDRQYvTuE59Lzn2b/jM4bMQxoR9owKcodt6duvQ79Ruefc+P2aWHTXBnOHNbvgUwpeqiAyW7sI98TRXzK4W3o1u4cakDJkbXJQ=="

// Js

// AES Encrypted Values:
var JS_AES_CFB_128 = "jHnsUoaMQ0qUeobWuVbKST9zwf1JsET2SaO4Srkvujg="
var JS_AES_CFB_192 = "HU9JZ/fXn4G/9n/5OELbeko3cIxUwy6w4dJT1IQuAJA="
var JS_AES_CFB_256 = "mMEED4K46PawCG4xqheUV3BJaQSEL1Wr5CZhdcHPwrw="
var JS_AES_CFB8_128 = "xdiYZL/Rz7X9mc+Bphe7sFab6c3slciabwSD0lvhnN4="
var JS_AES_CFB8_192 = "WtI13CkbWHwUduVFeIg9vMUCqJTAQZJjoOguAfuf23M="
var JS_AES_CFB8_256 = "0NGurdoTM9vj3gn+NQripcA02sF3P15LqClZuxvQB2Q="
var JS_AES_CBC_128 = "oGVboN1e39qClYUwURwDSQn7YNuoJOlfVVppfHwe6REUrdQLnwHVwivr/UQYX8Np"
var JS_AES_CBC_192 = "N9NtNJkz4hJQhAY3SupjlFQlaPPUuTuB58v8BU0kNjRmW/+TEnZpPtDY8iPpRAJd"
var JS_AES_CBC_256 = "C3O0fEMRV66OaeH/4iWDGwCf78pl4B6fxYga1b/9nS1jp9OQIEdrQEVFsCXMPsaQ"

// RSA Encrypted Values:
var JS_RSA_PKCS1V1_5 = "piGSxFiu3GSUQ0V/AJb2AjWRzfG2zDVbdEqxnAtT9HOKhDWrKsAP1eOtZ069OFx+cU5wyUjdC4FfaK+IBeJSOBB4DBZmKU4qKuRGQrJvg8a2FeeuU/9aXddw7SnZZ/ne+JGcjbXGSgrlnZXCk89q+OXWpwTlu0KT7WSI52OirijYxuVLxhGx4FCwNASoe4McWEEFlctIVPnDbBQXMDosXd0PqB4cLgXH99dd7u8oLa5dAPj+GVwN1reQKDVMeWnKc5kC2YczX5BsseG9/nK52/KKqf9hv7unpv5Gd04eZ6+rcyH1yruIv5zmaP8/xs0LNOCK3JeqbSz5ee6wHlbxSA=="
var JS_RSA_OAEP = "p5yonTWh2mfZyJr4/EYCUQg5tv5+3u7f28LFNPN4amg9Rb7RtQIhjjlS1SF+HHSTSkUy4mr4B2EkkxnEhYK9y2l+HzNhDmIvf4xhzgvaSjC8CagQXQd3V8FmW0m8BE8nXU+lIdK2YnvknZkSA9b2JQ7rlnEhoTQJ8UaQfpU8E7g+RWwEUDFpQOwsDm+NtrMxv45+HQm2NZ2Wa460FWzo94gbYsPNpH58mxVuu4xZYpKPN8D6fdqcEqDbhlnAI1vbKFIC4gez0Z9iZAHla/91XnrzeuuZy7vpZkfI04oj+1Qn32UN85QOrhaGbeDJpq5q6pKp06h8ur/GVwTogpGFow=="

// X509 Encrypted Values:
var JS_X509_PKCS1V1_5 = "EdW2aFSxcMacbzemF9mphiSP976c2XBPK4sW58y2eCt5fpsI1xjmElDELCQ/YNqkV3GNw1UAQoKA4xE7uVrvY/y/EyHKCScOLp8XBsqNQSl8PR/FV2xjVwhiU1MGCSiFbTd09a7zgYbujA62C08pIPXW35HtQCscbGVvdTO4HLYTvZVhphM2rzIqs7Aj1yxUZ8g75t0gXpuQBfoGUmGEO+6Me0wQN1NwcTZpaw4XShO/ov3JduzSQIV5/lkPvkqvb62rBGJCIrf5p1aGEZ54NlKFxrXRiDpOV+396pbXmIgW5b/kANke6BdTDZvnYhIZWcW7YY5a9ueyd5DtrBN+SQ=="
var JS_X509_OAEP = "wtH9eCIOP5w7eDvll80VNJ3Jg4yT46XdegJi0McKQIiqukHUDvZ1oAAb0XAStUGPD+Qb+FfoEDpnqcPgOYVoqmf1zuVIN9IEahiGE1NfxmJeAZO/iqK2NMTWUMYhxG9HD5WcxJnIiDVTtxLKfona6aX15Dvd7GCFRkIpG6d3p4GWall8selKC34VB+w1DTBTDeDmrYWvtK/TREk1RUI01slycbm1SDF2H46wnrqK7KQH+zVKj2a4Tz7qWOq+BxNfpaTjo2oDm3zWFASmwBf6ELmSO/D9wXjiTFrFNG4WouoYjwDPDeQLiF6kcJT4KdsJKcml+Y/GOXjGUS6uV9N8jg=="

// Py

// AES Encrypted Values:
var PY_AES_CFB8_128 = "TEW315W1NwDlI10Rdt3McEJlhoJTwQP83t3J+IiYml4="
var PY_AES_CFB8_192 = "3DCi763sbaxr4AKmCr1TA4UcGv1V91F2jMVnePscVmE="
var PY_AES_CFB8_256 = "N3XEOxMRIdKKqnimwuZ/BXTKL7RDqO+v5rowpcnWZFY="
var PY_AES_CBC_128 = "jRzMYVSJNWd6zfocjxaVSYluh/HW5LzSMnvbKcpFMoZHI9Bxufa2kLoRzyK7FXV5"
var PY_AES_CBC_192 = "44sMPzEaKnxQfVC4HeLtWXTlvsuEgz6hcvquMp0CHlENmYpdL8LrdDlkEtuVzp7c"
var PY_AES_CBC_256 = "+dE+0TugA+x8dsH1Lpl2NEmxJ/EpLy7Tu7MpkR8GKuff1RgYTLkxPCYR3ZBh4rVD"

// RSA Encrypted Values:
var PY_RSA_PKCS1V1_5 = "Q202+RdR7lAqitYWJnDxN26ikUVf00DMYa4GpFhLqKqLz0uwiUqlUcp1fKRT/LvJmmZqfO6bFdRs4AaYWZucrNkN+v+kWaaAGw3v+KVtihNe6zvXzs6vUCkKXVBsenZhVB/USL0w9eGzi+f3tr97HLHBo8zotZmAeVPEcuVifkLqKAz5v4VtvbU91qMnkcqUOh+wHA9TUH1F7fs3WUD2rPj31FVZayGWO3a8+3vMLyPbOLPtRi8mFtf3JFb0xmNLyXSWFyN8DOmHxIriTyD1iCLUVMi7UIBdZg1nSJXEREdDRGwxD6w314qa5FNDUCapbwhHylFWTsZp9l8JLXf4EA=="
var PY_RSA_OAEP = "mShEcmazguR6KHX7lI8zOLVjuaHpzfrjImk6jvIRjOxyaLjHTzXwadk3skUr3RLgZZYnBB/tINbVAfPLmmYwxyPqAQFoqjSxsJRvEdxLzTQy30SpoQieneSQ2ppO044jp863qRyf7UFTBFIWr+4EY/fowZHUslXdJQ6kC3lAIkTrNrick/Qxlb7y5vstJxy4VBtipCbRlH8ihSsw2ox539hHx/nS2qS90weone6VGsE57SvOjeqT41Gn97MptMmw8alQYOsHQEkzF7UBtUpcD09ucGvQaKKONk7xCscHCqvNwf0b5Io6Z3b6w9TsdL7ggHKkB2SpvjglRI4+xHFsUQ=="

// X509 Encrypted Values:
var PY_X509_PKCS1V1_5 = "JU6dFJpLXDRl+d7N2LszsDPqakoj9Gxus+XykaYyBgRMeOcoKO0ijbG9WxXw7GsYZ1nK4C7WZavrtLaMS4InLEBWO52t5AkzJmUfdAHgP6Of+3668QF5UaDPyjWH7/j2S4UEl2AbBNJgqGEBsVlrbADjkdt4sargLtUPDOhEk1mdHiyHdyNX1GBz2h84CbGQbkvhgKGHp9r46ZIismX1fgeIMWQR+Q9x8dZNDR0Z7/VQnuW0NbIe0RJaGLeiw767/ZAuTuC+VlQs/7wnmvHdUnEgu65oSEGNtCS7nj+2RLVRcfJA2YxkQ5vSNn9ldva/t7PsX4fINotYXX7YpwR3SA=="
var PY_X509_OAEP = "Noo4J2qGlyzZcz2rk1cw1Lr2R73a7meIRu/CZV21Su+4rla4rj2CrKgjQDQwFRQ8qRk5JJgE81YturfVMAymoA9Xp40PhM+tlhE/k2BZiJOXv3QIpLJCHq+2xqC3zxwXtlsGK0bq/dlB+Jr7Xqug3JXm9IpzAE/SN3lNS5NvlK8mRFYoz10atZlUuhkSh8CH9zZ+VCi/orSCRwbnq24uVB9cNS17u/VerNQIltSPfdOI3Ja7w0ddBE/V23Ub1dWmwCfRYXsnsJM37W5Awtht4an/+CDbdsbzNWc4GwlFtgkL1l0p04mwGYwq4aC3VYxFJFVDcThbzKTFRXdR9BipGA=="

func skip(t *testing.T, message string) {
	t.Skip(message)
}

func skipNotCompatible(t *testing.T) {
	// skip(t, "not compatible")
}

func skipNoCfb8Support(t *testing.T) {
	skip(t, "golang does not have CFB8 support, only CFB")
}

func skipOaepNotSupported(t *testing.T) {
	skip(t, "golang does not support OAEP comming from Js/Py")
}
