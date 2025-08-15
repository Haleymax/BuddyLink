<template>
  <div class="map-container">
    <!-- 在模态框中时的简化布局 -->
    <div v-if="props.isModal" class="modal-layout">
      <!-- 使用提示 -->
      <n-alert type="info" :show-icon="false" style="margin-bottom: 16px;" class="usage-tip">
        <template #icon>
          <n-icon size="18" color="#1890ff">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M11,9H13V7H11M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,17H13V11H11V17Z"/>
            </svg>
          </n-icon>
        </template>
        <div class="tip-content">
          <strong>快速定位指南:</strong> 
          <div class="tip-actions">
            <span>📍 点击地图选择位置</span>
            <span>🎯 GPS获取当前位置</span>
            <span>🔍 搜索目标地址</span>
          </div>
        </div>
      </n-alert>
      
      <!-- 地图容器 -->
      <div class="map-wrapper">
        <div id="amap-container" class="map"></div>
        <div class="map-overlay">
          <n-tag type="success" size="small" class="coordinate-display" v-if="currentLocation">
            📌 {{ currentLocation.longitude.toFixed(4) }}, {{ currentLocation.latitude.toFixed(4) }}
          </n-tag>
        </div>
      </div>
      
      <!-- 控制面板 -->
      <div class="control-panel">
        <n-grid x-gap="16" y-gap="16" cols="1 s:2 m:2 l:2">
          <!-- 左侧：定位和搜索 -->
          <n-grid-item>
            <n-card size="small" title="位置操作" class="operation-card">
              <template #header-extra>
                <n-icon size="16" color="#52c41a">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4M12,6A6,6 0 0,0 6,12A6,6 0 0,0 12,18A6,6 0 0,0 18,12A6,6 0 0,0 12,6M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8Z"/>
                  </svg>
                </n-icon>
              </template>
              
              <n-space vertical size="medium">
                <!-- 获取当前位置按钮 -->
                <n-button 
                  type="primary" 
                  @click="getCurrentLocation"
                  :loading="locationLoading"
                  block
                  size="large"
                  class="location-btn"
                >
                  <template #icon>
                    <n-icon>
                      <svg viewBox="0 0 24 24">
                        <path fill="currentColor" d="M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8M3.05,13H1V11H3.05C3.5,6.83 6.83,3.5 11,3.05V1H13V3.05C17.17,3.5 20.5,6.83 20.95,11H23V13H20.95C20.5,17.17 17.17,20.5 13,20.95V23H11V20.95C6.83,20.5 3.5,17.17 3.05,13M12,5A7,7 0 0,0 5,12A7,7 0 0,0 12,19A7,7 0 0,0 19,12A7,7 0 0,0 12,5Z"/>
                      </svg>
                    </n-icon>
                  </template>
                  获取当前位置
                </n-button>

                <!-- 地址搜索 -->
                <div class="search-section">
                  <n-input
                    v-model:value="searchForm.address" 
                    placeholder="输入地址、商店名称或地标"
                    clearable
                    size="large"
                    @keyup.enter="searchAddress"
                    class="search-input"
                  >
                    <template #prefix>
                      <n-icon color="#666">
                        <svg viewBox="0 0 24 24">
                          <path fill="currentColor" d="M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4Z"/>
                        </svg>
                      </n-icon>
                    </template>
                  </n-input>
                  
                  <n-space style="margin-top: 12px;">
                    <n-button 
                      type="success" 
                      @click="searchAddress"
                      :disabled="!searchForm.address || searchForm.address.trim() === ''"
                      :loading="searchLoading"
                      style="flex: 1"
                      ghost
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      搜索位置
                    </n-button>
                    <n-button 
                      type="warning" 
                      @click="clearMarker"
                      ghost
                      style="flex: 1"
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M19,4H15.5L14.5,3H9.5L8.5,4H5V6H19M6,19A2,2 0 0,0 8,21H16A2,2 0 0,0 18,19V7H6V19Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      清除
                    </n-button>
                  </n-space>
                </div>
              </n-space>
            </n-card>
          </n-grid-item>

          <!-- 右侧：位置信息 -->
          <n-grid-item>
            <n-card size="small" title="位置详情" class="info-card">
              <template #header-extra>
                <n-icon size="16" color="#1890ff">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2C13.1,2 14,2.9 14,4C14,5.1 13.1,6 12,6C10.9,6 10,5.1 10,4C10,2.9 10.9,2 12,2M21,9V7L15,1H5C3.89,1 3,1.89 3,3V21A2,2 0 0,0 5,23H19A2,2 0 0,0 21,21V9M12,19A5,5 0 0,1 7,14A5,5 0 0,1 12,9A5,5 0 0,1 17,14A5,5 0 0,1 12,19M12,11.5A2.5,2.5 0 0,0 9.5,14A2.5,2.5 0 0,0 12,16.5A2.5,2.5 0 0,0 14.5,14A2.5,2.5 0 0,0 12,11.5Z"/>
                  </svg>
                </n-icon>
              </template>
              
              <template v-if="currentLocation">
                <div class="location-info">
                  <n-descriptions :column="1" size="small" class="location-details">
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#fa8c16">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M7.07,18.28C7.5,17.38 10.12,16.5 12,16.5C13.88,16.5 16.5,17.38 16.93,18.28C15.57,19.36 13.86,20 12,20C10.14,20 8.43,19.36 7.07,18.28M18.36,16.83C16.93,15.09 13.46,14.5 12,14.5C10.54,14.5 7.07,15.09 5.64,16.83C4.62,15.5 4,13.82 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,13.82 19.38,15.5 18.36,16.83M12,6C10.06,6 8.5,7.56 8.5,9.5C8.5,11.44 10.06,13 12,13C13.94,13 15.5,11.44 15.5,9.5C15.5,7.56 13.94,6 12,6M12,11A1.5,1.5 0 0,1 10.5,9.5A1.5,1.5 0 0,1 12,8A1.5,1.5 0 0,1 13.5,9.5A1.5,1.5 0 0,1 12,11Z"/>
                            </svg>
                          </n-icon>
                          经度
                        </div>
                      </template>
                      <n-text code>{{ currentLocation.longitude }}</n-text>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#fa8c16">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M7.07,18.28C7.5,17.38 10.12,16.5 12,16.5C13.88,16.5 16.5,17.38 16.93,18.28C15.57,19.36 13.86,20 12,20C10.14,20 8.43,19.36 7.07,18.28M18.36,16.83C16.93,15.09 13.46,14.5 12,14.5C10.54,14.5 7.07,15.09 5.64,16.83C4.62,15.5 4,13.82 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,13.82 19.38,15.5 18.36,16.83M12,6C10.06,6 8.5,7.56 8.5,9.5C8.5,11.44 10.06,13 12,13C13.94,13 15.5,11.44 15.5,9.5C15.5,7.56 13.94,6 12,6M12,11A1.5,1.5 0 0,1 10.5,9.5A1.5,1.5 0 0,1 12,8A1.5,1.5 0 0,1 13.5,9.5A1.5,1.5 0 0,1 12,11Z"/>
                            </svg>
                          </n-icon>
                          纬度
                        </div>
                      </template>
                      <n-text code>{{ currentLocation.latitude }}</n-text>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#52c41a">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4Z"/>
                            </svg>
                          </n-icon>
                          地址
                        </div>
                      </template>
                      <n-ellipsis style="max-width: 200px;" :tooltip="{ maxWidth: 300 }">
                        {{ currentLocation.address || '地址获取中...' }}
                      </n-ellipsis>
                    </n-descriptions-item>
                  </n-descriptions>
                  
                  <!-- 确认位置按钮 -->
                  <div class="confirm-section">
                    <n-button 
                      type="primary" 
                      @click="confirmLocation"
                      block
                      size="large"
                      :disabled="!currentLocation"
                      class="confirm-btn"
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      确认使用此位置
                    </n-button>
                  </div>
                </div>
              </template>
              
              <template v-else>
                <div class="empty-location">
                  <n-empty size="small" description="">
                    <template #icon>
                      <n-icon size="48" color="#ccc">
                        <svg viewBox="0 0 24 24">
                          <path fill="currentColor" d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z"/>
                        </svg>
                      </n-icon>
                    </template>
                    <template #extra>
                      <div class="empty-tips">
                        <p>尚未选择位置</p>
                        <n-text depth="3" style="font-size: 12px;">
                          使用左侧功能获取位置信息
                        </n-text>
                      </div>
                    </template>
                  </n-empty>
                </div>
              </template>
            </n-card>
          </n-grid-item>
        </n-grid>
      </div>
    </div>

    <!-- 正常模式的完整布局 -->
    <n-card v-else title="高德地图定位" :bordered="false" class="main-card">
      <template #header-extra>
        <n-tag type="info" size="small">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12,2C13.1,2 14,2.9 14,4C14,5.1 13.1,6 12,6C10.9,6 10,5.1 10,4C10,2.9 10.9,2 12,2M21,9V7L15,1H5C3.89,1 3,1.89 3,3V21A2,2 0 0,0 5,23H19A2,2 0 0,0 21,21V9M12,19A5,5 0 0,1 7,14A5,5 0 0,1 12,9A5,5 0 0,1 17,14A5,5 0 0,1 12,19M12,11.5A2.5,2.5 0 0,0 9.5,14A2.5,2.5 0 0,0 12,16.5A2.5,2.5 0 0,0 14.5,14A2.5,2.5 0 0,0 12,11.5Z"/>
              </svg>
            </n-icon>
          </template>
          智能定位
        </n-tag>
      </template>
      
      <!-- 使用提示 -->
      <n-alert type="info" :show-icon="false" style="margin-bottom: 20px;" class="usage-tip">
        <template #icon>
          <n-icon size="18" color="#1890ff">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M11,9H13V7H11M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,17H13V11H11V17Z"/>
            </svg>
          </n-icon>
        </template>
        <div class="tip-content">
          <strong>快速定位指南:</strong> 
          <div class="tip-actions">
            <span>📍 点击地图选择位置</span>
            <span>🎯 GPS获取当前位置</span>
            <span>🔍 搜索目标地址</span>
          </div>
        </div>
      </n-alert>
      
      <!-- 地图容器 -->
      <div class="map-wrapper">
        <div id="amap-container" class="map"></div>
        <div class="map-overlay">
          <n-tag type="success" size="small" class="coordinate-display" v-if="currentLocation">
            📌 {{ currentLocation.longitude.toFixed(4) }}, {{ currentLocation.latitude.toFixed(4) }}
          </n-tag>
        </div>
      </div>
      
      <!-- 控制面板 -->
      <div class="control-panel">
        <n-grid x-gap="16" y-gap="16" cols="1 s:2 m:2 l:2">
          <!-- 左侧：定位和搜索 -->
          <n-grid-item>
            <n-card size="small" title="位置操作" class="operation-card">
              <template #header-extra>
                <n-icon size="16" color="#52c41a">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4M12,6A6,6 0 0,0 6,12A6,6 0 0,0 12,18A6,6 0 0,0 18,12A6,6 0 0,0 12,6M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8Z"/>
                  </svg>
                </n-icon>
              </template>
              
              <n-space vertical size="medium">
                <!-- 获取当前位置按钮 -->
                <n-button 
                  type="primary" 
                  @click="getCurrentLocation"
                  :loading="locationLoading"
                  block
                  size="large"
                  class="location-btn"
                >
                  <template #icon>
                    <n-icon>
                      <svg viewBox="0 0 24 24">
                        <path fill="currentColor" d="M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8M3.05,13H1V11H3.05C3.5,6.83 6.83,3.5 11,3.05V1H13V3.05C17.17,3.5 20.5,6.83 20.95,11H23V13H20.95C20.5,17.17 17.17,20.5 13,20.95V23H11V20.95C6.83,20.5 3.5,17.17 3.05,13M12,5A7,7 0 0,0 5,12A7,7 0 0,0 12,19A7,7 0 0,0 19,12A7,7 0 0,0 12,5Z"/>
                      </svg>
                    </n-icon>
                  </template>
                  获取当前位置
                </n-button>

                <!-- 地址搜索 -->
                <div class="search-section">
                  <n-input
                    v-model:value="searchForm.address" 
                    placeholder="输入地址、商店名称或地标"
                    clearable
                    size="large"
                    @keyup.enter="searchAddress"
                    class="search-input"
                  >
                    <template #prefix>
                      <n-icon color="#666">
                        <svg viewBox="0 0 24 24">
                          <path fill="currentColor" d="M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4Z"/>
                        </svg>
                      </n-icon>
                    </template>
                  </n-input>
                  
                  <n-space style="margin-top: 12px;">
                    <n-button 
                      type="success" 
                      @click="searchAddress"
                      :disabled="!searchForm.address || searchForm.address.trim() === ''"
                      :loading="searchLoading"
                      style="flex: 1"
                      ghost
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      搜索位置
                    </n-button>
                    <n-button 
                      type="warning" 
                      @click="clearMarker"
                      ghost
                      style="flex: 1"
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M19,4H15.5L14.5,3H9.5L8.5,4H5V6H19M6,19A2,2 0 0,0 8,21H16A2,2 0 0,0 18,19V7H6V19Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      清除
                    </n-button>
                  </n-space>
                </div>
              </n-space>
            </n-card>
          </n-grid-item>

          <!-- 右侧：位置信息 -->
          <n-grid-item>
            <n-card size="small" title="位置详情" class="info-card">
              <template #header-extra>
                <n-icon size="16" color="#1890ff">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2C13.1,2 14,2.9 14,4C14,5.1 13.1,6 12,6C10.9,6 10,5.1 10,4C10,2.9 10.9,2 12,2M21,9V7L15,1H5C3.89,1 3,1.89 3,3V21A2,2 0 0,0 5,23H19A2,2 0 0,0 21,21V9M12,19A5,5 0 0,1 7,14A5,5 0 0,1 12,9A5,5 0 0,1 17,14A5,5 0 0,1 12,19M12,11.5A2.5,2.5 0 0,0 9.5,14A2.5,2.5 0 0,0 12,16.5A2.5,2.5 0 0,0 14.5,14A2.5,2.5 0 0,0 12,11.5Z"/>
                  </svg>
                </n-icon>
              </template>
              
              <template v-if="currentLocation">
                <div class="location-info">
                  <n-descriptions :column="1" size="small" class="location-details">
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#fa8c16">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M7.07,18.28C7.5,17.38 10.12,16.5 12,16.5C13.88,16.5 16.5,17.38 16.93,18.28C15.57,19.36 13.86,20 12,20C10.14,20 8.43,19.36 7.07,18.28M18.36,16.83C16.93,15.09 13.46,14.5 12,14.5C10.54,14.5 7.07,15.09 5.64,16.83C4.62,15.5 4,13.82 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,13.82 19.38,15.5 18.36,16.83M12,6C10.06,6 8.5,7.56 8.5,9.5C8.5,11.44 10.06,13 12,13C13.94,13 15.5,11.44 15.5,9.5C15.5,7.56 13.94,6 12,6M12,11A1.5,1.5 0 0,1 10.5,9.5A1.5,1.5 0 0,1 12,8A1.5,1.5 0 0,1 13.5,9.5A1.5,1.5 0 0,1 12,11Z"/>
                            </svg>
                          </n-icon>
                          经度
                        </div>
                      </template>
                      <n-text code>{{ currentLocation.longitude }}</n-text>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#fa8c16">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M7.07,18.28C7.5,17.38 10.12,16.5 12,16.5C13.88,16.5 16.5,17.38 16.93,18.28C15.57,19.36 13.86,20 12,20C10.14,20 8.43,19.36 7.07,18.28M18.36,16.83C16.93,15.09 13.46,14.5 12,14.5C10.54,14.5 7.07,15.09 5.64,16.83C4.62,15.5 4,13.82 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,13.82 19.38,15.5 18.36,16.83M12,6C10.06,6 8.5,7.56 8.5,9.5C8.5,11.44 10.06,13 12,13C13.94,13 15.5,11.44 15.5,9.5C15.5,7.56 13.94,6 12,6M12,11A1.5,1.5 0 0,1 10.5,9.5A1.5,1.5 0 0,1 12,8A1.5,1.5 0 0,1 13.5,9.5A1.5,1.5 0 0,1 12,11Z"/>
                            </svg>
                          </n-icon>
                          纬度
                        </div>
                      </template>
                      <n-text code>{{ currentLocation.latitude }}</n-text>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <div class="info-label">
                          <n-icon color="#52c41a">
                            <svg viewBox="0 0 24 24">
                              <path fill="currentColor" d="M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4Z"/>
                            </svg>
                          </n-icon>
                          地址
                        </div>
                      </template>
                      <n-ellipsis style="max-width: 200px;" :tooltip="{ maxWidth: 300 }">
                        {{ currentLocation.address || '地址获取中...' }}
                      </n-ellipsis>
                    </n-descriptions-item>
                  </n-descriptions>
                  
                  <!-- 确认位置按钮 -->
                  <div class="confirm-section">
                    <n-button 
                      type="primary" 
                      @click="confirmLocation"
                      block
                      size="large"
                      :disabled="!currentLocation"
                      class="confirm-btn"
                    >
                      <template #icon>
                        <n-icon>
                          <svg viewBox="0 0 24 24">
                            <path fill="currentColor" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/>
                          </svg>
                        </n-icon>
                      </template>
                      确认使用此位置
                    </n-button>
                  </div>
                </div>
              </template>
              
              <template v-else>
                <div class="empty-location">
                  <n-empty size="small" description="">
                    <template #icon>
                      <n-icon size="48" color="#ccc">
                        <svg viewBox="0 0 24 24">
                          <path fill="currentColor" d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z"/>
                        </svg>
                      </n-icon>
                    </template>
                    <template #extra>
                      <div class="empty-tips">
                        <p>尚未选择位置</p>
                        <n-text depth="3" style="font-size: 12px;">
                          使用左侧功能获取位置信息
                        </n-text>
                      </div>
                    </template>
                  </n-empty>
                </div>
              </template>
            </n-card>
          </n-grid-item>
        </n-grid>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import type { Location } from '../model/location';
