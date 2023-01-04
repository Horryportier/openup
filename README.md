# OpenUp
___
[install](#installation) | [unistall](#uninstall) | [keybinds](#keybinds)

### About
> this is simple tui for opening project, config, etc files.

### How to use

```mermaid
flowchart TD;
        O[type openup in terminal]--> A[in app]
        A--> B[Press `A` to add new item]
        A--> C[Press `D` to delete item]
        A--> D[Press `E` to change editor]
        A--> F[Press `C` to change item]
        A--> E[Press `enter` to open file]
```
<p align="center">
<img src="https://raw.githubusercontent.com/Horryportier/openup/main/v1/openupvid.gif" width=500 />
</p>

### Installation:

#### Linux:

```bash
go install github.com/Horryportier/openup@latest 
```
### Uninstall:

#### Linux:

```bash
   rm -rf ~/.openup
   rm ~/go/bin/openup
```


## To implement
- [ ] lunch tmux/kitty session
- [ ] help for all views

## keybinds

- standard bubbletea list bindings
- change existing item {C} not working 
- delete item {D}
- add item {A}
- change editor {E}

