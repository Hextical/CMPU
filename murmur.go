/*

Source: https://github.com/modmuss50/CAV2/blob/master/murmur.go

-------------------------------------------------------------------------------

MIT License

Copyright (c) 2018 Modmuss50

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"io/ioutil"

	"github.com/aviddiviner/go-murmur"
)

func GetByteArrayHash(bytes []byte) int {
	return int(murmur.MurmurHash2(computeNormalizedArray(bytes), 1))
}

func GetFileHash(file string) (int, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	result := GetByteArrayHash(bytes)
	return result, nil
}

func computeNormalizedArray(bytes []byte) []byte {
	var newArray []byte
	for _, byte := range bytes {
		if !isWhitespaceCharacter(byte) {
			newArray = append(newArray, byte)
		}
	}
	return newArray
}

func isWhitespaceCharacter(b byte) bool {
	return b == 9 || b == 10 || b == 13 || b == 32
}
