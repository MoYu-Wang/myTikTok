
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

    showMessage("首页内容暂未实现,\n正在为您跳转热点内容")
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
document.getElementById("careVideo").addEventListener("click",function(){
    if(!UserIsLogin()){
        showMessage("用户未登录\n或登录信息已过期")
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
            showMessage(data.status_msg);
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
        showMessage("用户未登录\n或登录信息已过期")
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
            showMessage(data.status_msg);
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
        showMessage("用户未登录\n或登录信息已过期")
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
            showMessage(data.status_msg);
            return
        }
        console.log(data);
        videoInfos = videoInfos.concat(data.videoInfos)
        if(null == videoInfos[0]){
            showMessage("您好像还没有发布过视频,\n快来发布您的第一条视频吧!")
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
        showMessage("用户未登录\n或登录信息已过期")
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
    })
    .catch(error => {
        console.error('Error:', error);
    });

});

//历史记录
document.getElementById("myHistory").addEventListener("click",function(){
    if(!UserIsLogin()){
        showMessage("用户未登录\n或登录信息已过期")
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
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    if(null != videoInfos[index]){
        if(this.innerText == "+"){
            //获取登录用户信息
            var userData = JSON.parse(localStorage.getItem("userData"));
            POST_Req("/user/care",CareUserParam(userData.token,videoInfos[index].userID,1))
            .then(data => {
                if(data.status_code != 0){
                    showMessage(data.status_msg);
                    return
                }
                showMessage("成功关注");
                this.innerHTML = `√`;
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }
        else if(this.innerText == "√"){
            showMessage("您已经关注了该用户")
        }
    }else{
        showMessage("发布人信息获取失败")
    }
});

//点赞点击事件
document.getElementById("favorite").addEventListener("click",function(){
    if(!UserIsLogin()){
        showMessage("用户未登录\n或登录信息已过期")
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
            showMessage(data.status_msg);
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

//点击确认修改密码
document.getElementById("submitUpdate").addEventListener("click",function(){
    var pwd = document.getElementById("").value;
    var newpwd = document.getElementById("").value;
    var newpwd2 = document.getElementById("").value;
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if (newpwd != newpwd2){
        showMessage("两次输入的新密码不同")
        return;
    }
    POST_Req("/user/update/password",UpdatePasswordParam(userData.token,pwd,newpwd))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        alert("修改密码成功,请重新登录")
        nowData = {};
        localStorage.setItem("userData", JSON.stringify(nowData));
        window.location.href = "login.html";
    })
    .catch(error => {
        console.error('Error:', error);
    });

});


// 检查点击事件并隐藏浮窗
document.addEventListener('click', function(event) {
    // 检查点击的元素是否是浮窗或其子元素
    var isClickInside = floatWindow.contains(event.target);

    if (!isClickInside) {
        // 如果点击的不是浮窗区域，则隐藏浮窗
        floatWindow.style.display = 'none';
    }
});