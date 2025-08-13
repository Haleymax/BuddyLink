<template>
  <div class="map-container">
    <n-card title="高德地图定位" :bordered="false">
      <!-- 使用提示 -->
      <n-alert type="info" :show-icon="false" style="margin-bottom: 16px;">
        💡 <strong>操作提示:</strong> 支持地址搜索、实时搜索建议、附近POI查找，点击地图获取坐标，或手动输入经纬度定位
      </n-alert>
      
      <!-- 地图容器 -->
      <div id="amap-container" class="map"></div>
      
      <!-- 控制面板 -->
      <div class="control-panel">
        <n-space vertical>
          <!-- 获取当前位置按钮 -->
          <n-button 
            type="primary" 
            @click="getCurrentLocation"
            :loading="locationLoading"
            block
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
          <n-form :model="searchForm" label-placement="left" label-width="60">
            <n-form-item label="地址">
              <n-input
                v-model:value="searchForm.address" 
                placeholder="请输入地址、POI名称或关键词"
                clearable
                @keyup.enter="searchAddress"
              />
            </n-form-item>
            <n-form-item>
              <n-space>
                <n-button 
                  type="success" 
                  @click="searchAddress"
                  :disabled="!searchForm.address || searchForm.address.trim() === ''"
                  :loading="searchLoading"
                  style="flex: 1"
                >
                  <template #icon>
                    <n-icon>
                      <svg viewBox="0 0 24 24">
                        <path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"/>
                      </svg>
                    </n-icon>
                  </template>
                  搜索地址
                </n-button>
                <n-button 
                  type="info" 
                  @click="searchNearby"
                  :disabled="!currentLocation"
                  quaternary
                >
                  <template #icon>
                    <n-icon>
                      <svg viewBox="0 0 24 24">
                        <path fill="currentColor" d="M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4Z"/>
                      </svg>
                    </n-icon>
                  </template>
                  附近
                </n-button>
              </n-space>
            </n-form-item>
          </n-form>

          <!-- 分隔线 -->
          <n-divider>或</n-divider>

          <!-- 经纬度输入 -->
          <n-form :model="coordinateForm" label-placement="left" label-width="60">
            <n-form-item label="经度">
              <n-input-number 
                v-model:value="coordinateForm.longitude" 
                placeholder="请输入经度"
                :precision="6"
                :step="0.000001"
                style="width: 100%"
                clearable
              />
            </n-form-item>
            <n-form-item label="纬度">
              <n-input-number 
                v-model:value="coordinateForm.latitude" 
                placeholder="请输入纬度"
                :precision="6"
                :step="0.000001"
                style="width: 100%"
                clearable
              />
            </n-form-item>
            <n-form-item>
              <n-space>
                <n-button 
                  type="info" 
                  @click="navigateToCoordinate"
                  :disabled="!isValidCoordinate"
                  style="flex: 1"
                >
                  定位到指定坐标
                </n-button>
                <n-button 
                  type="warning" 
                  @click="clearMarker"
                  quaternary
                >
                  清除标记
                </n-button>
              </n-space>
            </n-form-item>
          </n-form>

          <!-- 当前位置信息显示 -->
          <n-card v-if="currentLocation" size="small" title="当前位置信息">
            <n-descriptions :column="1" size="small">
              <n-descriptions-item label="经度">
                {{ currentLocation.longitude }}
              </n-descriptions-item>
              <n-descriptions-item label="纬度">
                {{ currentLocation.latitude }}
              </n-descriptions-item>
              <n-descriptions-item label="地址" v-if="currentLocation.address">
                {{ currentLocation.address }}
              </n-descriptions-item>
            </n-descriptions>
          </n-card>
        </n-space>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useMessage } from 'naive-ui';
import '../styles/MapComponent.css';

// 声明全局AMap变量
declare global {
  interface Window {
    AMap: any;
    _AMapSecurityConfig: any;
  }
}

const message = useMessage();

// 响应式数据
const map = ref<any>(null);
const marker = ref<any>(null);
const locationLoading = ref(false);
const searchLoading = ref(false);
const currentLocation = ref<{
  longitude: number;
  latitude: number;
  address?: string;
} | null>(null);

