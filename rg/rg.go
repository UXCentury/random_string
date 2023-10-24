package rg

import (
	"crypto/rand"
	"math/big"
)

// Options type for options
type Options struct {
	Length   int  `default:"8"`
	Alpha    bool `default:"true"`
	Numeric  bool `default:"true"`
	Special  bool `default:"false"`
	UseUpper bool `default:"true"`
	UseLower bool `default:"true"`
}

// GenerateRandomString with options for length and character set to use. options are length, alpha, numeric, special, useUpper, useLower
func GenerateRandomString(options Options) (string, error) {
	// generate random string by provided options
	// Define character sets we can use it also as global variable
	alphaCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numericCharset := "0123456789"
	specialCharset := "!@#$%^&*()_+[]{}|;':,.<>?"

	charset := ""
	if options.Alpha {
		if options.UseUpper {
			charset += alphaCharset[:26]
		}
		if options.UseLower {
			charset += alphaCharset[26:]
		}
	}
	if options.Numeric {
		charset += numericCharset
	}
	if options.Special {
		charset += specialCharset
	}

	// Calculate the length of the charset
	charsetLength := big.NewInt(int64(len(charset)))

	// Generate the random string with crypto/rand for secure random numbers
	result := make([]byte, options.Length)
	for i := 0; i < options.Length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}
