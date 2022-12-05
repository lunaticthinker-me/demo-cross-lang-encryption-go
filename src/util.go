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
var CS_X509 = "Kr3bWMPlCo/6ZqAebg3fOEKykYxMt/8ScWx8GOWy/1XZmFBWbLuQncd/qfwIQaZogPQkzO95WLXSq/+1CJfT2Jkk4yHfnkwrUmgT8+I84B7HCLx62uIIIjhfTtcN1J8SDdVE0AZx9BukHS7aBN218IH6p5qlqpVzjw/gI4tgmd2Ol/wz7Vzovnb4dfDwj4LgMHXDeXKhf7wW7DrAUGGADCOqYG9A1Iy6CBshvjD5MapDykTQ2gVZFQC2mrZnn/D/0SHYq3kCGLrCCcjBrAVU5ULOuPMsFsLbdCC0DZP3qCDo8Nc4BfQQy4jzXk1XwxIcqUXmmoUgtFrB8VPFUNUJvA=="

// AES Encrypted Values:
var JS_AES_CFB_128 = "z11VkLYh2sxTyMSh4wr/VQyNGS0IuJelSVJKSH1piYc="
var JS_AES_CFB_192 = "4+eUuv4Hr03hXXVz2utBVZmn9fT5E/N95Qq4qIuHWdw="
var JS_AES_CFB_256 = "1cXCXZZxT4sZc2FzzLGyrE7OZ32GxWLmv2irxJ4KzqI="
var JS_AES_CFB8_128 = "1ZzwOUXOhhbCrxOXQvnO8lA0he+bZo7KCfu6J1w75MI="
var JS_AES_CFB8_192 = "EJ9WAk9j4i7j9Dm0xiRh8x5r7/KZHLXO/Y34wC7m7jE="
var JS_AES_CFB8_256 = "Y7i1V8qehMsUs0Oo8NDWOqT7rSMr1rRY+gP6Pa+7peU="
var JS_AES_CBC_128 = "ysMa13YpdgO0zsRmZlKUF7uarZsCuDutYdjjgnntSRm8dAZWj4Nm2yo4bRGAZrWq"
var JS_AES_CBC_192 = "d52n/ivp5Vakrx30+IVn7kUjVc4gzCl1dZDa53+8GmvpgrH8GNQcnxJGPW9QkZd0"
var JS_AES_CBC_256 = "RK3HQA6FhecHKMGZIMR+jOs6VnKxLT2Y2YYY/IEPxeGp687czg6iGDomrH/K9Gsh"

// RSA Encrypted Values:
var JS_RSA_OAEP = "by+ECCKCT8YzHJSmZLZwJKDY01PpMVFOBHhbvZxIQju15Vbhjmg1U7UU+w4dPVGDuApOwY2sLsy+vcabsbqFBotBDQr0wNqPz6ZcjtjDbi5mXvytjd4YGFGDh40U2EiVrMI8lcfVQ/tgWqQ3o7IuOcllAaKFmWexKX+S6esMxkEwlWYQ/anU5wsYtG6E2KhxvI9uL0H/8ycWge0tAl0Vc1p0jEwjmAOozfz0Lzo6Z63tBM+8qNpveRUSLwVlNDW/2RqNYInyfld5vthJuON2EAAFrOaD6NivRLutLkTm+QcIHjMXtNgkYi4t4X0uoziAa8Fq/VdPgHrq9IXxEqz3TA=="
var JS_RSA_PKCS1V1_5 = "Qfe9Nd+KEVFgikZeUVAXoopFbJ1/nOWFjKIkhgKt794lZUzv4EbgF7MbVuKzwiI5XnLamJ+5Ki96xNjPEo7WQNbF9AQVwXhhuwj83OLxB8Jps6zyshQn4FbQ0q7trboscHD7rMDMBsn+MPyHz1oYI1C/C9v8eUoCTc7UEo1UcAoQTid1oEigayrR2x15TQY5gLuaTRvpMkqbUUOI1sOgHr5QOsDkqJJWdeE82xhjOthVFUJ+gJrDTFcq7tqYI9FL++nAHjDq4tLS/by1Lif8ivZ3Q+4x2NWlYjUYpFwBT4FULBX58v/PRpr3E3Puh5rs47jHzGdLyMYMx8xm5tC8Ww=="

