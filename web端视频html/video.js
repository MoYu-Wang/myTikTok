//项目服务器IP和端口
// var myURL = "http://124.221.123.251:11316";

//我本地IP和端口
var myURL = "http://192.168.56.1:11316";

//初始化
var s = document.getElementById("s");
//let arr = ["http://vjs.zencdn.net/v/oceans.mp4", "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4","https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4","https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218093206z8V1JuPlpe.mp4"]
var arr = [];
let current = 0;

//频道英文映射表
var myMap = new Map();
myMap.set("热点","top");
myMap.set("关注","care");
myMap.set("体育","sport");
myMap.set("游戏","game");
myMap.set("推荐","referee");


//处理发送的请求(未实现)
function getRequest(url){

    var xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState == 4 ){
            arr = [];
            if (xhr.status == 200) {
                // 处理成功的响应
                console.log("服务器响应:\n"+xhr.responseText);
                strArray = translateData(xhr.responseText);
                updateArray(strArray);
            }else{
                console.log("服务器未响应,http状态码: " + xhr.status + " 请求当前状态: " + xhr.readyState);
                alert("请求超时,服务器未响应")
                return;
            }
        }
    };
    xhr.send();
}


//解析成功返回的数据(未实现)
function translateData(data){
    var strArray = new Array();
    strArray = data.split('\n');
    return strArray;
}


//更新数组
function updateArray(strArray){
    current = 0;
    for(var i = 0;i < strArray.length; i++){
        //arr.unshift(strArray[i]);
        arr.push(strArray[i]);
    }
    current = (current + arr.length)%arr.length;
    document.getElementById("video").innerHTML = `
        <source id="s" src="${arr[current]}" type="video/mp4">
    `;
    document.getElementById("video").load();
}


//监听事件
//监听键盘
window.onkeydown = function (e)
{
    console.log(e.keyCode);
    if (e.keyCode == 38) {
        //上
        if(arr.length == 0)return;
        current = (current - 1 + arr.length)%arr.length;
        console.log(current);
        document.getElementById("video").innerHTML = `
            <source id="s" src="${arr[current]}" type="video/mp4">
        `;
        document.getElementById("video").load();
    }
    if(e.keyCode == 40){
        //下
        if(arr.length == 0)return;
        current = (current + 1 + arr.length)%arr.length;
        console.log(current);
        document.getElementById("video").innerHTML = `
            <source id="s" src="${arr[current]}" type="video/mp4">
        `;
        document.getElementById("video").load();
    }
}

//初始加载事件
//频道显示
//视频切换源
window.onload = function(){
    var header = document.querySelector('.taghead')
    var lis = header.querySelectorAll('li')
    for(var i = 0;i < lis.length;i++){
        lis[i].setAttribute('index',i)
        lis[i].onclick=function(){
            for(var j = 0;j < lis.length;j++){
                lis[j].className = '' 
            }
            this.className = 'tag'
            var index = this.getAttribute('index')
            getRequest(myURL + "/" + myMap.get(this.textContent));
            console.log(myURL + "/" + myMap.get(this.textContent));
        }
    }
    //默认发送热点频道请求
    getRequest(myURL + "/" + myMap.get("热点"));
    //加载视频
    if(arr.length == 0) return;
    current = (current + arr.length)%arr.length;
    document.getElementById("video").innerHTML = `
        <source id="s" src="${arr[current]}" type="video/mp4">
    `;
    document.getElementById("video").load();
}

