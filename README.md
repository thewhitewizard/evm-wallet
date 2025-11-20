# 🔐 EVM Wallet Generator

A secure and easy-to-use command-line tool for generating Ethereum wallets with support for keystore files and BIP39 mnemonic phrases.

## ✨ Features

- 🔑 **Secure Key Generation**: Uses cryptographically secure random number generation
- 📝 **BIP39 Mnemonic Support**: Generate 24-word recovery phrases
- 💾 **Keystore Format**: Export wallets in standard Ethereum keystore format
- 🔒 **Password Protection**: Secure keystore files with password encryption
- 🎨 **Beautiful CLI**: Clean, emoji-enhanced command-line interface
- ⚡ **Fast & Lightweight**: Minimal dependencies, quick execution

## 📦 Installation

Download the latest release from the [Releases](https://github.com/thewhitewizard/evm-wallet/releases) page and extract the binary for your platform.

## 🚀 Usage

### Basic Wallet Generation

Generate a new Ethereum wallet with all information displayed:

```bash
./evm-wallet
```

This will output:
- 🔑 Private key (hexadecimal)
- 🔓 Public key (hexadecimal)
- 📍 Ethereum address
- 📝 BIP39 mnemonic phrase (24 words)

### Generate Keystore Only

Generate only a keystore file without displaying sensitive information:

```bash
./evm-wallet --keystore
```

You'll be prompted to enter and confirm a password to encrypt the keystore file.

### Insecure Mode (Testing Only)

Generate a keystore without password protection (⚠️ **use only for testing**):

```bash
./evm-wallet --keystore --insecure
```

## 📋 Command-Line Options

| Flag | Description |
|------|-------------|
| `--keystore` | Generate wallet as keystore format only (hides sensitive info) |
| `--insecure` | Generate keystore without password (⚠️ insecure, testing only) |

## 📁 Output

### Standard Mode

When running without `--keystore`, the tool displays:
- Private key in hexadecimal format
- Public key in hexadecimal format
- Ethereum address (checksummed)
- BIP39 mnemonic phrase

### Keystore Mode

When using `--keystore`, the tool:
- Creates a `keystore/` directory in the current folder
- Generates an encrypted keystore file
- Displays the file path and account address

The keystore file is compatible with:
- Geth
- MyEtherWallet
- MetaMask (via import)
- Other Ethereum clients

## ⚠️ Security Warnings

1. **Never share your private key or mnemonic phrase** with anyone
2. **Store your keystore password securely** - if you lose it, you cannot recover your wallet
3. **Backup your mnemonic phrase** in a secure location
4. **Never use `--insecure` mode** for production wallets
5. **Keep your keystore files secure** - they contain encrypted private keys

---

**Made with ❤️ for the Ethereum community**

