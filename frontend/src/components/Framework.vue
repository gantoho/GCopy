<script setup lang='ts'>
import {onMounted, ref} from 'vue'
import { ShowWin, Greet, NowTime } from '../../wailsjs/go/main/App.js'

const msg = ref()
const nowTime = ref()

onMounted(async () => {
  msg.value = await Greet('World')
  nowTime.value = await NowTime()
})

</script>

<template>
  <div class="framework">
    <h1>Framework</h1>
    <span>{{ msg }}</span>
    <p>{{ nowTime }}</p>
  </div>
</template>

<style lang='scss' scoped>
.framework {
  background-color: #3b3b3b;
}
</style>
