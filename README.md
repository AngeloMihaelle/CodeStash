# 🧰 CodeStash

**CodeStash** is a local-first, privacy-friendly CLI tool to manage, search, and execute code snippets right from your terminal. Organize your favorite shell commands, code templates, or reusable bits of logic — without relying on the internet or cloud storage.

No telemetry. No sync. Just your code, your way.

---

## 🚀 Features

✅ 100% local — no network calls, no telemetry
✅ Store snippets in `JSON`
✅ Organize with tags, languages, and descriptions
✅ Fuzzy search to quickly find what you need
✅ Print, copy, or execute shell snippets
✅ Automatically track usage stats
🔲 Extend via local-only plugins (Planned)
🔲 Shell integration for `bash`, `zsh`, `fish` (Planned)
🔲 TOML and SQLite support (Planned)
🔲 macOS and Linux testing (Planned)

---

## 📦 Installation

> 🛠️ Go 1.20+ required

```bash
git clone [https://github.com/AngeloMihaelle/CodeStash.git](https://github.com/AngeloMihaelle/CodeStash.git)
cd codestash
go build -o codestash
mv codestash /usr/local/bin/  # or somewhere in your $PATH
````

-----

## 🔧 Usage Overview

```bash
codestash <command> [flags]
```

| Command | Description |
|---|---|
| `add` | Add a new snippet interactively |
| `edit` | Edit a snippet by ID or title |
| `list` | List all snippets (with language/tag filters) |
| `use` | Print, copy, or execute a snippet |
| `delete` | Delete a snippet |
| `search` | Fuzzy search snippets |
| `exec` | Execute a snippet directly |
| `copy` | Copy a snippet to the clipboard |
| `print` | Print a snippet to the terminal |
| `stats` | Show usage analytics |
| `config` | **Planned:** Configure CodeStash settings (e.g., storage format) |
| `tag` | **Planned:** Add or remove tags from a snippet |

-----

## 📖 Commands Documentation

### `codestash add`

Add a new snippet interactively. You will be prompted for the snippet's title, description, language, tags, whether it's executable, and the code content.

```bash
codestash add
```

**Example:**

```
📝 Title: My Docker Command
🧾 Description: Command to list all Docker containers.
💻 Language: shell
🏷️ Tags (comma separated): docker,containers
🚀 Is this snippet executable? (y/N): y
📋 Enter code (end with 'EOF' on a new line):
docker ps -a
EOF
✅ Snippet added successfully!
🚀 This snippet is marked as executable and can be run with 'codestash exec'
```

### `codestash copy [snippet-id-or-title]`

Copy a snippet's code to your system clipboard.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to copy.

<!-- end list -->

```bash
codestash copy my-snippet-title
codestash copy abc123
```

**Example:**

```bash
codestash copy "Restart Docker"
# 📋 Copied 'Restart Docker' to clipboard
```

### `codestash delete [snippet-id-or-title]`

Delete a snippet from your collection. By default, it will ask for confirmation.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to delete.

**Flags:**

  * `-f`, `--force`: Delete without asking for confirmation.

<!-- end list -->

```bash
codestash delete my-old-script
codestash delete abc123 --force
```

**Example:**

```bash
codestash delete "Unused Python Script"
# ⚠️  Are you sure you want to delete 'Unused Python Script'? [y/N]: y
# ✅ Deleted snippet 'Unused Python Script'
```

### `codestash edit [snippet-id-or-title]`

Edit an existing snippet. If no field flag is provided, it will launch an interactive editing session for all fields.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to edit.

**Flags:**

  * `-f`, `--field string`: Edit a specific field (e.g., `title`, `description`, `language`, `tags`, `executable`, `code`).

<!-- end list -->

```bash
codestash edit my-snippet-title
codestash edit abc123 --field code
```

**Example (Interactive):**

```bash
codestash edit "My Docker Command"
# 📝 Editing snippet: My Docker Command
# ─────────────────────────────────────
# Press Enter to keep current value, or type new value:
#
# 📝 Title [My Docker Command]: New Docker Command Title
# 🧾 Description [Command to list all Docker containers.]:
# 💻 Language [shell]:
# 🏷️ Tags [docker, containers]: docker, devops
# 🚀 Executable [Yes] (y/n):
# 📋 Edit code? (y/N): n
# ✅ Snippet 'New Docker Command Title' updated successfully!
```

**Example (Specific Field):**

```bash
codestash edit "New Docker Command Title" --field description
# 🧾 Current description: Command to list all Docker containers.
# 🧾 New description: Lists all Docker containers, including stopped ones.
# ✅ Snippet 'New Docker Command Title' updated successfully!
```

### `codestash exec [snippet-id-or-title]`

Execute a snippet directly. By default, the snippet must be marked as executable.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to execute.

**Flags:**

  * `-f`, `--force`: Force execution even if the snippet is not marked as executable.

<!-- end list -->

```bash
codestash exec my-executable-script
codestash exec another-script --force
```

**Example:**

```bash
codestash exec "Restart Docker"
# 🚀 Executing 'Restart Docker'...
# ─────────────────────────────────────
# (Output of 'sudo systemctl restart docker' will appear here)
```

### `codestash list`

List all stored snippets. You can filter by language or tag, and choose to show the full code content.

**Flags:**

  * `-l`, `--language string`: Filter snippets by a specific language.
  * `-t`, `--tag string`: Filter snippets by a specific tag.
  * `-e`, `--expanded`: Show the full code content for each listed snippet.

<!-- end list -->

```bash
codestash list
codestash list --language python
codestash list --tag git --expanded
```

**Example:**

```bash
codestash list --language shell
# 📚 Found 2 snippet(s):
#
# 🔹 ID: abc123
#    Title: Restart Docker
#    Language: shell
#    Tags: docker, linux, system
#    Description: Restarts the Docker daemon on Linux.
#    Executable: true
#    Used: 3 times
#
# 🔹 ID: def456
#    Title: Git Status
#    Language: shell
#    Tags: git, cli
#    Description: Show git status.
#    Executable: false
#    Used: 1 time
```

### `codestash print [snippet-id-or-title]`

Print a snippet's details and code content to the terminal.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to print.

<!-- end list -->

```bash
codestash print my-snippet-id
codestash print "My Python Template"
```

**Example:**

```bash
codestash print "Restart Docker"
# 📄 Restart Docker
# 📝 Restarts the Docker daemon on Linux.
# ─────────────────────────────────────
# sudo systemctl restart docker
# ─────────────────────────────────────
```

### `codestash search [query]`

Search snippets by title, description, tags, language, or code content using a fuzzy search.

**Arguments:**

  * `query`: The search term.

**Flags:**

  * `-e`, `--expanded`: Show full code content for each matching snippet.
  * `-x`, `--executable`: Show only executable snippets among the search results.

<!-- end list -->

```bash
codestash search docker
codestash search "create table" -e -x
```

**Example:**

```bash
codestash search "list files"
# 🔍 Found 1 snippet(s) matching 'list files':
#
# 🔹 ID: ghi789
#    Title: List Files in Directory
#    Language: bash
#    Tags: filesystem, cli
#    Description: Lists files in the current directory.
#    Used: 0 times
#    📄 Executable: No
#    Preview: ls -la
```

### `codestash stats`

Show usage statistics and analytics for your snippets.

**Flags:**

  * `-d`, `--detailed`: Show detailed statistics, including unused snippets and full language/tag breakdowns.

<!-- end list -->

```bash
codestash stats
codestash stats --detailed
```

**Example:**

```bash
codestash stats
# 📊 CodeStash Statistics
# ═══════════════════════════════════════
#
# 📚 Total Snippets: 5
# 🚀 Executable Snippets: 2
# 📋 Non-executable Snippets: 3
# 📈 Total Usage: 15 times
# 📊 Average Usage: 3.0 times per snippet
#
# 🏆 Top 5 Most Used Snippets:
# ───────────────────────────────────────
# 1. git-push-fix.sh — used 12 times (last used: 3 hours ago)
# 2. restart-nginx — used 9 times (last used: 1 day ago)
# ...
#
# 💻 Top Languages:
# ───────────────────────────────────────
# 1. shell — 3 snippets (60.0%)
# 2. python — 1 snippet (20.0%)
# ...
#
# 🏷️  Top Tags:
# ───────────────────────────────────────
# 1. docker — 2 snippets
# 2. linux — 2 snippets
# ...
#
# 🆕 Recently Created:
# ───────────────────────────────────────
# • My New Snippet — created just now
# • Another Recent One — created 5 minutes ago
#
# 🕒 Recently Used:
# ───────────────────────────────────────
# • My Favorite Snippet — used just now
# • Docker Restart — used 2 minutes ago
```

### `codestash use [snippet-id-or-title]`

A versatile command to print, copy, or execute a snippet. If no action flag is provided, it defaults to printing the snippet.

**Arguments:**

  * `snippet-id-or-title`: The ID or title of the snippet to use.

**Flags:**

  * `-c`, `--copy`: Copy the snippet's code to the clipboard.
  * `-x`, `--execute`: Execute the snippet.
  * `-f`, `--force`: Force execution even if the snippet is not marked as executable (only applies with `-x`).

<!-- end list -->

```bash
codestash use my-template
codestash use my-script -c
codestash use my-shell-command -x
codestash use non-exec-script -x -f
```

**Example:**

```bash
codestash use "Daily Report Script" -c
# 📋 Copied 'Daily Report Script' to clipboard

