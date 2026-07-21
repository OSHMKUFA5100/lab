> 感觉还是不如typora好使啊,以后买个正版吧

```css
/**
 * MarkText One Light Dark 主题 (仿 Typora onelight-dark：https://github.com/caolib/typora-onelight-theme)
 *
 * 使用方法:
 *   在 MarkText 中打开 设置 → 主题 → 自定义 CSS，
 *   将本文件全部内容粘贴进去，保存即可。
 *
 * 配色来源: Typora onelight-dark theme by caolib
 * 兼容: MarkText v0.19.x (旧引擎) / v0.20+ (新引擎)
 */

/* ====================================================================
 * CSS 变量 — 使用 camelCase 命名 (旧引擎 v0.19.x 兼容)
 * ==================================================================== */
:root {
  /* —— 编辑器核心 —— */
  --themeColor: #5c99f3;
  --themeColor90: rgba(92, 153, 243, 0.9);
  --themeColor80: rgba(92, 153, 243, 0.8);
  --themeColor70: rgba(92, 153, 243, 0.7);
  --themeColor60: rgba(92, 153, 243, 0.6);
  --themeColor50: rgba(92, 153, 243, 0.5);
  --themeColor40: rgba(92, 153, 243, 0.4);
  --themeColor30: rgba(92, 153, 243, 0.3);
  --themeColor20: rgba(92, 153, 243, 0.2);
  --themeColor10: rgba(92, 153, 243, 0.1);
  --highlightThemeColor: #e94ab9;

  --highlightColor: #ffffff10;
  --selectionColor: #3266d0;
  --editorColor: #c3c3c3;
  --editorColor80: #c3c3c3;
  --editorColor60: rgba(195, 195, 195, 0.6);
  --editorColor50: rgba(195, 195, 195, 0.5);
  --editorColor40: rgba(195, 195, 195, 0.4);
  --editorColor30: rgba(195, 195, 195, 0.3);
  --editorColor10: rgba(195, 195, 195, 0.1);
  --editorColor04: rgba(195, 195, 195, 0.04);
  --editorBgColor: #1f1f1e;
  --editorAreaWidth: 900px;
  --deleteColor: #ff6b6b;
  --iconColor: rgba(195, 195, 195, 0.56);
  --codeBgColor: #282c34;
  --codeBlockBgColor: #282c34;
  --footnoteBgColor: rgba(51, 51, 51, 0.5);
  --inputBgColor: #3d404e;
  --focusColor: #5c99f3;

  /* —— 按钮 —— */
  --buttonFontColor: #c3c3c3;
  --buttonBgColor: #333333;
  --buttonBorder: 1px solid #353a45;
  --buttonShadow: none;
  --buttonFontColorHover: #ffffff;
  --buttonBgColorHover: #414855;
  --buttonBorderHover: var(--buttonBorder);
  --buttonFontColorActive: var(--buttonFontColor);
  --buttonBgColorActive: #2a2d35;
  --buttonBorderActive: var(--buttonBorder);
  --buttonFocusBorder: 1px solid #5c99f3;
  --buttonFocusShadow: 0 0 0 1px rgba(92, 153, 243, 0.3);

  --buttonPrimaryFontColor: #ffffff;
  --buttonPrimaryBgColor: #5c99f3;
  --buttonPrimaryBorder: 1px solid #4a8ad9;
  --buttonPrimaryShadow: none;
  --buttonPrimaryFontColorHover: var(--buttonPrimaryFontColor);
  --buttonPrimaryBgColorHover: #4a8ad9;
  --buttonPrimaryBorderHover: var(--buttonPrimaryBorder);
  --buttonPrimaryFontColorActive: var(--buttonPrimaryFontColor);
  --buttonPrimaryBgColorActive: #3d7bc4;
  --buttonPrimaryBorderActive: var(--buttonPrimaryBorder);
  --buttonPrimaryFocusBorder: none;
  --buttonPrimaryFocusShadow: inset 0 0 0 1px rgba(24, 26, 31, 0.5), 0 0 0 1px #5c99f3;

  /* —— Markdown 元素 —— */
  --headingColor: #c3c3c3;
  --h1Color: #c3c3c3;
  --h2Color: #c3c3c3;
  --h3Color: #c3c3c3;
  --h4Color: #ffffff;
  --h5Color: #c3c3c3;
  --h6Color: #c3c3c3;
  --blockquoteTextColor: rgba(195, 195, 195, 0.8);
  --blockquoteBorderColor: #5c99f3;
  --hrColor: #5c99f3;
  --linkColor: #e94ab9;
  --strongColor: #5c99f3;
  --emColor: #e94ab9;
  --listMarkerColor: #5c99f3;
  --tableBorderColor: #393e49;

  /* —— 侧边栏 —— */
  --sideBarColor: #c3c3c3;
  --sideBarIconColor: var(--iconColor);
  --sideBarTitleColor: #c3c3c3;
  --sideBarTextColor: #7a7a7a;
  --sideBarBgColor: #1f1f1e;
  --sideBarItemHoverBgColor: #333333;
  --itemBgColor: #333333;

  /* —— 浮动层 —— */
  --floatFontColor: #c3c3c3;
  --floatBgColor: #353a45;
  --floatHoverColor: #414855;
  --floatBorderColor: #353a45;
  --floatShadow: rgba(0, 0, 0, 0.3);
  --maskColor: rgba(0, 0, 0, 0.7);
}

/* ====================================================================
 * 全局背景 — 统一为 Typora onelight-dark 底色 #1f1f1e
 * ==================================================================== */
html, body {
  background-color: #1f1f1e;
}
.editor-wrapper {
  background-color: #1f1f1e;
}
.editor-component {
  background-color: #1f1f1e;
}
#ag-editor-id {
  background-color: #1f1f1e;
}

/* ====================================================================
 * 滚动条
 * ==================================================================== */
::-webkit-scrollbar,
::-webkit-scrollbar-corner {
  background: var(--editorBgColor);
}
::-webkit-scrollbar:vertical { width: 8px; }
::-webkit-scrollbar:vertical:hover { width: 10px; }
::-webkit-scrollbar-thumb {
  background: rgba(134, 134, 134, 0.27);
  border-radius: 4px;
}
::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.5);
}

/* ====================================================================
 * 弹窗 (Element UI)
 * ==================================================================== */
.v-modal { background: #08090a !important; opacity: 0.65 !important; }
.el-dialog,
.el-dialog.ag-dialog-table {
  border-radius: 6px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3) !important;
  border: 1px solid #353a45 !important;
  background: #2a2a29 !important;
}
.el-button:focus { color: var(--buttonFontColor); }
.el-button--primary:focus {
  color: var(--buttonPrimaryFontColor);
  background: var(--buttonPrimaryBgColor);
}

/* ====================================================================
 * 标题栏 / 浮动菜单
 * ==================================================================== */
.title-bar .frameless-titlebar-button > div > svg { fill: #c3c3c3; }
.title-bar .frameless-titlebar-minimize:hover,
.title-bar .frameless-titlebar-toggle:hover {
  background-color: rgba(195, 195, 195, 0.05) !important;
}
.input-wrapper {
  border-radius: 5px !important;
  border: 1px solid #353a45 !important;
  background: #2a2a29 !important;
}
.ag-front-menu,
.ag-float-wrapper {
  box-shadow: 0 4px 8px 0 var(--floatShadow) !important;
}

/* ====================================================================
 * 侧边栏
 * ==================================================================== */
.side-bar { border-right: 1px solid #333333 !important; }
.left-column ul > li > svg { fill: rgba(195, 195, 195, 0.6) !important; }
.left-column ul > li:hover > svg { fill: #ffffff !important; }
.left-column ul > li.active > svg { fill: #5c99f3 !important; }

/* ====================================================================
 * 标签栏
 * ==================================================================== */
.editor-tabs {
  box-shadow: none !important;
  background: var(--editorBgColor) !important;
}
.editor-tabs:after {
  position: absolute;
  content: '';
  border-bottom: 1px solid #333333;
  bottom: 0; left: 0; right: 0; z-index: 1;
}
.editor-tabs ul.tabs-container:after {
  position: absolute;
  content: '';
  border-bottom: 1px solid #333333;
  bottom: 0; left: 0; right: 0; z-index: 2;
}
.tabs-container > li,
.tabs-container > li.active { background: var(--editorBgColor) !important; }
.tabs-container > li.active {
  border: 1px solid #333333;
  border-bottom: none;
}
.tabs-container > li.active:after {
  top: 0 !important; right: auto !important;
  width: 2px !important; height: auto !important;
  background: #5c99f3 !important;
}
.tabs-container svg.close-icon #unsaved-circle-icon { fill: #5c99f3; }

/* ====================================================================
 * 编辑器内容区 (v0.19.x 旧引擎: 容器是 #ag-editor-id)
 * ==================================================================== */

/* ---- 标题 ---- */
#ag-editor-id h1 {
  text-align: center;
  padding-bottom: 0.3em;
}
#ag-editor-id h2 {
  border-bottom: 2px solid var(--themeColor);
}
#ag-editor-id h3 {
  border-left: 4px solid var(--themeColor);
  padding-left: 8px;
}
#ag-editor-id h4 {
  color: #ffffff !important;
  background: var(--themeColor);
  border-radius: 5px;
  width: fit-content;
  padding: 2px 12px;
}

/* ---- 粗体 / 斜体 ---- */
#ag-editor-id strong { color: var(--strongColor); font-weight: 700; }
#ag-editor-id em     { color: var(--emColor); }

/* ---- 链接 ---- */
#ag-editor-id a {
  color: var(--linkColor);
  font-weight: 700;
  text-decoration: underline;
}

/* ---- 行内代码 ---- */
#ag-editor-id code {
  background-color: var(--codeBlockBgColor);
  color: #eb4c37;
  border-radius: 3px;
  padding: 2px 4px;
}

/* ---- 代码块 ---- */
#ag-editor-id :not(pre) > code[class*='language-'],
#ag-editor-id pre:not(.CodeMirror-line),
#ag-editor-id pre[class*='language-'],
#ag-editor-id pre.ag-paragraph {
  background: var(--codeBlockBgColor) !important;
  border: none !important;
  border-radius: 5px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  transition: box-shadow 0.3s ease;
}

/* ---- 引用块 (四面边框 + 左侧加粗) ---- */
#ag-editor-id blockquote {
  padding: 10px 10px 10px 20px;
  font-size: 0.9em;
  border: 2px solid var(--blockquoteBorderColor);
  border-left: 5px solid var(--blockquoteBorderColor);
  border-radius: 0 5px 5px 0;
  background: transparent;
  overflow: auto;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
}
#ag-editor-id blockquote:hover {
  background: var(--editorBgColor);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
}
/* 隐藏引擎默认的 ::before 竖条，改用 border-left */
#ag-editor-id blockquote::before {
  content: none;
}

/* ---- 分割线 ---- */
#ag-editor-id hr {
  height: 2px;
  background-color: var(--hrColor);
}

/* ---- 表格 ---- */
#ag-editor-id table tr th::before,
#ag-editor-id table tr td::before {
  border-color: var(--tableBorderColor);
}

/* ---- 图片 ---- */
#ag-editor-id img {
  border-radius: 10px;
  display: block;
  margin: 0 auto;
  max-width: 100%;
}

/* ---- 任务列表复选框 ---- */
#ag-editor-id li.ag-task-list-item {
  list-style-type: none;
  position: relative;
}
#ag-editor-id li.ag-task-list-item > input[type='checkbox'] {
  position: absolute;
  cursor: pointer;
  width: 16px; height: 16px; top: 0.1em;
  margin: 0; left: -24px;
  transform-origin: center;
  transition: all 0.2s ease;
  appearance: none;
  -webkit-appearance: none;
}
#ag-editor-id li.ag-task-list-item > input[type='checkbox']::before {
  content: '';
  width: 16px; height: 16px;
  box-sizing: border-box;
  display: inline-block;
  border: 2px solid var(--editorColor50);
  border-radius: 2px;
  background-color: var(--editorBgColor);
  position: absolute; top: 0; left: 0;
  transition: all 0.2s ease;
}
#ag-editor-id li.ag-task-list-item > input[type='checkbox']:checked::before {
  border-color: transparent;
  background-color: var(--themeColor);
}
#ag-editor-id li.ag-task-list-item > input[type='checkbox']::after {
  content: '';
  transform: rotate(-45deg) scale(0);
  width: 9px; height: 5px;
  border: 2px solid #fff;
  border-top: none; border-right: none;
  position: absolute; display: inline-block;
  top: 1px; left: 5px;
  transition: all 0.2s ease;
}
#ag-editor-id li.ag-task-list-item > input[type='checkbox']:checked::after {
  transform: rotate(-45deg) scale(1);
}

/* ---- 段落悬停 ---- */
#ag-editor-id p.ag-paragraph {
  padding: 5px;
  margin: 0;
  border-radius: 5px;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
}
#ag-editor-id p.ag-paragraph:hover {
  background: #141414;
  box-shadow: 0 4px 4px rgba(0, 0, 0, 0.3);
}

/* ---- 已完成任务文字 ---- */
#ag-editor-id li.ag-task-list-item > input.ag-checkbox-checked ~ * {
  color: var(--editorColor50);
}

/* ---- 列表标记 ---- */
#ag-editor-id ul > li::marker,
#ag-editor-id ol > li::marker {
  color: var(--listMarkerColor);
}

/* ---- 搜索高亮 ---- */
.ag-highlight {
  background: #ff0 !important;
  color: red !important;
  border-radius: 3px;
  padding: 0 2px;
  height: auto !important;
  line-height: inherit !important;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

/* ---- 表格表头 ---- */
#ag-editor-id table tr th {
  background: var(--editorColor04);
}
#ag-editor-id table tr:first-child td {
  background: var(--editorColor04);
}

/* ---- 脚注 ---- */
#ag-editor-id sup.md-footnote {
  color: var(--strongColor);
  font-weight: 700;
}

/* ---- 前置图标 ---- */
.ag-front-icon {
  fill: var(--editorColor30);
}

/* ====================================================================
 * CodeMirror 源码模式 (v0.19.x 旧引擎)
 * ==================================================================== */
.source-code {
  background: var(--editorBgColor);
}
.source-code .CodeMirror {
  color: var(--editorColor);
  background: transparent;
}
.source-code .CodeMirror-gutters {
  background: transparent;
  border-right-color: #333333;
}
.source-code .CodeMirror-linenumber {
  color: var(--editorColor30);
}
.source-code .CodeMirror-cursor {
  border-left-color: var(--themeColor);
}
.source-code .CodeMirror-activeline-background {
  background: var(--editorColor04);
}
.source-code .CodeMirror-selected {
  background: var(--selectionColor) !important;
}

/* CodeMirror 语法高亮色 — 参考 onelight 配色 */
.cm-s-one-dark .cm-comment  { color: #7f848e; font-style: italic; }
.cm-s-one-dark .cm-keyword  { color: #c678dd; }
.cm-s-one-dark .cm-operator  { color: #56b6c2; }
.cm-s-one-dark .cm-builtin   { color: #e5c07b; }
.cm-s-one-dark .cm-atom      { color: #d19a66; }
.cm-s-one-dark .cm-number    { color: #d19a66; }
.cm-s-one-dark .cm-def       { color: #61afef; }
.cm-s-one-dark .cm-string    { color: #98c379; }
.cm-s-one-dark .cm-string-2  { color: #e06c75; }
.cm-s-one-dark .cm-variable  { color: #e06c75; }
.cm-s-one-dark .cm-variable-2 { color: #e5c07b; }
.cm-s-one-dark .cm-variable-3 { color: #d19a66; }
.cm-s-one-dark .cm-type      { color: #e5c07b; }
.cm-s-one-dark .cm-meta      { color: #e06c75; }
.cm-s-one-dark .cm-attribute { color: #d19a66; }
.cm-s-one-dark .cm-property  { color: #e06c75; }
.cm-s-one-dark .cm-qualifier { color: #e5c07b; }
.cm-s-one-dark .cm-tag       { color: #e06c75; }
.cm-s-one-dark .cm-link      { color: #c678dd; }
.cm-s-one-dark .cm-header    { color: #61afef; }
.cm-s-one-dark .cm-quote     { color: #98c379; }

/* ====================================================================
 * Mermaid/图表 SVG 暗色适配
 * ==================================================================== */
figure[data-role] svg rect[fill='#ffffff']       { fill: var(--editorBgColor); }
figure[data-role] svg rect[stroke='#000000']     { stroke: var(--editorColor); }
figure[data-role] svg text[fill='#000000']       { fill: var(--editorColor); }
figure[data-role] svg path[stroke='#000000']     { stroke: var(--editorColor); }
figure[data-role] svg path[fill='#ffffff']        { fill: var(--editorBgColor); }
figure[data-role] svg path[fill='#000000']        { fill: var(--editorColor); }
figure[data-role] svg use[fill='black']           { fill: var(--editorColor); }
figure[data-role] svg use[fill='#000000']         { fill: var(--editorColor); }

```