const coordinateForm = ref({
  longitude: null as number | null,
  latitude: null as number | null
});

const searchForm = ref({
  address: ''
});

// 计算属性：验证坐标是否有效
const isValidCoordinate = computed(() => {
  return coordinateForm.value.longitude !== null && 
         coordinateForm.value.latitude !== null &&
         coordinateForm.value.longitude >= -180 && 
         coordinateForm.value.longitude <= 180 &&
         coordinateForm.value.latitude >= -90 && 
         coordinateForm.value.latitude <= 90;
});

// 高德地图配置
const AMAP_KEY = import.meta.env.VITE_AMAP_KEY || '42707d19daa52635acb92b215df96bcc'; // 请在.env.local中配置或直接替换
const AMAP_SECURITY_CODE = import.meta.env.VITE_AMAP_SECURITY_CODE || '45a9990b03da96393396d53446d5eb6e'; // 请在.env.local中配置或直接替换

// 初始化地图
const initMap = () => {
  try {
    // 设置安全密钥
    window._AMapSecurityConfig = {
      securityJsCode: AMAP_SECURITY_CODE,
    };

    // 创建地图实例
    map.value = new window.AMap.Map('amap-container', {
      zoom: 13,
      center: [116.39, 39.9], // 默认中心点（北京）
      mapStyle: 'amap://styles/normal',
      viewMode: '3D'
    });

    // 等待地图加载完成后再添加控件
    map.value.on('complete', () => {
      try {
        // 尝试添加比例尺控件
        if (window.AMap.Scale) {
          const scale = new window.AMap.Scale();
          map.value.addControl(scale);
          console.log('比例尺控件加载成功');
        }
        
        // 尝试添加工具栏控件
        if (window.AMap.ToolBar) {
          const toolBar = new window.AMap.ToolBar();
          map.value.addControl(toolBar);
          console.log('工具栏控件加载成功');
        }
      } catch (error) {
        console.warn('地图控件加载失败:', error);
        // 控件加载失败不影响地图基本功能
      }
    });
    
    // 添加地图点击事件监听
    map.value.on('click', (e: any) => {
      const { lng, lat } = e.lnglat;
      
      // 更新坐标表单
      coordinateForm.value.longitude = Number(lng.toFixed(6));
      coordinateForm.value.latitude = Number(lat.toFixed(6));
      
      // 添加标记点
      addMarker(lng, lat);
      
      // 获取地址信息（可选）
      const geocoder = new window.AMap.Geocoder();
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

// 加载高德地图脚本
const loadAmapScript = () => {
  return new Promise((resolve, reject) => {
    if (window.AMap) {
      resolve(window.AMap);
      return;
    }

    const script = document.createElement('script');
    script.src = `https://webapi.amap.com/maps?v=2.0&key=${AMAP_KEY}&plugin=AMap.Geolocation,AMap.Geocoder,AMap.AutoComplete,AMap.PlaceSearch`;
    script.async = true;
    script.onload = () => resolve(window.AMap);
    script.onerror = reject;
    document.head.appendChild(script);
  });
};

// 获取当前位置
const getCurrentLocation = async () => {
  if (!map.value) {
    message.error('地图未初始化');
    return;
  }

  locationLoading.value = true;

  try {
    const geolocation = new window.AMap.Geolocation({
      enableHighAccuracy: true,
      timeout: 10000,
      maximumAge: 1000 * 60 * 60 * 24,
      convert: true,
      showButton: false,
      buttonPosition: 'RB',
      buttonOffset: new window.AMap.Pixel(10, 20),
      showMarker: false,
      showCircle: false,
      panToLocation: true,
      zoomToAccuracy: true
    });

    geolocation.getCurrentPosition((status: string, result: any) => {
      locationLoading.value = false;
      
      if (status === 'complete') {
        const { lng, lat, formattedAddress } = result.position;
        
        currentLocation.value = {
          longitude: lng,
          latitude: lat,
          address: formattedAddress
        };

        // 更新输入框的值
        coordinateForm.value.longitude = Number(lng.toFixed(6));
        coordinateForm.value.latitude = Number(lat.toFixed(6));

        // 在地图上标记位置
        addMarker(lng, lat);
        
        // 移动地图中心到当前位置
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
  // 移除已有标记
  if (marker.value) {
    map.value.remove(marker.value);
  }

  // 创建新标记
  marker.value = new window.AMap.Marker({
    position: [lng, lat],
    icon: new window.AMap.Icon({
      size: new window.AMap.Size(32, 32),
      image: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_r.png',
      imageOffset: new window.AMap.Pixel(-16, -32),
      imageSize: new window.AMap.Size(32, 32)
    }),
    title: `坐标: ${lng.toFixed(6)}, ${lat.toFixed(6)}`,
    anchor: 'bottom-center'
  });

  // 添加到地图
  map.value.add(marker.value);
  
  // 创建信息窗体
  const infoWindow = new window.AMap.InfoWindow({
    content: `
      <div style="padding: 8px; min-width: 200px;">
        <div style="font-weight: bold; margin-bottom: 8px; color: #333;">📍 位置坐标</div>
        <div style="margin-bottom: 4px;"><strong>经度:</strong> ${lng.toFixed(6)}</div>
        <div style="margin-bottom: 4px;"><strong>纬度:</strong> ${lat.toFixed(6)}</div>
      </div>
    `,
    offset: new window.AMap.Pixel(0, -30),
    closeWhenClickMap: true
  });

  // 点击标记显示信息窗体
  marker.value.on('click', () => {
    infoWindow.open(map.value, [lng, lat]);
  });
};

// 定位到指定坐标
const navigateToCoordinate = async () => {
  if (!isValidCoordinate.value || !map.value) {
    return;
  }

  const { longitude, latitude } = coordinateForm.value;
  
  try {
    // 移动地图到指定位置
    map.value.setCenter([longitude!, latitude!]);
    map.value.setZoom(16);
    
    // 添加标记
    addMarker(longitude!, latitude!);
    
    // 逆地理编码获取地址信息
    const geocoder = new window.AMap.Geocoder();
    geocoder.getAddress([longitude!, latitude!], (status: string, result: any) => {
      if (status === 'complete' && result.regeocode) {
        currentLocation.value = {
          longitude: longitude!,
          latitude: latitude!,
          address: result.regeocode.formattedAddress
        };
      } else {
        currentLocation.value = {
          longitude: longitude!,
          latitude: latitude!
        };
      }
    });
    
    message.success('定位成功');
  } catch (error) {
    message.error('定位失败');
    console.error('定位错误:', error);
  }
};

// 搜索附近POI
const searchNearby = () => {
  if (!currentLocation.value || !map.value) {
    message.warning('请先获取当前位置');
    return;
  }

  const { longitude, latitude } = currentLocation.value;
  
  // 使用高德地图PlaceSearch服务搜索附近POI
  const placeSearch = new window.AMap.PlaceSearch({
    pageSize: 10,
    pageIndex: 1,
    city: '',
    map: map.value,
    panel: false
  });

  placeSearch.searchNearBy('', [longitude, latitude], 1000, (status: string, result: any) => {
    if (status === 'complete' && result.poiList && result.poiList.pois.length > 0) {
      // 清除之前的标记
      if (marker.value) {
        map.value.remove(marker.value);
      }

      // 添加附近POI标记
      const pois = result.poiList.pois.slice(0, 5); // 显示前5个POI
      const markers: any[] = [];

      pois.forEach((poi: any, index: number) => {
        const poiMarker = new window.AMap.Marker({
          position: [poi.location.lng, poi.location.lat],
          icon: new window.AMap.Icon({
            size: new window.AMap.Size(25, 30),
            image: `https://webapi.amap.com/theme/v1.3/markers/n/mark_${String.fromCharCode(65 + index)}.png`
          }),
          title: poi.name
        });

        // 创建信息窗体
        const infoWindow = new window.AMap.InfoWindow({
          content: `
            <div style="padding: 8px; min-width: 200px;">
              <div style="font-weight: bold; margin-bottom: 8px; color: #333;">📍 ${poi.name}</div>
              <div style="margin-bottom: 4px;"><strong>地址:</strong> ${poi.address || '暂无地址'}</div>
              <div style="margin-bottom: 4px;"><strong>类型:</strong> ${poi.type || '未知'}</div>
              <div><strong>距离:</strong> ${Math.round(poi.distance)}米</div>
            </div>
          `,
          offset: new window.AMap.Pixel(0, -30)
        });

        poiMarker.on('click', () => {
          infoWindow.open(map.value, [poi.location.lng, poi.location.lat]);
        });

        markers.push(poiMarker);
        map.value.add(poiMarker);
      });

      // 调整地图视野以包含所有POI
      const bounds = new window.AMap.Bounds();
      pois.forEach((poi: any) => {
        bounds.extend([poi.location.lng, poi.location.lat]);
      });
      map.value.setBounds(bounds);

      message.success(`找到${pois.length}个附近地点`);
    } else {
      message.warning('附近没有找到相关地点');
    }
  });
};

// 地址搜索功能
const searchAddress = async () => {
  if (!searchForm.value.address || searchForm.value.address.trim() === '' || !map.value) {
    message.warning('请输入有效的地址');
    return;
  }

  searchLoading.value = true;
  console.log('开始搜索地址:', searchForm.value.address);

  try {
    // 确保AMap.Geocoder已加载
    if (!window.AMap || !window.AMap.Geocoder) {
      message.error('地图服务未完全加载，请稍后再试');
      searchLoading.value = false;
      return;
    }

    const geocoder = new window.AMap.Geocoder({
      city: '', // 全国范围搜索
      radius: 1000, // 搜索半径
      extensions: 'base' // 返回基本信息
    });

    geocoder.getLocation(searchForm.value.address, (status: string, result: any) => {
      searchLoading.value = false;
      console.log('搜索结果:', status, result);
      
      if (status === 'complete' && result.geocodes && result.geocodes.length > 0) {
        const geocode = result.geocodes[0];
        const location = geocode.location;
        const lng = location.lng;
        const lat = location.lat;
        const formattedAddress = geocode.formattedAddress || geocode.district || searchForm.value.address;
        
        console.log('找到位置:', lng, lat, formattedAddress);
        
        // 更新坐标表单
        coordinateForm.value.longitude = Number(lng.toFixed(6));
        coordinateForm.value.latitude = Number(lat.toFixed(6));
        
        // 更新当前位置信息
        currentLocation.value = {
          longitude: Number(lng.toFixed(6)),
          latitude: Number(lat.toFixed(6)),
          address: formattedAddress
        };
        
        // 添加标记点
        addMarker(lng, lat);
        
        // 移动地图中心到搜索位置
        map.value.setCenter([lng, lat]);
        map.value.setZoom(16);
        
        message.success(`找到地址: ${formattedAddress}`);
        
        // 搜索成功后清空搜索框
        searchForm.value.address = '';
      } else {
        console.error('搜索失败:', status, result);
        message.error('未找到该地址，请检查地址是否正确或尝试更具体的地址');
      }
    });
  } catch (error) {
    searchLoading.value = false;
    console.error('地址搜索出错:', error);
    message.error('地址搜索失败，请稍后重试');
  }
};

// 清除标记
const clearMarker = () => {
  if (map.value) {
    // 清除所有标记
    map.value.clearMap();
    marker.value = null;
    
    // 清空位置信息
    currentLocation.value = null;
    
    // 清空坐标表单
    coordinateForm.value.longitude = null;
    coordinateForm.value.latitude = null;
    
    // 清空搜索框
    searchForm.value.address = '';
    
    message.success('已清除所有标记点');
  }
};

// 监听坐标输入变化
watch([() => coordinateForm.value.longitude, () => coordinateForm.value.latitude], 
  () => {
    // 当坐标发生变化且有效时，自动更新地图（可选）
    // 这里暂时注释，避免输入时频繁更新
    // if (isValidCoordinate.value) {
    //   navigateToCoordinate();
    // }
  }
);

// 组件挂载时初始化
onMounted(async () => {
  try {
    await loadAmapScript();
    initMap();
  } catch (error) {
    message.error('地图加载失败，请检查网络连接和API密钥');
    console.error('地图加载错误:', error);
  }
});
</script>