import '../styles/MapComponent.css';
import AMapLoader from '@amap/amap-jsapi-loader';

// 定义组件的props和emits
interface Props {
  initialLocation?: Location | null
  isModal?: boolean  // 新增属性来标识是否在模态框中
}

const props = withDefaults(defineProps<Props>(), {
  initialLocation: null,
  isModal: false
})

// 定义事件
const emit = defineEmits<{
  locationSelected: [location: Location]
}>()

const message = useMessage();

// 响应式数据
const map = ref<any>(null);
const marker = ref<any>(null);
const AMapRef = ref<any>(null); // 保存 AMap 命名空间
const locationLoading = ref(false);
const searchLoading = ref(false);
const currentLocation = ref<Location | null>(null);

const searchForm = ref({
  address: ''
});

// 计算属性已移除，不再需要验证坐标

// 高德地图配置
const AMAP_KEY = import.meta.env.VITE_AMAP_KEY || '42707d19daa52635acb92b215df96bcc';
const AMAP_SECURITY_CODE = import.meta.env.VITE_AMAP_SECURITY_CODE || '45a9990b03da96393396d53446d5eb6e';

// 初始化地图（接收 AMap 命名空间）
const initMap = () => {
  try {
    const AMap = AMapRef.value;
    if (!AMap) {
      message.error('AMap 未加载');
      return;
    }

    map.value = new AMap.Map('amap-container', {
      zoom: 13,
      center: [116.39, 39.9],
      mapStyle: 'amap://styles/normal',
      viewMode: '3D'
    });

    map.value.on('complete', () => {
      try {
        // 插件已通过 loader 预加载，直接实例化
        const scale = new AMap.Scale();
        map.value.addControl(scale);
        const toolBar = new AMap.ToolBar();
        map.value.addControl(toolBar);
      } catch (error) {
        console.warn('地图控件加载失败:', error);
      }
    });

    map.value.on('click', (e: any) => {
      const { lng, lat } = e.lnglat;
      addMarker(lng, lat);
      const geocoder = new AMap.Geocoder();
      geocoder.getAddress([lng, lat], (status: string, result: any) => {
        if (status === 'complete' && result.regeocode) {
          currentLocation.value = {
            longitude: Number(lng.toFixed(6)),
            latitude: Number(lat.toFixed(6)),
            address: result.regeocode.formattedAddress
          };
        } else {
          currentLocation.value = {
            longitude: Number(lng.toFixed(6)),
            latitude: Number(lat.toFixed(6))
          };
        }
      });
      message.info(`已选择位置: ${lng.toFixed(6)}, ${lat.toFixed(6)}`);
    });

    message.success('地图初始化成功，点击地图任意位置获取坐标');
  } catch (error) {
    console.error('地图初始化失败:', error);
    message.error('地图初始化失败');
  }
};

