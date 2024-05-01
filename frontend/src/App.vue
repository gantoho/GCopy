<script setup lang='ts'>
import { ref } from 'vue'
// import Home from '@/pages/Home.vue';
import Notes from '@/pages/Notes.vue'
import { NConfigProvider, NMessageProvider, zhCN, NGlobalStyle } from "naive-ui";
import type { GlobalTheme } from 'naive-ui'
import themeOverrides from '@/utils/theme'

const theme = ref<GlobalTheme | null>(null)
</script>

<template>
  <div class="app" style="--wails-draggable:no-drag">
    <!--    <n-theme-editor>-->
    <n-config-provider :theme="theme" :locale="zhCN" :theme-overrides="themeOverrides">
      <n-global-style/>
      <n-message-provider>
        <Notes v-model:theme="theme"/>
      </n-message-provider>
    </n-config-provider>
    <!--    </n-theme-editor>-->
  </div>
</template>

<style lang='scss' scoped>
.app {
  color: #bcbec4;
  i {
    font-style: normal;
  }
}
</style>
