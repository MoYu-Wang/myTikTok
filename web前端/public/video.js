alert("请根据鼠标滚轮上下滑动切换视频!")
var index = 0;
var videoInfos = [];
var videoOperateInfo = {};
let current = 0;
var userData = {};
var videoWatchTime = 0;//视频播放时间(ms)
var favIsClick = 0;//点赞操作缓存
var comNum = 0;//评论数量缓存
var comTexts = [];//评论文本缓存
var listIndex = 1;
var listValue = ["search","top","care","referee","works","favorite","history",""];

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
        //登录情况默认进入推荐页面
        document.getElementById("refereeVideo").click();
    } else {
        //游客登录的情况
        document.getElementById("user_name").innerText = "未登录"

        document.getElementById("userlogin").value = "登录账号";
        document.getElementById("userlogin").setAttribute("onclick", "UserLogin()");
        //未登录情况默认进入热点页面
        document.getElementById("topVideo").click();
    }
    

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
    videoOperateInfo = {};
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
    if(confirm("注销用户会导致您的用户所有信息全部删除,\n请问您真的要注销您的账户吗?")){
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
    POST_Req("/video/info",VideoOperateInfoParam(userData.token,videoInfo.videoID))
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoOperateInfo = data;
        //是否点赞视频
        if (data.isFavorite){
            document.getElementById("favorite").innerHTML = `取消点赞`;
        }else{
            document.getElementById("favorite").innerHTML = `点赞`;
        }
        //点赞数量
        document.getElementById("favoriteNum").innerHTML = data.videoFavoriteNum;
        //评论数量
        document.getElementById("commentNum").innerHTML = data.videoCommitNum;
        
    })
    .catch(error => {
        console.error('Error:', error);
    });
    
    
    //更改视频名称
    document.getElementById("vName").innerHTML = videoInfo.videoName;
    //更改视频标签
    document.getElementById("vTags").innerHTML = videoInfo.videoTags;
    //更新视频播放器
    document.getElementById("video").innerHTML = `
        <source id="s" src="${videoInfo.videoLink}" type="video/mp4">
    `;
    document.getElementById("video").load();

    //重置操作
    favIsClick = 0;
    comNum = 0;
    comText = [];
}

//视频划走之后的操作
function VideoCloseOperate(vID){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    var isfav = 0
    if(videoOperateInfo.isFavorite){
        isfav = -1 * favIsClick;
    }else{
        isfav = favIsClick;
    }
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

    alert("首页内容暂未实现,\n正在为您跳转热点内容")
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
    if (JSON.stringify(userData) !== "{}"){
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

    }else{
        //发送get请求获取top视频信息数组
        GET_Req("/video/top")
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

    }
    
});

//关注点击事件
document.getElementById("careVideo").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
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
        alert("用户未登录\n或登录信息已过期")
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

//我的作品点击事件
document.getElementById("myWorks").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //获取我的作品
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取works视频信息数组
    POST_Req("/user/works",UserWorksParam(userData.token,userData.userID))
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            alert("您好像还没有发布过视频,\n快来发布您的第一条视频吧!")
            UpLoadVideo()
            return
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 4;
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//我的喜爱
document.getElementById("myFavorite").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //获取我的喜爱
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取favorite视频信息数组
    GET_Req("/user/favorite","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            alert("您没有点赞过的视频!\n正在为您跳转热点频道观看视频")
            document.getElementById("topVideo").click();
            return
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 5;
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//历史记录
document.getElementById("myHistory").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
        return
    }
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    this.style.backgroundColor = '#5a8dd9';
    //获取历史记录
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取works视频信息数组
    GET_Req("/user/history","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            alert(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            alert("您还未观看过视频,正在为您跳转热点频道")
            document.getElementById("topVideo").click();
            return
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 6;
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//发布人点击事件
document.getElementById("publicUser2").addEventListener("click",function(){

});

//关注发布人点击事件
document.getElementById("careUser").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
        return
    }
    if(null != videoInfos[index]){
        if(this.innerText == "+"){
            //获取登录用户信息
            var userData = JSON.parse(localStorage.getItem("userData"));
            POST_Req("/user/care",CareUserParam(userData.token,videoInfos[index].userID,1))
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
            alert("您已经关注了该用户")
        }
    }else{
        alert("发布人信息获取失败")
    }
});

