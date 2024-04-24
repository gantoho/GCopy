<script setup lang='ts'>
import { ref } from 'vue'

defineProps(['item'])

const copyContent = async () => {
  const text = await navigator.clipboard.readText()
  console.log(text, "-----------")
}
</script>

<template>
  <div class="item" @click="copyContent">
    {{ item }}
  </div>
</template>

<style lang='scss' scoped>
.item {
  cursor: pointer;
}
</style>