// 获取当前位置
const getCurrentLocation = async () => {
  if (!map.value || !AMapRef.value) {
    message.error('地图未初始化');
    return;
  }
  const AMap = AMapRef.value;
  locationLoading.value = true;
  try {
    const geolocation = new AMap.Geolocation({
      enableHighAccuracy: true,
      timeout: 10000,
      maximumAge: 1000 * 60 * 60 * 24,
      convert: true,
      showButton: false,
      buttonPosition: 'RB',
      buttonOffset: new AMap.Pixel(10, 20),
      showMarker: false,
      showCircle: false,
      panToLocation: true,
      zoomToAccuracy: true
    });
    geolocation.getCurrentPosition((status: string, result: any) => {
      locationLoading.value = false;
      if (status === 'complete') {
        const { lng, lat, formattedAddress } = result.position;
        currentLocation.value = { longitude: lng, latitude: lat, address: formattedAddress };
        addMarker(lng, lat);
        map.value.setCenter([lng, lat]);
        map.value.setZoom(16);
        message.success('获取位置成功');
      } else {
        message.error('获取位置失败：' + result.message);
      }
    });
    map.value.addControl(geolocation);
  } catch (error) {
    locationLoading.value = false;
    message.error('定位服务出错');
    console.error('定位错误:', error);
  }
};

