安装

```bash
brew tap ethereum/ethereum
brew install ethereum
```

```bash
geth account new //创建账号，会提示输入密码
```


创建账号后，反解出私钥

```go
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

func main() {
	privKey, address, err := KeystoreToPrivateKey("UTC--2017-11-21T05-46-23.555205600Z--6e60f5243e1a3f0be3f407b5afe9e5395ee82aa2", "123456789")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("privKey:%s\naddress:%s\n", privKey, address)
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
```

