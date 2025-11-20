package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/term"
)

const (
	keystoreDirPerm = 0700
)

func main() {
	keystoreOnly := flag.Bool("keystore", false, "Generate as keystore format")
	insecure := flag.Bool("insecure", false, "Generate keystore without password (insecure)")

	flag.Parse()

	fmt.Println("🔐 Generating secure Ethereum wallet...")
	fmt.Println()

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(fmt.Sprintf("❌ Error generating private key: %v", err))
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		panic("❌ Error converting public key")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !*keystoreOnly {
		mnemonic, err := bip39.NewMnemonic(privateKeyBytes)
		if err != nil {
			panic(fmt.Sprintf("❌ Error generating mnemonic phrase: %v", err))
		}

		fmt.Println("✨ === Ethereum Wallet Generated ===")
		fmt.Println()
		fmt.Printf("🔑 Private Key (hex): %s\n", hexutil.Encode(privateKeyBytes))
		fmt.Printf("🔓 Public Key (hex): %s\n", hexutil.Encode(publicKeyBytes))
		fmt.Printf("📍 Address: %s\n", address.Hex())
		fmt.Println()
		fmt.Println("📝 === BIP39 Mnemonic Phrase ===")
		fmt.Println(mnemonic)
		fmt.Println()
		fmt.Println("⚠️  WARNING: Keep this information secure and never share it!")
	} else {
		var password string

		if *insecure {
			password = ""

			fmt.Println("⚠️  Insecure mode enabled: keystore will be generated without password!")
			fmt.Println()
		} else {
			fmt.Print("🔒 Enter password for keystore: ")

			passwordBytes, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				panic(fmt.Sprintf("❌ Error reading password: %v", err))
			}

			fmt.Println()

			fmt.Print("🔒 Confirm password: ")

			passwordConfirm, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				panic(fmt.Sprintf("❌ Error reading password confirmation: %v", err))
			}

			fmt.Println()

			if string(passwordBytes) != string(passwordConfirm) {
				panic("❌ Passwords do not match!")
			}

			password = string(passwordBytes)
		}

		keystoreDir := "./keystore"
		if err := os.MkdirAll(keystoreDir, os.FileMode(keystoreDirPerm)); err != nil {
			panic(fmt.Sprintf("❌ Error creating keystore directory: %v", err))
		}

		fmt.Println("💾 Creating keystore file...")
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

		account, err := ks.ImportECDSA(privateKey, password)
		if err != nil {
			panic(fmt.Sprintf("❌ Error creating keystore: %v", err))
		}

		files, err := os.ReadDir(keystoreDir)
		if err != nil {
			panic(fmt.Sprintf("❌ Error reading keystore directory: %v", err))
		}

		var keystoreFile string

		for _, file := range files {
			if !file.IsDir() {
				keystoreFile = filepath.Join(keystoreDir, file.Name())
				break
			}
		}

		fmt.Println()
		fmt.Println("✅ === Keystore Generated ===")
		fmt.Printf("📁 Keystore file: %s\n", keystoreFile)
		fmt.Printf("📍 Account address: %s\n", account.Address.Hex())
		fmt.Println()

		if *insecure {
			fmt.Println("⚠️  WARNING: Keystore was generated without password (insecure mode)!")
			fmt.Println("⚠️  It is extremely vulnerable. Use only for testing!")
		} else {
			fmt.Println("⚠️  WARNING: Keep this information secure and never share it!")
			fmt.Println("⚠️  The keystore file is protected by your password. Don't lose it!")
		}
	}
}