// 添加标记点
const addMarker = (lng: number, lat: number) => {
  if (!AMapRef.value) return;
  const AMap = AMapRef.value;
  if (marker.value) {
    map.value.remove(marker.value);
  }
  marker.value = new AMap.Marker({
    position: [lng, lat],
    icon: new AMap.Icon({
      size: new AMap.Size(32, 32),
      image: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_r.png',
      imageOffset: new AMap.Pixel(-16, -32),
      imageSize: new AMap.Size(32, 32)
    }),
    title: `坐标: ${lng.toFixed(6)}, ${lat.toFixed(6)}`,
    anchor: 'bottom-center'
  });
  map.value.add(marker.value);
  const infoWindow = new AMap.InfoWindow({
    content: `
      <div style="padding: 8px; min-width: 200px;">
        <div style="font-weight: bold; margin-bottom: 8px; color: #333;">📍 位置坐标</div>
        <div style="margin-bottom: 4px;"><strong>经度:</strong> ${lng.toFixed(6)}</div>
        <div style="margin-bottom: 4px;"><strong>纬度:</strong> ${lat.toFixed(6)}</div>
      </div>
    `,
    offset: new AMap.Pixel(0, -30),
    closeWhenClickMap: true
  });
  marker.value.on('click', () => {
    infoWindow.open(map.value, [lng, lat]);
  });
};