codestash use "Setup Environment" -x
# 🚀 Executing 'Setup Environment'...
# ─────────────────────────────────────
# (Execution output)
```

Currently multi-lined commands aren't  supported.

-----

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
  "created_at": "2025-07-01T15:00:00Z",
  "executable": true
}
```

### TOML example (Planned)

```toml
# NOTE: TOML support is planned for future releases.
id = "abc123"
title = "Restart Docker"
code = "sudo systemctl restart docker"
tags = ["docker", "linux", "system"]
language = "shell"
description = "Restarts the Docker daemon on Linux."
usage_count = 3
last_used = "2025-07-03T21:14:00Z"
created_at = "2025-07-01T15:00:00Z"
executable = true
```

-----

## 🗃️ Storage Options

Snippets are currently stored locally in:

  * `~/.codestash/snippets.json`

**Planned:**

  * `~/.codestash/snippets.toml`
  * `~/.codestash/snippets.sqlite`

Set your preferred format:

```bash
# NOTE: This command is planned for future releases.
codestash config set format toml     # or json, or sqlite
```

Switching formats will automatically convert existing snippets to the new format (if implemented).


-----

## 🔌 Local Plugins (Planned)

Drop executable files in:

```bash
~/.codestash/plugins/
```

They’ll be auto-detected and can modify or extend snippet behavior (formatting, integration, etc). All offline.

