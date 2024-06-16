package handlers

import (
	"hash/crc32"
	"os"
)

var logo, _ = os.ReadFile("logo.jpg")
var targetCRC32 = crc32.ChecksumIEEE(logo)
var flag = "SgffCTF{c0ll1s10n_h4ck3r-_-}"
