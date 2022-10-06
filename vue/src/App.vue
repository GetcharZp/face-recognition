<template>
  <div class="page">
    <div class="video-box">
      <video id="video" width="800" height="600" preload autoplay loop muted></video>
      <canvas id="canvas" width="800" height="600"></canvas>
      <canvas width="800" height="600" ref="photo" class="photo"></canvas>
    </div>
    <div class="btn">
      <el-button @click="dialogVisible = true">录入人脸</el-button>
      <el-button @click="draw">人脸识别</el-button>
    </div>
  </div>
  <el-dialog
      v-model="dialogVisible"
      title="录入人脸数据"
      width="30%"
  >
    <el-input v-model="name" placeholder="请输入用户姓名" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="enterFaceData">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import api from '@/api/api.js'
import {onMounted, ref} from 'vue'
import { ElMessage } from 'element-plus'

let canvas
let video
const photo=ref()
const dialogVisible = ref(false)
const name = ref('')

onMounted(()=>{
  video = document.getElementById('video');
  canvas = document.getElementById('canvas');
  let context = canvas.getContext('2d');
  let img = new Image();

  let tracker = new tracking.ObjectTracker('face');
  tracker.setInitialScale(4);
  tracker.setStepSize(2);
  tracker.setEdgesDensity(0.1);

  tracking.track('#video', tracker, { camera: true });

  tracker.on('track', function(event) {
    context.clearRect(0, 0, canvas.width, canvas.height);

    event.data.forEach(function(rect) {
      context.strokeStyle = '#a64ceb';
      context.strokeRect(rect.x, rect.y, rect.width, rect.height);
      context.font = '11px Helvetica';
      context.fillStyle = "#fff";
      context.fillText('x: ' + rect.x + 'px', rect.x + rect.width + 5, rect.y + 11);
      context.fillText('y: ' + rect.y + 'px', rect.x + rect.width + 5, rect.y + 22);
    });
  });
})
const enterFaceData=()=> {
  if (name.value === '') {
    ElMessage.error('姓名不能为空')
    return
  }
  let ctx=photo.value.getContext('2d')
  ctx.drawImage(video,0 ,0, 800,600)
  api.enterFaceData({
    name: name.value,
    file:convertBase64UrlToBlob()
  }).then(res=>{
    if (res.data.code === 200) {
      ElMessage.success(res.data.msg)
    } else {
      ElMessage.error(res.data.msg)
    }
  })
}
const convertBase64UrlToBlob=()=>{
  var img = new Image();
  img.src = photo.value.toDataURL("image/jpeg", 0.8);   //转换为图片形式  打印出二进制数据
  let urlData=img.src
  var bytes=window.atob(urlData.split(',')[1]);        //去掉url的头，并转换为byte
  //处理异常,将ascii码小于0的转换为大于0
  var ab = new ArrayBuffer(bytes.length);
  var ia = new Uint8Array(ab);
  for (var i = 0; i < bytes.length; i++) {
    ia[i] = bytes.charCodeAt(i);
  }
  return new Blob( [ab] , {type : 'image/jpg'});
}
</script>

<style scoped>
video, canvas {
  position: absolute;
}
button {
  /*margin-top: 600px;*/
}
.photo{
  position:absolute;
  right:0;
  bottom:0;
  transform-origin: right bottom;
  transform: scale(0.1);
}
html,body,#app,.page{
  width: 100%;
  height: 100%;

}
.video-box{
  width: 800px;
  height: 600px;
  margin: 0 auto;
}
.btn{
  text-align: center;
  padding-top: 20px;
}
</style>
