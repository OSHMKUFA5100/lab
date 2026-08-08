### excelize库

读写Excel的Go库

[github](https://github.com/qax-os/excelize)

## 基本用法

#### 创建工作簿

```go
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // 创建一个新的Sheet,返回的index为该工作簿的索引
    index, err := f.NewSheet("SheetName")
    if err != nil {
        return fmt.Errorf("创建Sheet失败: %w", err)
    }
```

#### 打开文件

```go
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("打开 Excel失败: %w", err)
	}
	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
```

#### 写入cell值

```go
  f.SetCellValue("SheetName", "B2", 100)
```

#### 设置工作簿默认活动工作表

```go
f.SetActiveSheet(index)
```

#### 创建或更新excel表

```go
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        return fmt.Errorf("保存失败: %w", err)
    }
```

#### 获取指定cell值

```go
    cell, err := f.GetCellValue("Sheet1", "B2")
    if err != nil {
        return fmt.Errorf("读取cell失败: %w", err)
    }
    fmt.Println(cell)
```

#### 获取所有工作簿的cell值（二维数组）

```go
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        return fmt.Errorf("读取失败: %w", err)
    }
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
```

