package factionhash

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FLFACTIONHASH_POLYNOMIAL = 0x1021
)

type hasher struct {
	table [256]uint32
}

// Function for calculating the Freelancer data nickname hash.
// Algorithm from flhash.exe by sherlog@t-online.de (2003-06-11)
func (h *hasher) RawHash(data []byte) uint32 {
	var hash uint32 = 0xFFFF
	for _, b := range data {
		hash = (hash >> 8) ^ h.table[byte(hash&0xFF)^b]
	}
	return hash
}

// NicknameHasher implements the hashing algorithm used by item, base, etc. nicknames
type NicknameHasher struct {
	hasher
}

type HashCode int

func (h HashCode) ToIntStr() string {
	return fmt.Sprintf("%d", int32(h))
}

func (h HashCode) ToUintStr() string {
	return strconv.FormatUint(uint64(int(h)), 10)
}

func (h HashCode) ToHexStr() string {
	return fmt.Sprintf("%x", int(h))
}

func (h *NicknameHasher) Hash(name string) HashCode {
	bytes := []byte(strings.ToLower(name))
	hash := h.RawHash(bytes)
	// hash = (hash >> (physicalBits - logicalBits)) | 0x80000000
	return HashCode(hash)
}

func NewHasher() *NicknameHasher {
	h := NicknameHasher{}
	NotSimplePopulateTable(FLFACTIONHASH_POLYNOMIAL, &h.table)

	return &h
}

var nick = NewHasher()

func HashFaction(name string) HashCode {
	return nick.Hash(name)
}

func NotSimplePopulateTable(poly uint32, t *[256]uint32) {
	for i := 0; i < 256; i++ {
		crc := uint32(i) << 8
		for j := 0; j < 8; j++ {
			crc <<= 1

			if crc&0x10000 != 0 {
				crc = (crc ^ poly) & 0xFFFF
			}
		}
		t[i] = crc
	}
}
