var index = 0;
var videoInfos = [];
let current = 0;
var userData = {};
var videoWatchTime = 0;//视频播放时间(ms)
var isfav = 0;//点赞操作缓存
var comNum = 0;//评论数量缓存
var comTexts = [];//评论文本缓存
var listIndex = 1;
var listValue = ["","top","care","referee","search"];

//初始加载事件
//频道显示
//视频切换源
window.onload = function(){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    console.log(userData)
    //判断是否登录
    if (JSON.stringify(userData) !== "{}") {
        //判断用户登录信息是否过期
        GET_Req("/user/updatetoken","token",userData.token)
        .then(data => {
            if(data.status_code != 0){
                alert(data.status_msg);
                return
            }
            console.log(data);
            userData.token = data.token;
        })
        .catch(error => {
            console.error('Error:', error);
        });
        //获取基本信息
        GET_Req("/user/base","token" , userData.token)
        .then(data => {
            if(data.status_code != 0){
                alert(data.status_msg);
                return
            }
            console.log(data);
            userData.userID = data.userID
            userData.userName = data.userName
            userData.iphoneID = data.iphoneID
            document.getElementById("user_name").innerText = data.userName;
            localStorage.setItem("userData", JSON.stringify(userData));
        })
        .catch(error => {
            console.error('Error:', error);
        });
        //
        document.getElementById("userlogin").value = "退出登录";
        document.getElementById("userlogin").setAttribute("onclick", "UserExit()");
    } else {
        //游客登录的情况
        document.getElementById("user_name").innerText = "未登录"

        document.getElementById("userlogin").value = "登录账号";
        document.getElementById("userlogin").setAttribute("onclick", "UserLogin()");
    }
    

    //默认进入热点页面
    document.getElementById("topVideo").click();
}


//右上角用户信息
const userAvatar = document.getElementById('user_avatar');
const userMenu = document.getElementById('user_menu');

userAvatar.addEventListener('click', (event) => {
    userMenu.classList.toggle('active');
    event.stopPropagation(); // 阻止事件冒泡
});

userMenu.addEventListener('click', (event) => {
    event.stopPropagation(); // 阻止事件冒泡
});

document.addEventListener('click', () => {
    userMenu.classList.remove('active');
});


//个人中心
function UserInfo(){
    if(!UserIsLogin()){
        alert("用户未登录")
        return 
    }
}

//视频初始化
function initVideo(){
    index = 0;
    videoInfos = [];
}


//用户登录页面跳转
function UserLogin(){
    window.location.href = "login.html";
}

//退出登录
function UserExit(){
    if(!UserIsLogin()){
        alert("用户未登录")
        return 
    }
    if(confirm("请问是否退出登录")){
        var nowData = {};
        localStorage.setItem("userData", JSON.stringify(nowData));
        window.location.reload();
    }
}

