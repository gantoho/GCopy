<script setup lang='ts'>
import { ref } from 'vue'
import Framework from '@/components/Framework.vue';
import Item from '@/components/Item.vue'
import { EventsOn, BrowserOpenURL, ClipboardGetText, ClipboardSetText, EventsEmit } from '../../wailsjs/runtime/runtime'

// ClipboardSetText('Hello World').then(async () => {
//   await ClipboardGetText()
// })

EventsOn('Copy', async (text) => {
  console.log('Copy: ', text)
})

// BrowserOpenURL('https://ganto.me')
</script>

<template>
  <div class="home">
    <Framework />
    <template v-for="item in [1,2,3]" :key="item">
      <Item :item/>
    </template>
  </div>
</template>

<style lang='scss' scoped>

</style>