// 地址搜索
const searchAddress = async () => {
  if (!searchForm.value.address || searchForm.value.address.trim() === '' || !map.value || !AMapRef.value) {
    message.warning('请输入有效的地址');
    return;
  }
  searchLoading.value = true;
  try {
    const geocoder = new AMapRef.value.Geocoder({ city: '', radius: 1000, extensions: 'base' });
    geocoder.getLocation(searchForm.value.address, (status: string, result: any) => {
      searchLoading.value = false;
      if (status === 'complete' && result.geocodes && result.geocodes.length > 0) {
        const geocode = result.geocodes[0];
        const location = geocode.location;
        const lng = location.lng;
        const lat = location.lat;
        const formattedAddress = geocode.formattedAddress || geocode.district || searchForm.value.address;
        currentLocation.value = { longitude: Number(lng.toFixed(6)), latitude: Number(lat.toFixed(6)), address: formattedAddress };
        addMarker(lng, lat);
        map.value.setCenter([lng, lat]);
        map.value.setZoom(16);
        message.success(`找到地址: ${formattedAddress}`);
        searchForm.value.address = '';
      } else {
        message.error('未找到该地址，请检查地址是否正确或尝试更具体的地址');
      }
    });
  } catch (error) {
    searchLoading.value = false;
    message.error('地址搜索失败，请稍后重试');
    console.error('地址搜索出错:', error);
  }
};

