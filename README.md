# OpenUp
___

### About
> this is simple tui for opening project, config, etc files.

### How to use

```mermaid
flowchart TD;
        O[type openup in terminal]--> A[in app]
        A--> B[Press `A` to add new item]
        A--> C[Press `D` to delete item]
        A--> D[Press `E` to change editor]
        A--> E[Press `enter` to open file]
```
        ___
![](./v1/openupvid.gif)
        ___

### Installation:

#### Linux:

1. Clone repo:
```
   git clone https://github.com/Horryportier/openup
   cd openup
```
2. execute install script.
```
chmod +x install.sh

./install.sh
```



## To implement
- [x] Adding/delting items
- [x] adding data to json
- [x] choice of editor
- [x] un/install script
- [ ] better style
- [ ] changing existing item
- [ ] choice to switch to dir or stay in one you opened the app

#### Maybe
- [ ] custom key binds

## keybinds (may change)

- standard bubbletea list bindings
- change existing item {c}
- delete item {D}
- add item {A}
- change editor {E}

