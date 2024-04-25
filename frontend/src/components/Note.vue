<script setup lang="ts">
import { onMounted, withDefaults } from 'vue'
withDefaults(defineProps<{note: any}>(), {
  note: () => '默认值'
})
</script>

<template>
  <slot :data="note"></slot>
</template>
