# bookmark

## Description

`bookmark` is a tool that allows you to bookmark and navigate directories via the CLI command line.  
This tool is available in `bash`, `fish` and `powershell` shells.

Once installed, `bookmark` command and `bm` and `bma` commands are available.

**Note: This tool requires [peco](https://github.com/peco/peco)**

## Usage

Bookmark the current directory:

```bash
bma # or 'bookmark add'
```

Go to the selected directory from the bookmark list:

```bash
bm
```

Open the bookmark list and delete the selected directory:

```bash
bookmark remove
```

Delete non-existent directories from the database:

```bash
bookmark refresh
```

## Install & Setup

Currently binaries are released for **Windows (amd64) and Linux (amd64, arm64)**.

Binaries for each platform are available from the [release page](https://github.com/kawana77b/bookmark/releases).  
Place the executable in your shell's `$PATH`ed directory.

The following script must be added to the shell configuration file.

### bash

Add the following to `~/.bashrc`

```bash
eval "$(bookmark init bash)"
```

### fish

Add the following to `~/.config/fish/config.fish`

```fish
bookmark init fish | source
```

### powershell

Add the following to your `$PROFILE`

```powershell
Invoke-Expression (& {
    (bookmark init powershell | Out-String)
})
```

## Database

This tool uses a SQLite database. The database is created in the `bookmark` directory at the following location:

- `%LocalAppData%` (Windows)
- `$HOME/.cache` (Linux)

The location of the path can be checked with `bookmark status`

**If you want to delete the database completely, please manually delete `bookmark` directory.**

## Lisence

The source code license for this repository is MIT.  
Feel free to fork it, as it is a personal tool.  
I am a learner of the program, so please feel free to send me any bug report, feedback or advice you may have :wink:
