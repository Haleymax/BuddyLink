<template>
  <div class="map-page">
    <n-space vertical :size="24">
      <n-page-header 
        title="地图定位" 
        subtitle="使用高德地图进行位置定位和导航"
      >
        <template #avatar>
          <n-icon size="32" color="#18a058">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M12,2C15.31,2 18,4.66 18,7.95C18,12.41 12,19 12,19S6,12.41 6,7.95C6,4.66 8.69,2 12,2M12,6A2,2 0 0,0 10,8A2,2 0 0,0 12,10A2,2 0 0,0 14,8A2,2 0 0,0 12,6Z"/>
            </svg>
          </n-icon>
        </template>
      </n-page-header>
      
      <MapComponent />
    </n-space>
  </div>
</template>

<script lang="ts" setup>
import MapComponent from '../../components/MapComponent.vue';
import '../../styles/MapView.css';
</script>
