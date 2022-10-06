import http from './http.js'

 
export default{
	//  get请求示例
	isWorker(param){
		return http.get(`/h5/meeting/person/check`,param)
	},
	// 录入人脸
	enterFaceData(param){
		return http.upFile('enter/face-data',param)
	},
	
	
 
	 
}
