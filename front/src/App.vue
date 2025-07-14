<template>
  <n-message-provider>
    <Login/>
  </n-message-provider>
</template>

<script setup lang="ts">
import Login from "@/views/user/login.vue"
import { NMessageProvider } from 'naive-ui'
</script>