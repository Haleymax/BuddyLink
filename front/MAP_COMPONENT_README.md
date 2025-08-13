# 地图组件使用说明

## 简介

这是一个基于高德地图API的Vue 3地图组件，提供了位置定位、经纬度输入定位等功能。

## 功能特性

- ✨ 基于高德地图API 2.0
- 📍 获取当前位置（GPS定位）
- 🎯 通过经纬度坐标定位
- 🗺️ 地图标记和中心点控制
- 📱 响应式设计，支持移动端
- 🎨 美观的UI界面，基于Naive UI

## 安装和配置

### 1. 获取高德地图API密钥

1. 访问[高德开放平台](https://lbs.amap.com/)
2. 注册账号并创建应用
3. 获取Web服务API密钥和安全密钥

### 2. 配置API密钥

在 `src/components/MapComponent.vue` 文件中，找到以下配置并替换为你的密钥：

```typescript
// 高德地图配置
const AMAP_KEY = 'YOUR_AMAP_KEY'; // 请替换为你的高德地图API密钥
const AMAP_SECURITY_CODE = 'YOUR_SECURITY_CODE'; // 请替换为你的安全密钥
```

### 3. 域名白名单配置

在高德开放平台的控制台中，需要配置域名白名单：
- 开发环境：`localhost`、`127.0.0.1`
- 生产环境：你的实际域名

## 使用方法

### 在页面中使用

```vue
<template>
  <div>
    <MapComponent />
  </div>
</template>

<script setup>
import MapComponent from '@/components/MapComponent.vue';
</script>
```

### 组件功能

#### 1. 获取当前位置
- 点击"获取当前位置"按钮
- 系统会请求地理位置权限
- 获取成功后在地图上显示标记点
- 经纬度输入框会自动填充

#### 2. 手动输入坐标定位
- 在经度和纬度输入框中输入坐标
- 点击"定位到指定坐标"按钮
- 地图会移动到指定位置并显示标记

#### 3. 位置信息显示
- 显示当前位置的经纬度
- 显示详细地址信息（通过逆地理编码）

## 组件API

### Props
目前组件不接受外部props，所有配置都在组件内部。

### Events
组件内部处理所有交互，不对外暴露事件。

### 方法
- `getCurrentLocation()`: 获取当前位置
- `navigateToCoordinate()`: 定位到指定坐标
- `addMarker(lng, lat)`: 在地图上添加标记点

## 技术栈

- Vue 3 Composition API
- TypeScript
- Naive UI
- 高德地图API 2.0

## 注意事项

### 1. HTTPS要求
现代浏览器要求地理位置API必须在HTTPS环境下使用。如果在HTTP环境下，地理位置功能可能无法正常工作。

### 2. 权限问题
首次使用时，浏览器会请求地理位置权限，用户需要允许才能使用定位功能。

### 3. 网络要求
组件需要联网加载高德地图资源，请确保网络连接正常。

### 4. 坐标系说明
高德地图使用GCJ-02坐标系（火星坐标系），与GPS原始坐标（WGS-84）可能存在偏差。

## 常见问题

### Q: 地图无法加载？
A: 检查API密钥是否正确配置，网络是否正常，域名是否在白名单中。

### Q: 定位功能不工作？
A: 检查是否在HTTPS环境下，是否授权了地理位置权限。

### Q: 地图显示空白？
A: 检查容器是否有明确的宽高设置，API密钥是否有效。

## 自定义扩展

### 添加地图样式
可以在 `initMap()` 方法中修改 `mapStyle` 参数：

```typescript
map.value = new window.AMap.Map('amap-container', {
  zoom: 13,
  center: [116.39, 39.9],
  mapStyle: 'amap://styles/dark', // 修改为其他样式
  viewMode: '3D'
});
```

### 添加更多控件
在 `initMap()` 方法中添加其他控件：

```typescript
// 添加比例尺控件
map.value.addControl(new window.AMap.Scale());
// 添加工具栏
map.value.addControl(new window.AMap.ToolBar());
// 添加鹰眼控件
map.value.addControl(new window.AMap.OverView());
```

## 更新日志

### v1.0.0
- 初版发布
- 支持位置获取和坐标定位
- 基础地图功能完善