// X509 Encrypted Values:
var JS_X509 = "qfbx31aXmLDPflz+dBwSEowlFEkavSBq1yDtIA0sDR2/sYaEkdaPRdB43aU8Ng+OHwUFisHlYOlZeylt6XcbrbRrDzo/0fVA0bqY3atbr3r8If7DPGYeV/3rxrKxLrb4n/k6Q/Z/9sJAiZv3OTcC66UPt4UMgP/ObP0u1Crwr0oKJythuvIDXoi+/ALCqG8jWKLiNV+KnPcIFvas80Qlqzcnk0ylTGdaxfcfoR1gXXy2a1vISTAdx3DQD6gjoBhZ+V7s13RVu4Go+Mv5ei4uxKPf3YG4XuycEq4O23qUGtf7HvXEEPsyhIdxC4sF7UJSU004fNSIaG17dIjQRKqw1A=="

// AES Encrypted Values:
var PY_AES_CFB8_128 = "mllIL/0l03VULBmJZpLfHgnpEIgBY20C6eUDd/C4nf0="
var PY_AES_CFB8_192 = "BKf2TvLAX/XhhobBBI90h9TapCweKd27M3vk5OEUcEc="
var PY_AES_CFB8_256 = "0/2/ZETH0ItnFwq2HG8Bz+1yRyi8Qn43GMBBJgbWxyM="
var PY_AES_CBC_128 = "w+9OjMn4Z6YeCrzo7DmhTpL++Udp/vMU+0nulOMcN0JL7iwYIEzKl9iGCVgzCuRf"
var PY_AES_CBC_192 = "0PuhDzNNaDdFuIFYFN7wmt4ucCWNqYLs3mJdauJgOLx9GBH5zQ6BNQwdjd85I4Hd"
var PY_AES_CBC_256 = "BALwcIvh+HaYRkkbuNuUfBWSfUlElpJNOY5YX8AvvLA9bI2N+966RAKRZFG0FZuW"

// RSA Encrypted Values:
var PY_RSA_PKCS1V1_5 = "fGbqSNgYCioKV7Kz/RXsOqSLp2e1vUAGnOJ3MMCMoOP+tTLiz69hyJRoPYlZwVkepOWmnEM2i3IZzAfqA7FzICecsjjk/hJOC3D6I+D2L4PVXPyTBsLSxvu9va8F/KzEfUBh/ysuw+JeSUc7k6OAsMPdGZPa8cyGDagKCPrGonaCe8Jnqac+zf81/a5sIche2YOqvQnz3V8ESR3XYuS3WYIn7LHkIFvfb4oxDgN1jSpxSZNImdmmLgjN0pW1Vp/UGQQAhnGgJW2Hs3noQtswq0O0XM3hpHJGzV6NA7eKTUSe1X1BrcAI3cUqwlL0XNBVnsNoq/Mj/Z1/4DgvMl34yg=="
var PY_RSA_OAEP = "jbEtUjq4qFzMb1+1LZRO8AZi6UHqm2YYAQjEzYfM5XAYysSdPaFy9EtPFZmgM5ZxbNkThOSvRU2yw9tKUR46F/HRxtqCcnRnbNr+nZoGYRSZzqjyQQsNOGxLBG0X4K0TugvZAu3zU/nFbytEZX/OM/RkiagztXcTGWGWbvfRZ/wvLZTuVH1nAdI9r33PQGceUwplwJ8WFl3amMCC2OejitIXBrkgiHpuAz2IMfj/PM3HN2hykPCb7tTxNmUtc0050NMLaVHYQgxtmQSSCpUPaCz9kZxyjqU1+QYGWw03bgeA02D9tXz13h3j2E9a7CW+Ut6LBja6vOPqnu3oF/p9dw=="

// X509 Encrypted Values:
var PY_X509 = "vHOnktfekirpxRsXh5ffJjpy9IabBoHqeXcKhcX3saAQPbfbUNUHqSSUhFSDnF1/YKvZCGs9rUFjKyyrEPz4RCFFLB6leNq8y5eERJl6R8i987ZUkMBUmKIMLwMn8oZgEdTIn9VBdFP0/z44UQW5YKCadtQHik2HDVRIxYF4SO+q2sI5HoA/cGi9Jezbz7LkAvAXvom2EBNnYOO/6OCM8RewfR4poMBIlVdleNbxzmHFP4AxWYfgBw6bJLCjlKhP9gwFFObRsywxUHTZ6cxECXHTCMsmT5o+p4f8IAjkOxTkE/7bVdEK51au7W4VlUTRtBEB/T9nxM8cisIyGFGbqw=="

func skipNotCompatible(t *testing.T) {
	t.Skip("not compatible")
}