//注销用户
function UserDelete(){
    if(!UserIsLogin()){
        alert("用户未登录")
        return 
    }
    // 弹出可输入的对话框
    const userInput = window.prompt('请输入你的密码:', '');

    // 检查用户是否输入了内容
    if (userInput !== null) {
        // 用户输入了内容，你可以在这里处理用户的输入
        console.log('你输入的密码是:', userInput);
    } else {
        // 用户取消了对话框
        console.log('用户取消了输入');
    }
    if(confirm("注销用户会导致您的用户所有信息全都删除,请问您真的要注销您的账户吗?")){
        POST_Req("/user/delete",DeleteUserParam(userData.token,userInput))
        .then(data => {
            if(data.status_code != 1100){
                alert(data.status_msg);
                return
            }
            alert(data.status_msg);
            nowData = {};
            localStorage.setItem("userData", JSON.stringify(nowData));
            window.location.reload();
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
}

//判断用户是否登录
function UserIsLogin(){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if (JSON.stringify(userData) !== "{}")return true;
    return false;
}

//上传视频
function UpLoadVideo(){
    if(!UserIsLogin()){
        alert("用户未登录")
        return 
    }
    window.location.href = "upload.html";
}

//视频嵌入之后的操作
function VideoLoadOperate(){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //获取视频信息
    var videoInfo = videoInfos[index];
    //获取视频发布人与当前用户关系信息
    POST_Req("/user/info",UserInfoParam(userData.token,videoInfo.userID))
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        
        //更改发布人1
        document.getElementById("publicUser1").innerHTML = data.name;
        //更改发布人2
        document.getElementById("publicUser2").innerHTML = data.name;
        //是否关注用户
        if (data.isCare){
            document.getElementById("careUser").innerHTML = `√`;
        }else{
            document.getElementById("careUser").innerHTML = `+`;
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
    //更改视频名称
    document.getElementById("vName").innerHTML = videoInfo.videoName;
    //更改视频标签
    document.getElementById("vTags").innerHTML = videoInfo.videoTags;
    //是否点赞视频
    if (videoInfo.isFavorite){
        document.getElementById("favorite").innerHTML = `取消点赞`;
        isfav = 1;
    }else{
        document.getElementById("favorite").innerHTML = `点赞`;
        isfav = 0;
    }
    //点赞数量
    document.getElementById("favoriteNum").innerHTML = videoInfo.videoFavoriteNum;
    //评论数量
    document.getElementById("commentNum").innerHTML = videoInfo.videoCommitNum;
    //更新视频播放器
    document.getElementById("video").innerHTML = `
        <source id="s" src="${videoInfo.videoLink}" type="video/mp4">
    `;
    document.getElementById("video").load();

    //重置操作
    comNum = 0;
    comText = [];
}

//视频划走之后的操作
function VideoCloseOperate(vID){
    if(videoInfos[index].isFavorite){
        isfav = isfav - 1;
    }
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    var test = [{
        token:userData.token,
        videoID:vID,
        time:videoWatchTime,
        isf:isfav,
        cnum:comNum,
        cTexts:comTexts
    }]
    POST_Req("/video/operate",OperateVideoParam(userData.token,vID,videoWatchTime,isfav,comNum,comTexts))
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

//各种点击事件
//首页点击事件
document.getElementById("home").addEventListener("click", function() {
    // 在此处添加您想要执行的操作
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';

    alert("首页内容暂未实现,正在为您跳转热点内容")
    document.getElementById("topVideo").click()
});

//热点点击事件
document.getElementById("topVideo").addEventListener("click", function() {
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取top视频信息数组
    GET_Req("/video/top","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        //嵌入视频
        VideoLoadOperate();
        listIndex = 1;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

//关注点击事件
document.getElementById("careVideo").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取top视频信息数组
    GET_Req("/video/care","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        //嵌入视频
        VideoLoadOperate();
        listIndex = 2;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

//推荐点击事件
document.getElementById("refereeVideo").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取top视频信息数组
    GET_Req("/video/referee","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        //嵌入视频
        VideoLoadOperate();
        listIndex = 3;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

//我的点击事件
document.getElementById("userCenter").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    window.location.href = "userCenter.html";
});

//发布人点击事件
document.getElementById("publicUser2").addEventListener("click",function(){

});

//关注发布人点击事件
document.getElementById("careUser").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }
    if(null != videoInfos[index]){
        if(this.innerText == "+"){
            //获取登录用户信息
            var userData = JSON.parse(localStorage.getItem("userData"));
            POST_Req("",CareUserParam(userData.token,videoInfos[index].userID,1))
            .then(data => {
                if(data.status_code != 0){
                    alert(data.status_msg);
                    return
                }
                alert("成功关注");
                this.innerHTML = `√`;
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }
        else if(this.innerText == "√"){
            alert("该用户已被关注")
        }
    }else{
        alert("发布人信息获取失败")
    }
});

//点赞点击事件
document.getElementById("favorite").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }
    isfav ^= 1;
    if(isfav){
        document.getElementById("favorite").innerHTML = `取消点赞`;
    }else{
        document.getElementById("favorite").innerHTML = `点赞`;
    }
});

//评论点击事件
document.getElementById("comment").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录或登录信息已过期")
        return
    }

});

//搜索点击事件
document.getElementById("search").addEventListener("click",function(){
    var searchText = document.getElementById("searchText").value
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取查找的视频信息数组
    POST_Req("video/search",SearchVideoParam(userData.token,searchText))
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        //嵌入视频
        VideoLoadOperate();
        listIndex = 4;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});




//各种监听事件
//监听键盘
window.onkeydown = function (e)
{
    if (e.keyCode == 38) {
        
    }
    if(e.keyCode == 40){
       
    }
}

//监听鼠标滚轮
document.querySelector('body').addEventListener('wheel', function(event) {
    // 阻止默认的滚动行为
    // event.preventDefault();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if(listIndex != 1){
        if(!UserIsLogin()){
            return
        }
    }
    // 获取滚轮滚动的距离
    const deltaY = event.deltaY;

    // 输出滚动方向
    if (deltaY > 0) {
        if (null != videoInfos[index]){
            VideoCloseOperate(videoInfos[index].videoID)
        }
        //切换下个视频
        index += 1;
        if (index < 0 || index >= videoInfos.length){
            if(index >= videoInfos.length){
                //再次调用获取视频添加到videoInfos后面
                if (listIndex < 4 || listIndex > 0){
                    GET_Req("/video/"+listValue[listIndex] , "token" , userData.token)
                    .then(data => {
                        if(data.status_code != 0){
                            alert(data.status_msg);
                            return
                        }
                        console.log(data);
                        videoInfos = videoInfos.concat(data.videoInfos)
                    })
                    .catch(error => {
                        console.error('Error:', error);
                    });
                }
            }
        }else{
            VideoLoadOperate();
        }
    } else if (deltaY < 0) {
        if (null != videoInfos[index]){
            VideoCloseOperate(videoInfos[index].videoID)
        }
        //切换上个视频
        index -= 1
        if (index < 0 || index >= videoInfos.length){
            if(index < 0){
                alert("前面已经没有视频了,正在为您刷新页面");
                //刷新videoInfos
                initVideo();
                GET_Req("/video/"+listValue[listIndex] , "token" , userData.token)
                .then(data => {
                    if(data.status_code != 0){
                        alert(data.status_msg);
                        return
                    }
                    console.log(data);
                    videoInfos = videoInfos.concat(data.videoInfos)
                    VideoLoadOperate();
                })
                .catch(error => {
                    console.error('Error:', error);
                });
            }
        }else{
            VideoLoadOperate();
        }
    }
});

//监听视频播放时长
document.getElementById("video").addEventListener("timeupdate", function() {
    // 获取当前视频播放时间（秒）
    videoWatchTime = parseInt(video.currentTime * 1000);
});
