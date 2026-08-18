# 🐚 Mini Bash in GoLang

A simple command-line interface (CLI) shell written in **Go**. This project demonstrates basic interaction with the file system and operating system processes using Go's standard `os` package.

---

## 🚀 Features

The interactive shell supports the following core commands:

* 📂 **`ls`** — List files and directories in the current folder.
* 📁 **`cd`** — Change the current working directory.
* 🛠️ **`mkdir`** — Create a new directory.
* 📄 **`touch`** — Create a new empty file.
* 🚪 **`exit`** — Exit the interactive shell session.

---

## 🛠 Prerequisites

* **Go** compiler installed (version 1.18 or higher).

---

## 📥 Getting Started

1. **Clone the repository** (or download `main.go` directly):
   ```bash
   git clone [https://github.com/your-username/bash-in-golang.git](https://github.com/your-username/bash-in-golang.git)
   cd bash-in-golang

2. **Run the program:**
    ```go run bash-in-golang.go```

3. **Or build an executable:**
```go build -o mybash main.go
./mybash
```
**💻 Usage Example**
Welcome to Bash in GoLang
Please choose operator: ls, cd, mkdir, touch, exit

> ls
main.go

> mkdir test_folder
Enter folder name:
test_folder
Directory successfully created: test_folder

> cd test_folder
Enter the path to the folder:
test_folder
We have successfully transitioned to: /path/to/test_folder

> touch file.txt
Enter file name:
file.txt
File successfully created: file.txt

> exit
Goodbye!