//点赞点击事件
document.getElementById("favorite").addEventListener("click",function(){
    if(!UserIsLogin()){
        alert("用户未登录\n或登录信息已过期")
        return
    }
    favIsClick ^= 1;
    if(favIsClick){
        document.getElementById("favorite").innerHTML = `取消点赞`;
        document.getElementById("favoriteNum").innerText = String(parseInt(document.getElementById("favoriteNum").innerText) + 1);
    }else{
        document.getElementById("favorite").innerHTML = `点赞`;
        document.getElementById("favoriteNum").innerText = String(parseInt(document.getElementById("favoriteNum").innerText) - 1);
    }
});

//评论点击事件
document.getElementById("comment").addEventListener("click",function(){
    //设置频道栏
    var sidebar = document.getElementById('sidebar');
    // 检查当前频道栏的显示状态
    if (sidebar.style.display === 'none') {
        // 如果当前是隐藏状态，则显示频道栏
        sidebar.style.display = 'block';
        addScrollEventListener();
    } else {
        // 如果当前是显示状态，则隐藏频道栏
        sidebar.style.display = 'none';
        removeScrollEventListener();
    }
    //设置评论区
    var commentSection = document.getElementById("video-comments");
    if (commentSection) { // 检查commentSection是否存在
        if (commentSection.style.display === "none") {
            commentSection.style.display = "block";
            // 可以在这里添加加载评论的逻辑，例如发送请求获取评论数据并填充到评论区域
            // loadComments();
        } else {
            commentSection.style.display = "none";
        }
    } else {
        console.error("commentSection is null or undefined");
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
        listIndex = 0;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});


// 添加鼠标滚轮事件监听器
function addScrollEventListener() {
    document.querySelector('body').addEventListener('wheel', scrollEventHandler);
}

// 移除鼠标滚轮事件监听器
function removeScrollEventListener() {
    document.querySelector('body').removeEventListener('wheel', scrollEventHandler);
}


//各种监听事件
//监听键盘
window.onkeydown = function (e)
{
    if (e.keyCode == 38) {
        
    }
    if(e.keyCode == 40){
       
    }
}
// 定义滚轮事件处理函数
function scrollEventHandler(event) {
    // 你的滚轮事件处理逻辑
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
        if (null != videoInfos[index] && null != userData.token){
            VideoCloseOperate(videoInfos[index].videoID)
        }
        //判断下个index合不合法
        if (index + 1 < 0 || index + 1 >= videoInfos.length){
            if(index + 1 >= videoInfos.length){
                //再次调用获取视频添加到videoInfos后面
                if (listIndex < 4 && listIndex > 0){
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
                }else{
                    alert("已经是最后一个视频了")
                    return 
                }
            }
        }else{
            //切换下个视频
            index += 1;
            VideoLoadOperate();
        }
    } else if (deltaY < 0) {
        if (null != videoInfos[index] && null != userData.token){
            VideoCloseOperate(videoInfos[index].videoID)
        }
        if (index - 1 < 0 || index - 1 >= videoInfos.length){
            if(index - 1 < 0){
                if (listIndex < 4 && listIndex > 0){
                    alert("前面已经没有视频了,\n正在为您刷新页面");
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
                }else{
                    alert("已经是第一个视频了!")
                    return
                }
            }
        }else{
            //切换上个视频
            index -= 1
            VideoLoadOperate();
        }
    }
}
//监听鼠标滚轮
document.querySelector('body').addEventListener('wheel', scrollEventHandler);

//监听视频播放时长
document.getElementById("video").addEventListener("timeupdate", function() {
    // 获取当前视频播放时间（秒）
    videoWatchTime = parseInt(video.currentTime * 1000);
});
