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
var CS_AES_CFB8_128 = "B/c9vvIehMurMmOGDr2UK3BnLTdzJ6wtIRKCrc7HaBI="
var CS_AES_CFB8_192 = "r4tz8b2BB0WLk0Ib3Cgm5r/KltPJmG731ReaFcK56lQ="
var CS_AES_CFB8_256 = "o09jrVm34bTtP03UqtKCuJ+a6iI/JfIc0b7dOcHNk3A="
var CS_AES_CBC_128 = "oenNpLqtCkSwG2SQVc3uYLw4WJ/wtDZ1Dm8bPKjS7K8DXr/ZZfE0eCI46DkyozVj"
var CS_AES_CBC_192 = "kf9bDoBIVG2+pdx2IM7jcHhe818bX3b14L10CpyV8ZqbC0AdI5ZrWxucvfOz8LkK"
var CS_AES_CBC_256 = "i/Qsv1gJ2pNuvIPQjF5S6MAV6R07Blq7JX8hiSdwFNgKPNxnKUnUxQzfNshNZiYe"

// RSA Encrypted Values:
var CS_RSA_PKCS1V1_5 = "QAfP0mIDpc3ltgBDejTZobaA1UTpU9l/Pz8g6+FP18yJIW86Nz5Kt+2of1izXHOyox3S85uqY6aK0Mjft4NgVs4tJP6kU3oMdmhb7jUuhFWP+DgE4QVoCCjGyoH29Eb0rhyagitaAlBFJF7pZKMJp/QPHh0Ih06K82ZlwpAuLaQVvQ3GGEfTi5QBhPIzZc0Y61r8YlS35b7qcZzjQ+IjXlAizwN1TXxkJOrlWQVlYoyw0Fv2f0pTP1+dW4pkfXe6aGqfhDQX83JWfTam0b5Ad27szY0QPcA5NUPAgXvpHboS+eC4KXRTKSHeil/wDscJotCCACGUeBPS6DpmGwnRGw=="
var CS_RSA_OAEP = "hLaLqsg/+vciM4uPZVa3KGMsjALwCOhcBq0o06Q2rZqdFrEnnUXh1ZD5MInUQuU8HC0/na7as1s0hYRWfnKnsKC/jxIVgJ2+D2Nu7OzF/GQsXFY9yESNibVp63oOcpiSvS0UqYVRirGNL4INZ6jqZt8VbhRW1JgBno2Mxd8o9kUl4dxRzF+JR4j9Y0dLouVkSVC7hoYxfZ7kqWyG+ECG6E5BdWsk5fJqdsvpLi+qD62OgQvhjD5pHr6GtIbubcKxtI53yt/8qkVKmWLEuJNKEe8pD3IksRjpbzka8KzJ9UpCTeZ3OxzPLk3tm1obe6ny2t692/RiWr2zPZCJYbgmtg=="

// X509 Encrypted Values:
var CS_X509_OAEP = "N0rG61/96dy7vY/8yGubzMiA10Pz8H9rY12eJvpXyynnuJUXb6fP1CosuMb+qtgrRbkLAI3o+Tyk1bHe8bixaXkJuE7C39Yz6tmNsvWA/0KB1O6+VpB9pqY65l2xxr/YhlMznvPl4Qx9dboX9y03fG+xcvHfKsvjwtRVXw0KXUPcskIn9Uv9pVZYzrrtC+8s9xGWvWwrwH4zjtRjDhvTby8gTPdSgxXO4HuX/VTymeFAZUzZU64TdsGLpPyDTHWew3mWcN731r6WLpG4PkUaLZHDSA4QHexHib0io/UkermbdDOkXRnlwwGs9FZQw4Ozx17ceo32OIS5UGerUz6qJw=="
var CS_X509_PKCS1V1_5 = "s1MoudY8FSRkLCvECxhclLCLpChP85epbGtroCRcIxRMP0KgoB3R7cgpuP486TJWWKqhGhz+oWUh+ulqhXTd65aP5incSx9FVi+RcovnVR0a5XGxO8HhCRY62RF/YIbUOwnBwHLUVFQgRXIWeB3M+H5eCDWvIWmDS+n8KRicHslPiHecgSc0+0pcWe7V5hF8bZRzFuJzLoDuPIsvUrIrCOsHLh73syxw2m7FQHf0eJMX6s+AdvnnDLvxY3W9ol73RFddeHH3cAQzBg+PhjW2I5bymev9d8GVTanITdVnX9PmF6sq8Mb6ETZg36y07+tKvQtEK5f+QC0zDRYlUSeThw=="

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
