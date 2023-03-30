package account

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPrivateKeyByKeystore(fileKeystore, password string) (privKey, address string) {
	// privKey, address, err := KeystoreToPrivateKey("UTC--2017-11-21T05-46-23.555205600Z--6e60f5243e1a3f0be3f407b5afe9e5395ee82aa2", "123456789")
	privKey, address, err := KeystoreToPrivateKey(fileKeystore, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("privKey:%s\naddress:%s\n", privKey, address)
	return privKey, address
}

func KeystoreToPrivateKey(privateKeyFile, password string) (string, string, error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {

		return "", "", err

	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}
