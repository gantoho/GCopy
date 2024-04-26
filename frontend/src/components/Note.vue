<script setup lang="ts">
import { withDefaults } from 'vue'
withDefaults(defineProps<{note: any}>(), {
  note: () => '默认值'
})
</script>

<template>
  <slot :data="note"></slot>
</template>
