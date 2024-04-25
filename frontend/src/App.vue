<script setup lang='ts'>
// import Home from '@/pages/Home.vue';
import Notes from '@/pages/Notes.vue'
import { NConfigProvider, NMessageProvider, NThemeEditor } from "naive-ui";
import theme from '@/utils/theme'
</script>

<template>
  <div class="app" style="--wails-draggable:no-drag">
<!--    <n-theme-editor>-->
      <n-config-provider :theme-overrides="theme">
        <n-message-provider>
          <Notes/>
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
