# Ghid — Golang Hash IDentifier

**Ghid** is a lightweight CLI tool for identifying the type of a given hash string using Go.

---

## Usage

###  Commands

| Command | Description |
|--------|-------------|
| `ghid detect` | Identify the most probable hash type. |
| `ghid list` | Show the full list of supported hash types. |
| `ghid samples [hash name]` | Show sample hashes for the given hash type. |
| `ghid decode` | Try to decode hashes using a dictionary attack. |


---

###  Flags (general)

| Flag | Description |
|------|-------------|
| `-s`, `--short` | Show only name references. |
| `-c`, `--hashcat-only` | Show only Hashcat references. |
| `-j`, `--john-only` | Show only John the Ripper references. |
| `-e`, `--extended` | List all possible hash algorithms, including salted ones. |
| `-v`, `--version` | Show version information. |
| `-n`, `--no-color` | Disable color output. |

---

###  Flags (for `decode` command)

| Flag | Description |
|------|-------------|
| `-r`, `--read <file>` | Read input file containing hashes (format: `user:hash`). |
| `-w`, `--writer <file>` | Output file to save results (default: `decrypt.txt`). |
| `-d`, `--dictionary <file>` | Dictionary file to use for cracking. |
| `-t`, `--hash-type <type>` | Hash type to use (`md5`, `sha1`, `sha256`). Default: `md5`. |

---

## Examples

```bash
# Identify the most likely hash type
ghid 5f4dcc3b5aa765d61d8327deb882cf99

# Get only short name references
ghid -s 5f4dcc3b5aa765d61d8327deb882cf99

# List all supported hash types
ghid list

# Show samples for a given hash type
ghid samples sha256

# Decode hashes using a dictionary attack (md5)
ghid decode -r hashes.txt -w out.txt -t md5 -d words.txt
