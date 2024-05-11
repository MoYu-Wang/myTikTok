
//各种点击事件
//热点点击事件
document.getElementById("topVideo").addEventListener("click", function() {
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if (JSON.stringify(userData) !== "{}"){
        //发送get请求获取top视频信息数组
        GET_Req("/video/top","token",userData.token)
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            console.log(data);
            videoInfos = videoInfos.concat(data.videoInfos)
            //嵌入视频
            VideoLoadOperate();
            listIndex = 1;
            //更新频道栏
            updateBar();
        })
        .catch(error => {
            console.error('Error:', error);
        });

    }else{
        //发送get请求获取top视频信息数组
        GET_Req("/video/top")
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
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
document.getElementById("careVideo").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取top视频信息数组
    GET_Req("/video/care","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(videoInfos[0] == null){
            showMessage("您未关注过用户,自动为您跳转至热点频道")
            document.getElementById("topVideo").click();
            return;
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 2;
        //更新频道栏
        updateBar();
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

//推荐点击事件
document.getElementById("refereeVideo").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
   
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送get请求获取top视频信息数组
    GET_Req("/video/referee","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        //嵌入视频
        VideoLoadOperate();
        listIndex = 3;
        //更新频道栏
        updateBar();
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

//我的作品点击事件
document.getElementById("myWorks").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //获取我的作品
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取works视频信息数组
    POST_Req("/user/works",UserWorksParam(userData.token,userData.userID))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
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
        //更新频道栏
        updateBar();
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//我的喜爱
document.getElementById("myFavorite").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
   
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //获取我的喜爱
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取favorite视频信息数组
    GET_Req("/user/favorite","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            showMessage("您没有点赞过的视频!\n正在为您跳转热点频道观看视频")
            document.getElementById("topVideo").click();
            return
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 5;
        //更新频道栏
        updateBar();
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//历史记录
document.getElementById("myHistory").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    
    videoORUserVideo = false;
    //切换视频主体
    checkBody(0);
    //获取历史记录
    //初始化videoinfos数组和index
    initVideo();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //发送post请求获取works视频信息数组
    GET_Req("/user/history","token",userData.token)
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            showMessage("您还未观看过视频,正在为您跳转热点频道")
            document.getElementById("topVideo").click();
            return
        }
        //嵌入视频
        VideoLoadOperate();
        listIndex = 6;
        //更新频道栏
        updateBar();
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//发布人点击事件
document.getElementById("publicUser2").addEventListener("click",function(){
    //转到用户中心
    ToUserCenter(videoInfos[index].userID)
});

//关注发布人点击事件
document.getElementById("careUser").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    if(null != videoInfos[index]){
        if(document.getElementById("careUser").getAttribute("isCare") == 'false'){
            //获取登录用户信息
            var userData = JSON.parse(localStorage.getItem("userData"));
            POST_Req("/user/care",CareUserParam(userData.token,videoInfos[index].userID,1))
            .then(data => {
                if(data.status_code != 0){
                    showMessage(data.status_msg);
                    return
                }
                showMessage("成功关注");
                document.getElementById("careUser").setAttribute("isCare",true);
                document.getElementById("careUser").style.backgroundImage = "url('./Icon/已关注.png')";
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }
        else{
            showMessage("您已经关注了该用户")
        }
    }else{
        showMessage("发布人信息获取失败")
    }
});

//点赞点击事件
document.getElementById("favorite").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    if(document.getElementById("favorite").getAttribute("isFavorite") == 'false'){
        // 点赞操作
        document.getElementById("favorite").style.backgroundImage = "url('./Icon/点赞.png')";
        document.getElementById("favoriteNum").innerText = String(parseInt(document.getElementById("favoriteNum").innerText) + 1);
        document.getElementById("favorite").setAttribute("operate",1);
    }else{
        // 取消点赞操作
        document.getElementById("favorite").style.backgroundImage = "url('./Icon/未点赞.png')";
        document.getElementById("favoriteNum").innerText = String(parseInt(document.getElementById("favoriteNum").innerText) - 1);
        document.getElementById("favorite").setAttribute("operate",-1);
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
        //设置评论图标
        document.getElementById("comment").style.backgroundImage = "url('./Icon/评论.png')";
        //添加鼠标滚轮监听切换视频
        addScrollEventListener();
    } else {
        // 如果当前是显示状态，则隐藏频道栏
        sidebar.style.display = 'none';
        //设置评论图标
        document.getElementById("comment").style.backgroundImage = "url('./Icon/点击评论.png')";
        //移除鼠标滚轮监听切换视频
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
document.getElementById("search").addEventListener("click",async function(){
    var searchText = document.getElementById("searchText").value
    if (searchText === ""){
        showMessage("请输入内容");
        return ;
    }
    //重置频道栏按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    checkBody(2);
    UpdateSearchVideo(searchText)
});

//点击返回
document.getElementById("rebackBarVideo").addEventListener("click",async function(){
    //关闭评论区
    document.getElementById("video-comments").style.display = "none";
    // 如果当前是隐藏状态，则显示频道栏
    document.getElementById('sidebar').style.display = 'block';
    //设置评论图标
    document.getElementById("comment").style.backgroundImage = "url('./Icon/评论.png')";

    videoORUserVideo = false;
    checkBody(0)
    VideoLoadOperate();
    updateBar();
});

//点击编辑资料
document.getElementById("user-baseInfo-edit").addEventListener("click",async function(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //设置浮窗可见
    document.getElementById("floatWindow").style.display = "block";
    //移除滚轮监听
    removeScrollEventListener();
    //设置浮窗内容
    document.getElementById("floatWindow-data").innerHTML = `
    <span class="close" id="closeFloatWindows">&times;</span>
    
    
    <div class="Img" id="title" style=" background-image: url(./Icon/抖音.png);"></div>
    <h1 id="floatWindow-title">编辑信息界面</h1>

    <span class="Img" style="background-image: url(./Icon/用户信息.png);width:20px;height:20px"></span>
    <span>用户昵称:</span></br>
    <input type="edit-userName" id="edit-userName" minlength="5" maxlength="18" value=${userData.userName}><br>

    
    <span class="Img" style="background-image: url(./Icon/手机号.png);width:20px;height:20px"></span>
    <span>用户手机号:</span></br>
    <input type="edit-userIphone" id="edit-userIphone" minlength="11" maxlength="11" value=${userData.iphoneID}><br>

    <button id="edit-userInfo">确认修改</button>
    `;

    // 点击关闭按钮关闭模态浮窗
    document.getElementById("closeFloatWindows").onclick = function() {
        document.getElementById("floatWindow").style.display = "none";
        addScrollEventListener();
    }


    //点击确认修改信息
    document.getElementById("edit-userInfo").onclick = function(){
        var userName = document.getElementById("edit-userName").value;
        var userIphone = document.getElementById("edit-userIphone").value;
        if(userIphone.length != 11 && userIphone.length != 0){
            showMessage("手机号格式不正确")
            return;
        }
        POST_Req("/user/update/info",UpdateUserInfoParam(userData.token,userName,userIphone))
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            if(userName != ""){
                userData.userName = userName
            }
            if(userIphone != ""){
                userData.iphoneID = userIphone
            }
            localStorage.setItem("userData", JSON.stringify(userData));
            alert("修改信息成功")
            UserCenter();
            document.getElementById("closeFloatWindows").click();
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
});

// 点击关闭按钮或模态外部区域关闭模态浮窗
document.getElementById("closeFloatWindows").onclick = function() {
    document.getElementById("floatWindow").style.display = "none";
    addScrollEventListener();
}

window.onclick = function(event) {
  if (event.target == document.getElementById("floatWindow")) {
    document.getElementById("floatWindow").style.display = "none";
    addScrollEventListener();
  }
}
