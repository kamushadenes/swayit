package hash

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func HashAll(text string) []*Hashed {
	s := []byte(text)
	var hasheds []*Hashed


	hasheds = append(hasheds, &Hashed{Name:"MD4", Value: fmt.Sprintf("%x", md4.New().Sum(s))})

	hasheds = append(hasheds, &Hashed{Name:"MD5", Value: fmt.Sprintf("%x", md5.Sum(s))})

	hasheds = append(hasheds, &Hashed{Name:"SHA-1", Value: fmt.Sprintf("%x", sha1.Sum(s))})

	hasheds = append(hasheds, &Hashed{Name:"SHA-256", Value: fmt.Sprintf("%x", sha256.Sum256(s))})

	hasheds = append(hasheds, &Hashed{Name:"SHA-512", Value: fmt.Sprintf("%x", sha512.Sum512(s))})

	hasheds = append(hasheds, &Hashed{Name:"RIPEMD160", Value: fmt.Sprintf("%x", ripemd160.New().Sum(s))})

	hasheds = append(hasheds, &Hashed{Name:"SHA-3 256", Value: fmt.Sprintf("%x", sha3.Sum256(s))})
	
	hasheds = append(hasheds, &Hashed{Name:"SHA-3 512", Value: fmt.Sprintf("%x", sha3.Sum512(s))})
	
	hasheds = append(hasheds, &Hashed{Name:"Blake2b 256", Value: fmt.Sprintf("%x", blake2b.Sum256(s))})
	
	hasheds = append(hasheds, &Hashed{Name:"Blake2b 512", Value: fmt.Sprintf("%x", blake2b.Sum512(s))})
	
	hasheds = append(hasheds, &Hashed{Name:"Blake2s 256", Value: fmt.Sprintf("%x", blake2s.Sum256(s))})

	bcryptHash, _ := bcrypt.GenerateFromPassword(s, 1)
	hasheds = append(hasheds, &Hashed{Name:"BCrypt", Value: fmt.Sprintf("%x", bcryptHash)})
	
	return hasheds
}
