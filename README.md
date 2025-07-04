# 🧰 CodeStash

> Your local-first code snippet manager

 CodeStash is a local-first, privacy-friendly CLI tool to manage, search, and execute code snippets right from your terminal. Organize your favorite shell commands, code templates, or reusable bits of logic — without relying on the internet or cloud storage.

## ✨ Features

- **Local-First**: All data stored locally in JSON format
- **Cross-Platform**: Works on macOS, Linux, and Windows
- **Executable Snippets**: Mark snippets as executable and run them directly
- **Smart Search**: Search by title, description, tags, language, or code content
- **Usage Analytics**: Track snippet usage with detailed statistics
- **Clipboard Integration**: Copy snippets to clipboard with ease
- **Tagging System**: Organize snippets with custom tags
- **Language Support**: Syntax highlighting and language-specific execution
- **Multi-line Commands**: Support for complex scripts and multi-line code blocks

## 🚀 Installation

```bash
git clone https://github.com/AngeloMihaelle/CodeStash.git
cd codestash
go build -o codestash
mv codestash /usr/local/bin/  # or somewhere in your $PATH
```

## 📖 Usage

### Adding Snippets

Create a new snippet interactively:

```bash
codestash add
```

Example workflow:
```
📝 Title: Git force push safely
🧾 Description: Force push with lease to avoid overwriting others' work
💻 Language: bash
🏷️ Tags (comma separated): git, safety, push
🚀 Is this snippet executable? (y/N): y
📋 Enter code (end with 'EOF' on a new line):
git push --force-with-lease origin $(git branch --show-current)
EOF
```

### Listing Snippets

Show all snippets:
```bash
codestash list
```

**Flags:**
- `-l, --language <lang>`: Filter by programming language
- `-t, --tag <tag>`: Filter by tag
- `-e, --expanded`: Show full code content

**Examples:**
```bash
# List all Python snippets
codestash list --language python

# List snippets tagged with 'docker'
codestash list --tag docker

# List all snippets with full code content
codestash list --expanded
```

### Searching Snippets

Search across all snippet fields:
```bash
codestash search <query>
```

**Flags:**
- `-e, --expanded`: Show full code content in results
- `-x, --executable`: Show only executable snippets

**Examples:**
```bash
# Search for snippets containing 'docker'
codestash search docker

# Search for executable snippets only
codestash search --executable deploy

# Search with expanded view
codestash search --expanded "git push"
```

### Using Snippets

The `use` command is your primary interface for working with snippets:

```bash
codestash use <snippet-id-or-title>
```

**Flags:**
- `-c, --copy`: Copy snippet to clipboard
- `-x, --execute`: Execute the snippet (if marked as executable)
- `-f, --force`: Force execution even if not marked as executable

**Examples:**
```bash
# Print snippet to terminal
codestash use "git status"

# Copy snippet to clipboard
codestash use --copy "docker build"

# Execute snippet
codestash use --execute "deploy script"

# Force execute non-executable snippet
codestash use --execute --force "some command"
```

### Individual Commands

You can also use dedicated commands for specific actions:

#### Print Command
```bash
codestash print <snippet-id-or-title>
```
Prints the snippet content to the terminal.

#### Copy Command
```bash
codestash copy <snippet-id-or-title>
```
Copies the snippet code to your clipboard.

#### Execute Command
```bash
codestash exec <snippet-id-or-title>
```

**Flags:**
- `-f, --force`: Force execution even if not marked as executable

**Examples:**
```bash
# Execute an executable snippet
codestash exec "backup script"

# Force execute any snippet
codestash exec --force "some command"
```

### Editing Snippets

Edit existing snippets:
```bash
codestash edit <snippet-id-or-title>
```

**Flags:**
- `-f, --field <field>`: Edit specific field only

**Valid fields:** `title`, `description`, `language`, `tags`, `executable`, `code`

**Examples:**
```bash
# Interactive edit (all fields)
codestash edit "git push"

# Edit only the title
codestash edit --field title "old title"

# Edit only tags
codestash edit --field tags "docker script"
```

### Deleting Snippets

Remove snippets:
```bash
codestash delete <snippet-id-or-title>
```

**Flags:**
- `-f, --force`: Delete without confirmation

**Examples:**
```bash
# Delete with confirmation
codestash delete "old script"

# Delete without confirmation
codestash delete --force "unused snippet"
```

### Usage Statistics

View detailed analytics:
```bash
codestash stats
```

**Flags:**
- `-d, --detailed`: Show detailed statistics including unused snippets

**Example output:**
```
📊 CodeStash Statistics
═══════════════════════════════════════

📚 Total Snippets: 25
🚀 Executable Snippets: 15
📋 Non-executable Snippets: 10
📈 Total Usage: 127 times
📊 Average Usage: 5.1 times per snippet

🏆 Top 5 Most Used Snippets:
───────────────────────────────────────
1. docker-compose up — used 23 times (last used: 2 hours ago)
2. git status check — used 18 times (last used: 1 day ago)
3. npm run build — used 15 times (last used: 3 days ago)
```

## 🔧 Configuration

CodeStash stores all data in `~/.codestash/snippets.json`. The file is created automatically when you add your first snippet.

### Supported Languages for Execution

Executable snippets support various shell languages:
- **Unix/Linux/macOS**: `bash`, `shell`, `sh`, `zsh`, `fish`
- **Windows**: `powershell`, `ps1`, `cmd`, `bat`, `batch`

### Multi-line Command Support

CodeStash fully supports multi-line commands:

```bash
# Example: Multi-line Docker setup
#!/bin/bash
docker build -t myapp .
docker run -d -p 8080:80 myapp
echo "Application deployed on port 8080"
```

## 💡 Examples

### Common Workflow

1. **Add a new snippet:**
   ```bash
   codestash add
   # Enter: "Deploy to production"
   # Description: "Build and deploy the application"
   # Language: bash
   # Tags: deploy, production
   # Executable: y
   ```

2. **Search for deployment scripts:**
   ```bash
   codestash search deploy
   ```

3. **Execute the deployment:**
   ```bash
   codestash exec "Deploy to production"
   # or
   codestash use --execute "Deploy to production"
   ```

4. **View usage statistics:**
   ```bash
   codestash stats --detailed
   ```

### Organizing Your Snippets

Use tags to organize snippets by:
- **Technology**: `docker`, `kubernetes`, `git`, `npm`
- **Purpose**: `deploy`, `backup`, `cleanup`, `setup`
- **Environment**: `production`, `development`, `staging`

### Clipboard Integration

CodeStash integrates with your system clipboard:
- **macOS**: Uses `pbcopy`
- **Linux**: Uses `xclip` or `xsel`
- **Windows**: Uses `clip`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Uses Go's built-in JSON package for local storage
- Cross-platform clipboard support

---

**CodeStash** - Because every developer needs a good stash! 🧰✨