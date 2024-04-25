<script setup lang="ts">
import { onMounted, withDefaults } from 'vue'
const props = withDefaults(defineProps<{note: any}>(), {
  note: () => '默认值'
})
onMounted(() => {
  console.log(props.note)
})
</script>

<template>
  <slot :data="note"></slot>
</template>