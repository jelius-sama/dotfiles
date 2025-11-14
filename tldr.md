# dotfiles

This repository contains my system configuration, organized into **dedicated branches** so each component remains isolated and easy to deploy.

## Branch Layout

| Branch     | Purpose                                        |
| ---------- | ---------------------------------------------- |
| `zinit`    | Zsh plugin manager (zinit) + plugin submodules |
| `zsh`      | Zsh shell config                               |
| `nvim`     | Neovim configuration                           |
| `tmux`     | Tmux configuration                             |
| `anifetch` | Anifetch theme/config                          |
| `neofetch` | Neofetch setup                                 |

---

# Zinit Directory Structure

```
~/.local/share/zinit/
├── zinit.git/               # submodule
└── plugins/                 # all plugins as submodules
    ├── Aloxaf---fzf-tab/
    ├── romkatv---powerlevel10k/
    ├── tolkonepiu---catppuccin-powerlevel10k-themes/
    ├── zsh-users---zsh-autosuggestions/
    ├── zsh-users---zsh-completions/
    └── zsh-users---zsh-syntax-highlighting/
```

---

# Cloning on a New Machine

```sh
git clone --branch zinit --recursive https://github.com/jelius-sama/dotfiles.git ~/.local/share/zinit
```

If submodules are not pulled:

```sh
git submodule update --init --recursive
```

---

# Updating Plugins

Update everything:

```sh
git submodule update --remote --merge
```

Update one plugin:

```sh
git submodule update --remote plugins/<name>
```

Commit updates:

```sh
git add .
git commit -m "Update zinit plugins"
git push
```

---

# Adding a New Plugin

```sh
git rm --cached -r plugins/<folder>
git submodule add <url> plugins/<folder>
git add .
git commit -m "Add plugin <name>"
git push
```

---

# Resetting a Broken Plugin

```sh
git submodule deinit -f plugins/<name>
rm -rf .git/modules/plugins/<name>
rm -rf plugins/<name>
git submodule add <url> plugins/<name>
```
