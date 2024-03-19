<script setup lang='ts'>
import { ref } from 'vue'

defineProps(['item'])
</script>

<template>
  <div class="item">
    {{ item }}
  </div>
</template>

<style lang='scss' scoped>

</style>
