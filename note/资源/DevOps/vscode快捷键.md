配置文件位置

```
C:\Users\admin\AppData\Roaming\Code\User\keybindings.json
```

该文件保存的是用户级快捷键列表，每一个快捷键配置都是包含`key`、`command`、`when`（可选）字段

| 字段          | 作用                             |
| :------------ | :------------------------------- |
| **`key`**     | 你按下的键盘组合                 |
| **`command`** | 要执行的命令 ID                  |
| **`when`**    | 规定这条快捷键在什么“语境”下生效 |

完整示例：

```json
  // markdown 代码块
  {
    "key": "ctrl+shift+`",
    "command": "markdown.extension.editing.toggleCodeBlock",
    "when": "editorTextFocus && editorLangId == 'markdown'"
  },
```



## 按键的书写规则

### 基本格式

- **全小写**：修饰键和字母键都要小写，写 `ctrl` 而不是 `Ctrl`，写 `a` 而不是 `A`。
- **用 `+` 连接**：按键之间用加号连接，不带空格。例如 `ctrl+shift+a`。
- **修饰键顺序任意**：`ctrl+alt+a` 和 `alt+ctrl+a` 效果相同，但习惯按 `ctrl -> shift -> alt -> 字母` 的顺序写

### 支持的修饰键

| 修饰键        | 写法    | 说明                                            |
| :------------ | :------ | :---------------------------------------------- |
| Ctrl          | `ctrl`  | Windows / Linux 上常用 `ctrl`，macOS 对应 `cmd` |
| Shift         | `shift` |                                                 |
| Alt           | `alt`   | macOS 上是 `opt`（或者也支持 `alt`）            |
| Cmd (Mac)     | `cmd`   | 在 Windows 上无效，建议用 `ctrl` 兼顾跨平台     |
| Win (Windows) | `win`   | 很少用                                          |
| Meta (macOS)  | `meta`  | 通常等同 `cmd`                                  |

### 特殊按键名称

| 按键         | 写法                                   |
| :----------- | :------------------------------------- |
| 方向键       | `up` `down` `left` `right`             |
| 功能键       | `f1` ~ `f19`                           |
| 回车         | `enter`                                |
| 空格         | `space`                                |
| 退格         | `backspace`                            |
| Tab          | `tab`                                  |
| 反引号       | ``` (就是 backtick)                    |
| 逗号、句号   | `,` `.` 等直接写符号本身               |
| 数字键       | `0`~`9`                                |
| 计算器小键盘 | `numpad0` ~ `numpad9`，`numpad_add` 等 |

### 按键组合

VS Code 支持**二段按键**（一次按键序列），写法是：

```json
"key": "ctrl+k ctrl+b"
```

意思是先按 `Ctrl+K`，松开后再按 `Ctrl+B`。

## 命令ID

`command` 是一个**全局唯一的字符串**，由扩展或 VS Code 自己注册。不能自己编造，必须使用确切的 ID。

通过`ctrl+k ctrl+b`按键组合或者在设置中点击“键盘快捷方式”即可查看当前能设置的所有的命令。虽然显示的是命令，但是实际上是命令的名字，右键可以获取命令的ID。也可以不通过修改配置文件，在图形化界面设置快捷键。



## 上下文条件表达式

`when` 是一个**布尔表达式**，只有计算结果为 `true` 时，快捷键才会触发，可以组合变量使用逻辑运算符。

### 常用上下文变量

| 变量                            | 说明                                               | 常见值                                      |
| :------------------------------ | :------------------------------------------------- | :------------------------------------------ |
| `editorTextFocus`               | 焦点是否在编辑器的**文本区域**（光标可以输入文字） | `true` / `false`                            |
| `editorLangId`                  | 当前编辑器的语言 ID（不是文件后缀）                | `"markdown"`, `"javascript"`, `"python"` 等 |
| `editorHasSelection`            | 编辑器里是否有文本被选中                           | `true` / `false`                            |
| `isLinux`, `isWindows`, `isMac` | 当前操作系统                                       | `true` / `false`                            |
| `resourceExtname`               | 当前文件的后缀名（带`.`）                          | `".md"`, `".js"` 等                         |
| `activeEditor`                  | 是否有激活的编辑器                                 | `true` / `false`                            |
| `terminalFocus`                 | 焦点是否在集成终端                                 | `true` / `false`                            |
| `listFocus`                     | 焦点在列表/树视图吗                                | `true` / `false`                            |

### 逻辑运算符

| 运算符     | 含义                                                         |
| :--------- | :----------------------------------------------------------- |
| `&&`       | 与                                                           |
| `          |                                                              |
| `!`        | 非                                                           |
| `==`, `!=` | 等于 / 不等于（字符串）                                      |
| `=~`       | 正则匹配（例如：`editorLangId =~ /^markdown$                 |
| `in`       | 测试左侧值是否在给定的右值数组里，如 `editorLangId in ['markdown', 'rmd']` |

### 例子分析

```json
"when": "editorTextFocus && editorLangId == 'markdown'"
```

表达式含义：确保你正在编辑文本，而不是在侧边栏、终端或其他地方**并且**当前文件被 VS Code 识别为 Markdown 语言

## 常见问题排查

| 现象                             | 可能原因                                                |
| :------------------------------- | :------------------------------------------------------ |
| 按键完全没反应                   | `command` 拼写错误；`when` 条件一直为 false；扩展未激活 |
| 在其他语言文件里也生效了         | 缺少 `editorLangId == 'xxx'` 限制，考虑加上             |
| 一次性按了没反应，但按键序列可以 | 可能你定义成了 chord，但习惯是单键                      |
| 提示 “command 'xxx' not found”   | 命令 ID 不存在，检查扩展是否已安装并启用                |