-----

## 📈 Stats Example

```bash
codestash stats
```

Output:

```
📊 CodeStash Statistics
═══════════════════════════════════════

📚 Total Snippets: 5
🚀 Executable Snippets: 2
📋 Non-executable Snippets: 3
📈 Total Usage: 15 times
📊 Average Usage: 3.0 times per snippet

🏆 Top 5 Most Used Snippets:
───────────────────────────────────────
1. git-push-fix.sh — used 12 times (last used: 3 hours ago)
2. restart-nginx — used 9 times (last used: 1 day ago)
3. json-prettify — used 5 times (last used: 2 days ago)
4. clean-docker — used 2 times (last used: 1 week ago)
5. hello-world — used 1 time (last used: 1 month ago)

💻 Top Languages:
───────────────────────────────────────
1. shell — 3 snippets (60.0%)
2. python — 1 snippet (20.0%)
3. javascript — 1 snippet (20.0%)

🏷️  Top Tags:
───────────────────────────────────────
1. docker — 2 snippets
2. linux — 2 snippets
3. git — 1 snippet
4. system — 1 snippet
5. utility — 1 snippet

🆕 Recently Created:
───────────────────────────────────────
• My New Snippet — created just now
• Another Recent One — created 5 minutes ago
• Daily Report Script — created 1 hour ago

🕒 Recently Used:
───────────────────────────────────────
• My Favorite Snippet — used just now
• Docker Restart — used 2 minutes ago
• Git Fix — used 1 hour ago
```

-----

## 🐚 Shell Integration (Planned)

  * Completions for `bash`, `zsh`, `fish`
  * Auto-execute snippets from prompt
  * Key bindings for inserting snippets

-----

## 📁 Project Structure

```
codestash/
├── cmd/                # Subcommands (add, list, delete, etc.)
├── internal/
│   ├── snippet/        # Core snippet logic (model, ID generation)
│   └── store/          # JSON backend for snippet storage
├── data/               # Default location for snippets.json
└── main.go             # CLI entrypoint
```


-----

## 🛡 License

Licensed under the MIT License. See `LICENSE`.

-----

## 💬 Contributing

PRs, issues, and ideas are welcome\! See `CONTRIBUTING.md` for guidelines (soon™).

-----

## ❤️ Why?

I wanted a faster, offline alternative to snippet managers, built for developers who likes working on the terminal and want full control of their data.

> Made with Go and grit.
