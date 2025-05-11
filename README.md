# Git Project Renamer

This Go CLI tool recursively scans directories under your home directory, identifies Git repositories, extracts their first commit date, and prints the directory to include the project name and date.

## 🔧 How It Works

The core steps in `main.go` are:

1. Get the user's home directory.
2. Recursively list all subdirectories.
3. Filter only those that contain a `.git` directory (i.e. Git repositories).
4. For each Git repository:
    - Extract the date of the first commit.
    - Rename the directory by appending the date.
    - Print the renamed path.

## 🚀 Usage

### Build

```bash
go build -o git-renamer
````

### Run

```bash
./git-renamer
```

This will output something like:

```text
renamed: renamer-20250511
```

## 📁 Example

If you have this project structure:

```
/home/user/code/
├── myproject/.git
├── another/.git
├── notgit/
```

After running the tool:

```
myproject/      → myproject-20230101/
another/        → another-20221201/
```

## 📚 Requirements

* Go 1.16 or higher
* Git installed and accessible via `git` command

## 🧪 Testing

You can test the core logic using Go’s built-in testing framework:

```bash
go test ./...
```

## 🛠️ Project Structure

* `main.go` - The entry point of the CLI
* `listAllPaths.go` - Walks the file system
* `gitops.go` - Contains Git-related utilities
* `rename.go` - Handles renaming logic

## 📄 License

MIT License

---

Feel free to fork or contribute!