// 确认位置
const confirmLocation = () => {
  if (!currentLocation.value) {
    message.warning('请先选择一个位置');
    return;
  }
  
  // 发出位置选择事件给父组件
  emit('locationSelected', currentLocation.value);
  
  message.success(`已确认位置: ${currentLocation.value.address || `${currentLocation.value.longitude}, ${currentLocation.value.latitude}`}`);
};

const clearMarker = () => {
  if (map.value) {
    map.value.clearMap();
    marker.value = null;
    currentLocation.value = null;
    searchForm.value.address = '';
    message.success('已清除所有标记点');
  }
};


onMounted(async () => {
  try {
    // 安全密钥配置
    (window as any)._AMapSecurityConfig = { securityJsCode: AMAP_SECURITY_CODE };
    AMapLoader.load({
      key: AMAP_KEY,
      version: '2.0',
      plugins: ['AMap.Geolocation', 'AMap.Geocoder', 'AMap.AutoComplete', 'AMap.PlaceSearch', 'AMap.Scale', 'AMap.ToolBar'],
      // 如需 AMapUI 或 Loca 可在此添加
    }).then((AMap) => {
      AMapRef.value = AMap;
      initMap();
      
      // 如果有初始位置，设置到地图上
      if (props.initialLocation) {
        currentLocation.value = props.initialLocation;
        addMarker(props.initialLocation.longitude, props.initialLocation.latitude);
        map.value.setCenter([props.initialLocation.longitude, props.initialLocation.latitude]);
        map.value.setZoom(16);
      }
    }).catch((error) => {
      message.error('地图加载失败，请检查网络连接和API密钥');
      console.error('地图加载错误:', error);
    });
  } catch (error) {
    message.error('地图加载失败，请检查网络连接和API密钥');
    console.error('地图加载错误:', error);
  }
});
</script>

<style scoped>
.map-container {
  position: relative;
  width: 100%;
  height: 100%;
  min-height: 500px;
  display: flex;
  flex-direction: column;
}

.modal-layout {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.modal-layout .map {
  height: 320px;
}

.modal-layout .control-panel {
  flex: 1;
  margin-top: 16px;
  overflow-y: auto;
  min-height: 0;
}

.map {
  width: 100%;
  height: 350px;
  border-radius: 8px;
  overflow: hidden;
}

.control-panel {
  margin-top: 16px;
  flex: 1;
  min-height: 0;
}

/* 当在模态框中使用时的特殊样式 */
:deep(.n-card.main-card) {
  height: 100%;
  display: flex;
  flex-direction: column;
}

:deep(.n-card.main-card .n-card__content) {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
}

/* 确保地图容器在模态框中能够正确显示 */
.map-wrapper {
  flex-shrink: 0;
}

/* 控制面板在模态框中的样式 */
:global(.n-modal) .control-panel {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
}
</style>
