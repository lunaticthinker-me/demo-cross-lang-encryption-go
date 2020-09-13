package democrypt

import (
	"math/rand"
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
var CS_RSA = "r97MgnloMoX5K8LDxXLF7znhuk4vtjl9SlHaoQpQzX54jVFWQJZDIij7XUBLVuSu9CY6QjwZWG1DWYIPq7UfQkl7F0HQf2b1TKmnk4RfC9BMz5z3A8nMPkR24Ivt6gFh8yLG2gMdLErAuNutwDgZTVYLDdL3HX1XAvl/mN2cXV/NH+aLyWoLzFVcgokZWG+8uhICSWEvsgZL2mrjR+5Lhwb8B4VYvd0UqO/GQ+uLGUJHAe1U9bj2VGJ/z7WwPzruQjQmfr0earugokTv+F8zxjZB936GDtKJYFEXFqKP6Z15XwiSaYYucaBdBHkU5lFycau0PZyZCv9jqg6vMSF6Hw=="

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
var JS_RSA = "TcMeyv85OXgVkPhaK4qPC8ntrxMHuEAIzFEgSie17xYMfdIUxRm+TzpGPd6zNAEmEUTckyFQk5fziNdxWEOmGBnyQDObMAVrL+zXqC3HpFnXc+VxwP+suu58+7Ms3fGcKAWdixelqykMq+ewg/ESfmzaW3XksEalfuPd505ft399vcTX/bd8PdtI/QyXdHZeW1xyiclWhu6zAVZwI3VNJYusxE5ogpQetqK38fszzV46LrXnW2elhME0nsz5nNeZrQIhpLvU9ckBm6EvJmguHFeE/esnR2bueOr9gmGyoQJVzdyRZUU8/bNIpHTXhhQ+14QIs/eAtfCiCDxzyjamWg=="

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
var PY_RSA = "CYg15xBux9RIs4vhc7ylqcCSgdO4kJ0kds8TDJkJllB3RKDPSkxGLwBVZGiGqNKPidqlEymxSWCdAwxZsW2xi7EAXC4NCSkczeMMBdfnVaV7ZvNDewMB7WRo1j+6wj303ufavoDyqyO+SjJfg/K08pvS6hSTzNkSj9QvEFuaeLv47xmF2AV2dTil+alcLkK8DvWA5K8SY9LZMPqE1eLRT60EeuUOdhgGuUfd1ZHCXyZ5UKbxcpTJD5IWFpm+ShY0xi0sqBIQotfbHyJpauio+2uTZ0kbS5q4hIWN3aQSBndUV9hge7MNmZGRhQnHzivQ2HPDgJwQ+Wbri/P4EdW+OQ=="

// X509 Encrypted Values:
var PY_X509 = "vHOnktfekirpxRsXh5ffJjpy9IabBoHqeXcKhcX3saAQPbfbUNUHqSSUhFSDnF1/YKvZCGs9rUFjKyyrEPz4RCFFLB6leNq8y5eERJl6R8i987ZUkMBUmKIMLwMn8oZgEdTIn9VBdFP0/z44UQW5YKCadtQHik2HDVRIxYF4SO+q2sI5HoA/cGi9Jezbz7LkAvAXvom2EBNnYOO/6OCM8RewfR4poMBIlVdleNbxzmHFP4AxWYfgBw6bJLCjlKhP9gwFFObRsywxUHTZ6cxECXHTCMsmT5o+p4f8IAjkOxTkE/7bVdEK51au7W4VlUTRtBEB/T9nxM8cisIyGFGbqw=="
