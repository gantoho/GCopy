<script setup lang="ts">
import { ref, onMounted } from "vue";
import { GetContent, AddNote } from '@/../wailsjs/go/main/App'
import { NButton, NIcon, NModal, NInput, NScrollbar, NEmpty, useMessage } from 'naive-ui'
import { Add, HappyOutline } from '@vicons/ionicons5'
import Note from '@/components/Note.vue'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import { error } from "console";

const message = useMessage()

const notes = ref()
const getContent = async () => {
  notes.value = await GetContent()
}

onMounted(() => {
  getContent()
})

const showModal = ref(false)
const addNoteFn = async () => {
  const noteObj = {
    id: Date.now(),
    filename: fileName.value,
    content: textareaValue.value
  }
  const err = await AddNote(noteObj)
  if (err !== "ok") {
    message.error(`添加失败 ${err}`)
    return
  }
  message.info("add note ok")
  showModal.value = false
  getContent()
}

const fileName = ref()
const textareaValue = ref()
const editorOption= {
  modules: {
    toolbar: true // 工具栏隐藏显示
  }
}
const scrollbarMaxHeight = ref("150px")
</script>

<template>
  <div class="notes" v-if="notes">
    <Note v-for="note in notes" :key="note.id" :note="note">
      <template #default="{data}">
        <span class="filename">{{ data.filename }}</span>
        <n-scrollbar :style="{'max-height': scrollbarMaxHeight}">
          <pre class="content">{{ data.content }}</pre>
        </n-scrollbar>
      </template>
    </Note>
  </div>
  <n-empty v-else description="还没有任何笔记" size="small" style="justify-content: end; min-height: 200px;">
    <template #icon>
      <n-icon>
        <HappyOutline />
      </n-icon>
    </template>
    <template #extra>
      <n-button type="primary" secondary size="small" @click="showModal = true">
        写笔记
      </n-button>
    </template>
  </n-empty>
  <!--  添加按钮-->
  <n-button class="add-btn" type="primary" secondary circle @click="showModal = true">
    <n-icon>
      <Add/>
    </n-icon>
  </n-button>
  <n-modal
    v-model:show="showModal"
    preset="dialog"
    :show-icon="false"
    style="background-color: #3b3b3b;"
  >
    <template #header>
      <div style="color: #fff;">添加笔记</div>
    </template>
    <n-input type="text" placeholder="请输入文件名" v-model:value="fileName" />
    <QuillEditor :options="editorOption" v-model:content="textareaValue" contentType="text" />
    <template #action>
      <n-button type="primary" secondary @click="addNoteFn"> 添加 </n-button>
      <n-button type="primary" secondary ghost> 取消 </n-button>
    </template>
  </n-modal>
</template>

<style scoped lang="scss">
.notes {
  margin: 0 10px;
  .filename {
    display: inline-block;
    width: 100%;
    color: #0051ff;
    background-color: #4e3b3b;
    position: sticky;
    left: 0;
    top: 0;
    z-index: 1;
  }
  .content {
    background-color: #3b3b3b;
    color: #fff;
  }
}
.add-btn {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 99;
}
</style>
