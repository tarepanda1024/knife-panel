<template>
  <div>
    <div class="top-bar">
      <div class="top-path">
        <ion-icon
            v-if="currentPath !== '/'"
            @click="handleListFiles(indicators[indicators.length-2].path)"
            name="return-up-back"
            class="ivu-icon"
            style="margin-right: 5px;cursor: pointer">
        </ion-icon>
        <template v-for="(item,i) in indicators">
          <span v-if="i>1">/ </span>
          <span style="cursor: pointer" @click="handleListFiles(item.path)">{{ item.name }}
            </span>
        </template>
      </div>
      <div class="top-menu">
        <Button type="primary">删除</Button>
        <Button type="primary">下载</Button>
        <Button type="primary">上传</Button>
      </div>
    </div>

    <Table :columns="fileColumns"
           :data="files"
           :height="tableHeight"
           @on-row-click="handleRowClick"
    >

    </Table>
  </div>
</template>

<script>

import {listFiles} from "@/service/FileService";
import {formatSize} from "@/utils/FileSizeUtils"
export default {
  name: "FileManage",
  data() {
    return {
      fileColumns: [
        {
          "title": "文件名",
          "key": "name",
          "render": (h, params) => {
            let icon = params.row.dir ? 'ios-folder' : 'ios-document'
            return h('div', [
              h('Icon', {
                props: {
                  type: icon,
                },
                style: {
                  marginRight: '5px',
                }
              }),
              h('span', params.row.name)
            ])
          }
        },
        {
          "title": "大小",
          "key": "size",
          "width": 120,
          "align": "center",
          "render": (h, params) => {
            return h('span', formatSize( params.row.size))
          }
        },
        {
          "title": "修改时间",
          "key": "modifyTime",
          "width": 200,
          "align": "center",
          "render": (h, params) => {
            let time = params.row.modifyTime;
            let value = this.$moment.unix(time).format('YYYY-MM-DD HH:mm:ss')
            return h('span', value + ' ')
          }
        }
      ],
      files: [],
      tableHeight: 600,
      indicators: [],
      currentPath: ''
    }
  },
  methods: {
    initStyle() {
      this.tableHeight = document.body.clientHeight - 200;
    },
    handleBasePathChange(basePath) {
      let paths = basePath.split('/')
      let absPath = ''
      this.indicators = [{
        name: '/',
        path: '/'
      }]
      paths.forEach(value => {
        if (value !== '') {
          absPath = absPath + '/' + value;
          this.indicators.push({
            name: value,
            path: absPath
          })
        }
      })
    },
    handleRowClick(row, index) {
      if (row.dir) {
        this.handleListFiles(row.absPath)
      }
    },
    handleListFiles(basePath) {
      listFiles({"basePath": basePath}).then(data => {
        this.currentPath = basePath;
        this.files = data;
        this.handleBasePathChange(basePath);
      })
    }
  },
  mounted() {
    this.initStyle();
    this.handleListFiles("/")
  }
}
</script>

<style scoped>
.top-bar {
  height: 40px;
}

.top-path {
  float: left;
}

.top-menu {
  width: 200px;
  display: flex;
  justify-content: space-around;
  float: right;
}
</style>