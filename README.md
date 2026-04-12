# OpenList-115

基于 OpenList 的 115 云盘增强版本，包含以下功能：

## 主要特性

### 1. 115 Open 驱动缩略图支持
- 支持通过 Cookie 或 QRCode 登录 115 账号获取缩略图
- 文件列表使用官方 API 返回，缩略图通过 115 逆向驱动获取
- 支持二维码源设备选择（web, android, ios, tv 等）

### 2. 115 逆向驱动默认排序
- 默认按文件名升序排序，与 115 官方客户端一致

### 3. ArtPlayer 进度条优化
- DPlayer 风格的进度条拖动，拖动时只更新 UI
- 释放鼠标时才真正跳转，避免发送大量请求

## 安装

### 下载预编译二进制

从 [Releases](https://github.com/xuebiya/Openlist-115/releases) 页面下载对应平台的二进制文件。

### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/xuebiya/Openlist-115.git
cd Openlist-115

# 编译
go build -o openlist -ldflags="-w -s" -tags=jsoniter .
```

## 115 Open 驱动配置

在管理面板中添加 115 Open 存储时，配置选项：

| 选项 | 说明 |
|------|------|
| Access Token | 必填，115 官方 API Access Token |
| Refresh Token | 必填，115 官方 API Refresh Token |
| Cookie 115 | 可选，115 逆向 Cookie，用于获取缩略图 |
| QRCode Token 115 | 可选，115 二维码 Token，用于获取缩略图 |
| QRCode Source 115 | 可选，二维码源设备（web, android, ios 等）|

> 注意：缩略图功能需要配置 Cookie 115 或 QRCode Token 115 二者之一

## 构建

本项目使用 GitHub Actions 自动构建：

- Linux amd64/arm64
- Windows amd64/386 (支持 Windows 7)

触发构建：访问 Actions 页面，点击 "Build" workflow 并运行。

## License

基于 OpenList 项目，遵循 AGPL-3.0 许可证。
