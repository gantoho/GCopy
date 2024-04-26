<script setup lang="ts">
import { ref, onMounted, watchEffect } from "vue";
import { GetContent, AddNote, DeleteNote } from '@/../wailsjs/go/main/App'
import { WindowSetAlwaysOnTop, WindowSetPosition, WindowSetSystemDefaultTheme, WindowSetLightTheme , WindowSetDarkTheme } from '@/../wailsjs/runtime'
import { NButton, NIcon, NModal, NInput, NScrollbar, NEmpty, NPopconfirm, useMessage } from 'naive-ui'
import { Add, HappyOutline, Rocket, TrashBin, InvertMode, RadioButtonOff, RadioButtonOn } from '@vicons/ionicons5'
import Note from '@/components/Note.vue'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

const message = useMessage()

const notes = ref()
const getContent = async () => {
  notes.value = await GetContent()
}

onMounted(() => {
  getContent()
  WindowSetPosition(0, 0)
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

const windowTop = ref(false)
const handleWindowTop = () => {
  windowTop.value = !windowTop.value
  console.log(windowTop.value)
  WindowSetAlwaysOnTop(windowTop.value)
}

const deleteNoteFn = (fileName: string) => {
  DeleteNote(fileName)
  getContent()
}

const theme = ref("auto")
watchEffect(() => {
  if (theme.value === "auto") {
    WindowSetSystemDefaultTheme()
  } else if (theme.value === "light") {
    WindowSetLightTheme()
  } else if (theme.value === "dark") {
    WindowSetDarkTheme()
  }
})
const handleTheme = () => {
  if (theme.value === "auto") {
    theme.value = "light"
  } else if (theme.value === "light") {
    theme.value = "dark"
  } else {
    theme.value = "auto"
  }
}
</script>

<template>
  <div class="notes" v-if="notes">
    <Note v-for="note in notes" :key="note.id" :note="note">
      <template #default="{data}">
        <div class="title">
          <span class="filename">{{ data.filename }}</span>
          <n-popconfirm
            @positive-click="deleteNoteFn(data.filename)"
            :show-icon="false"
            style="background-color: #220; color: #9c9c9c;"
          >
            <template #trigger>
              <n-button type="error" size="small">
                <n-icon>
                  <TrashBin />
                </n-icon>
              </n-button>
            </template>
            确认删除吗?
          </n-popconfirm>
        </div>
        <n-scrollbar class="scrollbar" :style="{'max-height': scrollbarMaxHeight}">
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
  <div class="btns">
    <n-button class="theme-btn" type="primary" secondary circle @click="handleTheme">
      <n-icon>
        <InvertMode v-if="theme === 'auto'" />
        <RadioButtonOff v-if="theme === 'light'"/>
        <RadioButtonOn v-if="theme === 'dark'"/>
      </n-icon>
    </n-button>
    <n-button class="top-btn" type="primary" :ghost="!windowTop" secondary circle @click="handleWindowTop">
      <n-icon>
        <Rocket/>
      </n-icon>
    </n-button>
    <n-button class="add-btn" type="primary" secondary circle @click="showModal = true">
      <n-icon>
        <Add/>
      </n-icon>
    </n-button>
  </div>
  <n-modal
    v-model:show="showModal"
    preset="dialog"
    :show-icon="false"
    style="background-color: #3b3b3b; color: #9c9c9c; padding: 10px;"
  >
    <template #header>
      <div style="color: #fff;">添加笔记</div>
    </template>
    <n-input type="text" placeholder="请输入文件名" v-model:value="fileName" />
    <QuillEditor :options="editorOption" v-model:content="textareaValue" contentType="text" />
    <template #action>
      <n-button type="primary" secondary @click="addNoteFn"> 添加 </n-button>
      <n-button type="primary" secondary ghost @click="showModal = false"> 取消 </n-button>
    </template>
  </n-modal>
</template>

<style scoped lang="scss">
.notes {
  .title {
    display: flex;
    .filename {
      display: inline-block;
      width: 100%;
      color: #fb7299;
      background-color: #582b38;
      position: sticky;
      left: 0;
      top: 0;
      z-index: 1;
      padding: 2px 4px;
      border-radius: 10px;
    }
    &+.content {
      margin-top: 10px;
    }
  }
  .scrollbar {
    .content {
      background-color: #3b3b3b;
      color: #fff;
      border-radius: 10px;
      padding: 5px;
    }
  }
}
.btns {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 99;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
