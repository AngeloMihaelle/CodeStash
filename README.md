# 🧰 CodeStash

**CodeStash** is a local-first, privacy-friendly CLI tool to manage, search, and execute code snippets right from your terminal. Organize your favorite shell commands, code templates, or reusable bits of logic — without relying on the internet or cloud storage.

No telemetry. No sync. Just your code, your way.

---

## 🚀 Features

✅ 100% local — no network calls, no telemetry  
✅ Store snippets in `JSON`, `TOML`, or `SQLite` (planned)  
✅ Organize with tags, languages, and descriptions  
✅ Fuzzy search to quickly find what you need  
✅ Print, copy, or execute shell snippets  
✅ Automatically track usage stats  
✅ Extend via local-only plugins  
✅ Shell integration for `bash`, `zsh`, `fish`

---

## 📦 Installation

> 🛠️ Go 1.20+ required

```bash
git clone https://github.com/AngeloMihaelle/CodeStash.git
cd codestash
go build -o codestash
mv codestash /usr/local/bin/  # or somewhere in your $PATH
````

---

## 🔧 Usage Overview

```bash
codestash <command> [flags]
```

| Command  | Description                                   |
| -------- | --------------------------------------------- |
| `add`    | Add a new snippet interactively or via flags  |
| `edit`   | Edit a snippet by ID or title                 |
| `list`   | List all snippets (with tag/language filters) |
| `use`    | Print/copy/execute a snippet                  |
| `delete` | Delete a snippet                              |
| `tag`    | Add or remove tags from a snippet             |
| `search` | Fuzzy search snippets                         |
| `exec`   | Execute a snippet directly                    |
| `stats`  | Show usage analytics                          |

---

## ✏️ Example

```bash
# Add a new snippet
codestash add

# List all Python snippets
codestash list --language python

# Use a snippet (copy to clipboard)
codestash use my-snippet-id --copy

# Execute a shell snippet
codestash exec deploy-script

# Search by keyword
codestash search docker

# View usage stats
codestash stats
```

---

## 🧠 Snippet Structure

All snippets follow the same structure regardless of the backend format:

### JSON example

```json
{
  "id": "abc123",
  "title": "Restart Docker",
  "code": "sudo systemctl restart docker",
  "tags": ["docker", "linux", "system"],
  "language": "shell",
  "description": "Restarts the Docker daemon on Linux.",
  "usage_count": 3,
  "last_used": "2025-07-03T21:14:00Z",
  "created_at": "2025-07-01T15:00:00Z"
}
```

### TOML example

```toml
id = "abc123"
title = "Restart Docker"
code = "sudo systemctl restart docker"
tags = ["docker", "linux", "system"]
language = "shell"
description = "Restarts the Docker daemon on Linux."
usage_count = 3
last_used = "2025-07-03T21:14:00Z"
created_at = "2025-07-01T15:00:00Z"
```

---

## 🗃️ Storage Options

Snippets are stored locally in either:

* `~/.codestash/snippets.json`
* `~/.codestash/snippets.toml`

Set your preferred format:

```bash
codestash config set format toml     # or json
```

Switching formats will automatically convert existing snippets to the new format (if implemented).

---

## 🔌 Local Plugins (Coming Soon)

Drop executable files in:

```bash
~/.codestash/plugins/
```

They’ll be auto-detected and can modify or extend snippet behavior (formatting, integration, etc). All offline.

---

## 📈 Stats Example

```bash
codestash stats
```

Output:

```
Top 5 Snippets:
1. git-push-fix.sh — used 12 times
2. restart-nginx — used 9 times
3. json-prettify — used 5 times
...
```

---

## 🐚 Shell Integration (Planned)

* Completions for `bash`, `zsh`, `fish`
* Auto-execute snippets from prompt
* Key bindings for inserting snippets

---

## 📁 Project Structure

```
codestash/
├── cmd/                # Subcommands (add, list, delete, etc.)
├── internal/
│   ├── snippet/        # Core snippet logic
│   └── store/          # JSON and TOML backends
├── data/               # snippets.json or snippets.toml
└── main.go             # CLI entrypoint
```

---

## 🛡 License

Licensed under the MIT License. See `LICENSE`.

---

## 💬 Contributing

PRs, issues, and ideas are welcome! See `CONTRIBUTING.md` for guidelines (soon™).

---

## ❤️ Why?

I wanted a faster, offline alternative to snippet managers, built for developers who likes working on the terminal and want full control of their data.

> Made with Go and grit. 

