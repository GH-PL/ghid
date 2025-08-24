# Ghid — Golang Hash IDentifier

**Ghid** is a lightweight CLI tool for identifying the type of a given hash string using Go.

---

## Usage

### Commands

| Command | Description |
|--------|-------------|
| `ghid samples [hash name]` | Show sample hashes. |
| `ghid list` | Show the full list of supported hash types. |

### Flags

| Command | Description |
|--------|-------------|
| `ghid <hash>` | Show only the most common hash type. |
| `ghid -s <hash>` or `ghid --short <hash>` | Show only name references. |
| `ghid -c <hash>` or `ghid --hashcat-only <hash>` | Show only HashCat references. |
| `ghid -j <hash>` or `ghid --john-only <hash>` | Show only John the Ripper references. |
| `ghid -e <hash>` or `ghid --extended <hash>` | List all possible hash algorithms, including those using salt. |
| `ghid -v` or `ghid --version` | Show version information. |
| `ghid -n` or `ghid --no-color [command]` | Disable color output. |

---

## Examples

```bash
# Identify the most likely hash type
ghid 5f4dcc3b5aa765d61d8327deb882cf99

# Get only short name references
ghid -s 5f4dcc3b5aa765d61d8327deb882cf99

# List all supported hashes
ghid list

# Show samples for a given hash type
ghid samples md5
