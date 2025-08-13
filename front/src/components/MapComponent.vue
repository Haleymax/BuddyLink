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
import AMapLoader from '@amap/amap-jsapi-loader';

const message = useMessage();

// 响应式数据
const map = ref<any>(null);
const marker = ref<any>(null);
const AMapRef = ref<any>(null); // 保存 AMap 命名空间
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
      coordinateForm.value.longitude = Number(lng.toFixed(6));
      coordinateForm.value.latitude = Number(lat.toFixed(6));
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
        coordinateForm.value.longitude = Number(lng.toFixed(6));
        coordinateForm.value.latitude = Number(lat.toFixed(6));
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

// 定位到指定坐标
const navigateToCoordinate = async () => {
  if (!isValidCoordinate.value || !map.value || !AMapRef.value) return;
  const { longitude, latitude } = coordinateForm.value;
  try {
    map.value.setCenter([longitude!, latitude!]);
    map.value.setZoom(16);
    addMarker(longitude!, latitude!);
    const geocoder = new AMapRef.value.Geocoder();
    geocoder.getAddress([longitude!, latitude!], (status: string, result: any) => {
      if (status === 'complete' && result.regeocode) {
        currentLocation.value = { longitude: longitude!, latitude: latitude!, address: result.regeocode.formattedAddress };
      } else {
        currentLocation.value = { longitude: longitude!, latitude: latitude! };
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
  if (!currentLocation.value || !map.value || !AMapRef.value) {
    message.warning('请先获取当前位置');
    return;
  }
  const AMap = AMapRef.value;
  const { longitude, latitude } = currentLocation.value;
  const placeSearch = new AMap.PlaceSearch({ pageSize: 10, pageIndex: 1, city: '', map: map.value, panel: false });
  placeSearch.searchNearBy('', [longitude, latitude], 1000, (status: string, result: any) => {
    if (status === 'complete' && result.poiList && result.poiList.pois.length > 0) {
      if (marker.value) map.value.remove(marker.value);
      const pois = result.poiList.pois.slice(0, 5);
      const markers: any[] = [];
      pois.forEach((poi: any, index: number) => {
        const poiMarker = new AMap.Marker({
          position: [poi.location.lng, poi.location.lat],
          icon: new AMap.Icon({ size: new AMap.Size(25, 30), image: `https://webapi.amap.com/theme/v1.3/markers/n/mark_${String.fromCharCode(65 + index)}.png` }),
          title: poi.name
        });
        const infoWindow = new AMap.InfoWindow({
          content: `
            <div style="padding: 8px; min-width: 200px;">
              <div style="font-weight: bold; margin-bottom: 8px; color: #333;">📍 ${poi.name}</div>
              <div style="margin-bottom: 4px;"><strong>地址:</strong> ${poi.address || '暂无地址'}</div>
              <div style="margin-bottom: 4px;"><strong>类型:</strong> ${poi.type || '未知'}</div>
              <div><strong>距离:</strong> ${Math.round(poi.distance)}米</div>
            </div>
          `,
          offset: new AMap.Pixel(0, -30)
        });
        poiMarker.on('click', () => {
          infoWindow.open(map.value, [poi.location.lng, poi.location.lat]);
        });
        markers.push(poiMarker);
        map.value.add(poiMarker);
      });
      const bounds = new AMap.Bounds();
      pois.forEach((poi: any) => bounds.extend([poi.location.lng, poi.location.lat]));
      map.value.setBounds(bounds);
      message.success(`找到${pois.length}个附近地点`);
    } else {
      message.warning('附近没有找到相关地点');
    }
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
        coordinateForm.value.longitude = Number(lng.toFixed(6));
        coordinateForm.value.latitude = Number(lat.toFixed(6));
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

// 清除标记
const clearMarker = () => {
  if (map.value) {
    map.value.clearMap();
    marker.value = null;
    currentLocation.value = null;
    coordinateForm.value.longitude = null;
    coordinateForm.value.latitude = null;
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
}

.map {
  width: 100%;
  height: 400px;
  border-radius: 8px;
  overflow: hidden;
}

.control-panel {
  margin-top: 16px;
}
</style>
