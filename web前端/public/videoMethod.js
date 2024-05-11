showMessage("请根据鼠标滚轮上下滑动切换视频!")
var index = 0;
var videoInfos = [];
var videoOperateInfo = {};
var userVideoIndex = 0;
var userVideoInfos = [];
var videoORUserVideo = false;//false:正常视频表,true:用户中心视频表

let current = 0;
var userData = {};
var videoWatchTime = 0;//视频播放时间(ms)
var favIsClick = 0;//点赞操作缓存
var listIndex = 1;
var listValue = ["search","top","care","referee","works","favorite","history",""];
var videoComments = [];//视频评论





function loadImg(){
    
    document.getElementById("title").style.backgroundImage = "url('./Icon/抖音.png')";
    document.getElementById("comment").style.backgroundImage = "url('./Icon/评论.png')";
    document.getElementById("search").style.backgroundImage = "url('./Icon/搜索.png')";
    document.getElementById("careUser").style.backgroundImage = "url('./Icon/未关注.png')";
    document.getElementById("favorite").style.backgroundImage = "url('./Icon/未点赞.png')";
    document.getElementById("rebackBarVideo").style.backgroundImage = "url('./Icon/返回.png')";
    document.getElementById("downloadPC").style.backgroundImage = "url('./Icon/下载客户端.png')";
    document.getElementById("user_avatar").style.backgroundImage = "url('./Icon/用户.png')";
    document.getElementById("profile-picture").style.backgroundImage = "url('./Icon/发布人.png')";
    document.getElementById("publicUser2").style.backgroundImage = "url('./Icon/发布人.png')";
}

//初始加载事件
//频道显示
//视频切换源
window.onload = async function(){
    loadImg();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //判断是否登录
    if (await UserIsLogin()) {
        //获取基本信息
        GET_Req("/user/base","token" , userData.token)
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            userData.userID = data.userID
            userData.userName = data.userName
            userData.iphoneID = data.iphoneID
            document.getElementById("user_name").innerText = data.userName;
            document.getElementById("user_id").innerText = "AHUT_TikTok  ID号:\n" + data.userID;
            localStorage.setItem("userData", JSON.stringify(userData));
        })
        .catch(error => {
            console.error('Error:', error);
        });
        //
        
        document.getElementById("userlogin").value = "退出登录";
        document.getElementById("userlogin").setAttribute("onclick", "UserExit()");
        
        document.getElementById("user_id").style.display = "block";
        document.getElementById("updatePassword").style.display = "block";
        //登录情况默认进入推荐页面
        document.getElementById("refereeVideo").click();
    } else {
        //游客登录的情况
        document.getElementById("user_name").innerText = "未登录"
        document.getElementById("user_id").style.display = "none";
        document.getElementById("updatePassword").style.display = "none";
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

//切换主体 0:视频主体 1:个人中心主体
function checkBody(checkEnum){
    switch(checkEnum){
        case 0:
            // 切换视频主体
            document.getElementById("content").style.display = 'block';
            document.getElementById("user-profile").style.display = 'none';
            document.getElementById("search-body").style.display = 'none';
            //添加滚轮监听切换视频
            addScrollEventListener();
            break;
        case 1:
            // 切换个人中心主体
            //暂停播放
            document.getElementById("video").pause();
            document.getElementById("content").style.display = 'none';
            document.getElementById("user-profile").style.display = 'flex';
            document.getElementById("search-body").style.display = 'none';
            //移除滚轮监听
            removeScrollEventListener();
            break;
        case 2:
            // 切换搜索主体
            //暂停播放
            document.getElementById("video").pause();
            document.getElementById("content").style.display = 'none';
            document.getElementById("user-profile").style.display = 'none';
            document.getElementById("search-body").style.display = 'flex';
            //移除滚轮监听
            removeScrollEventListener();
        default:
            break;
    }
}

async function UserCenter(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    // 获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    ToUserCenter(userData.userID);
}

//去到uID用户中心
async function ToUserCenter(uID){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    //重置频道栏按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    buttons.forEach(function(button){
        button.style.backgroundColor = '#666';
    });
    //设置个人中心主体
    checkBody(1);
    // 获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //获取个人信息
    POST_Req("/user/info",UserInfoParam(userData.token,uID))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        document.getElementById("profile-userName").innerText = data.name;
        document.getElementById("profile-userID").innerText = "AHUT_TikTok  ID号:"+ data.id;
        document.getElementById("profile-getLikes").innerText = "获赞 " + data.getLikes;
        document.getElementById("profile-careCount").innerText = "关注 " + data.careCount;
        document.getElementById("profile-fansCount").innerText = "粉丝 " + data.fansCount;
        document.getElementById("user-works").onclick = function(){
            //更新用户中心视频内容
            UpdateUserCenterVideo(0,uID);
        }
        if(data.id === userData.userID){
            //本人主页
            //点击关注显示关注列表
            document.getElementById("profile-careCount").onclick = function(){

            }
            //点击粉丝显示粉丝列表
            document.getElementById("profile-fansCount").onclick = function(){

            }

            document.getElementById("user-careOperate").style.display = 'none';
            document.getElementById("user-baseInfo-edit").style.display = 'block';
            document.getElementById("user-baseInfo-edit").onclick = function(){

            }
            document.getElementById("user-likes").style.display = 'block';
            document.getElementById("user-likes").onclick = function(){
                //更新用户中心视频内容
                UpdateUserCenterVideo(1,uID);
            }
            document.getElementById("user-history").style.display = 'block';
            document.getElementById("user-history").onclick = function(){
                //更新用户中心视频内容
                UpdateUserCenterVideo(2,uID);
            }
        }else{
            document.getElementById("user-careOperate").style.display = 'flex';
            document.getElementById("user-careOperate").onclick = function(){
                var f;
                if(data.isCare){
                    f = -1
                }else{
                    f = 1
                }
                POST_Req("/user/care",CareUserParam(userData.token,uID,f))
                .then(response => {
                    if(response.status_code != 0){
                        showMessage(response.status_msg);
                        return
                    }
                    if(data.isCare){
                        showMessage("取消关注")
                        document.getElementById("user-careOperate").innerText = "关注";
                        data.isCare = false;
                    }else{
                        showMessage("关注成功")
                        document.getElementById("user-careOperate").innerText = "取消关注";
                        data.isCare = true;
                    }
                })
                .catch(error => {
                    console.error('Error:', error);
                });
            }
            document.getElementById("user-baseInfo-edit").style.display = 'none';
            document.getElementById("user-likes").style.display = 'none';
            document.getElementById("user-history").style.display = 'none';
            if(data.isCare){
                //已关注
                document.getElementById("user-careOperate").innerText = "取消关注";
            }else{
                //未关注
                document.getElementById("user-careOperate").innerText = "关注";
            }
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
    //更新用户中心视频内容
    UpdateUserCenterVideo(0,uID);
}

async function UpdateUserCenterVideo(videoEnum,uID){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    // 获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    var UrlStr;
    switch(videoEnum){
        case 0:
            UrlStr = "/user/works";
            document.getElementById("user-works").style.color = 'red';
            document.getElementById("user-likes").style.color = 'white';
            document.getElementById("user-history").style.color = 'white';
            break;
        case 1:
            UrlStr = "/user/favorite";
            document.getElementById("user-works").style.color = 'white';
            document.getElementById("user-likes").style.color = 'red';
            document.getElementById("user-history").style.color = 'white';
            break;
        case 2:
            UrlStr = "/user/history";
            document.getElementById("user-works").style.color = 'white';
            document.getElementById("user-likes").style.color = 'white';
            document.getElementById("user-history").style.color = 'red';
            break;
    }
    
    //初始化videoinfos数组和index
    initUserVideo()
    //查询videoInfos
    if(videoEnum == 0){
        POST_Req(UrlStr,UserWorksParam(userData.token,uID))
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            userVideoInfos = userVideoInfos.concat(data.videoInfos)
            if(null == userVideoInfos[0]){
                //没有作品
                const videoContainer = document.querySelector('.user-videoInfo');
                videoContainer.innerHTML = ''; // 清空现有视频
                return;
            }else{
                //显示作品
                const videoContainer = document.querySelector('.user-videoInfo');
                videoContainer.innerHTML = ''; // 清空现有视频
                userVideoInfos.forEach(async videoInfo => {
                    if (videoInfo == null){
                        return;
                    }
                    // 创建视频元素
                    const li = document.createElement('li');
                    li.className = 'user-videoInfo-data';
                    li.onclick = function(){
                        videoORUserVideo = true;
                        userVideoIndex = userVideoInfos.indexOf(videoInfo);
                        checkBody(0);
                        VideoLoadOperate();
                    }
                    videoContainer.appendChild(li);
                    //获取单个视频信息
                    POST_Req("/video/info",VideoOperateInfoParam(userData.token,videoInfo.videoID))
                    .then(async data => {
                        if(data.status_code != 0){
                            showMessage(data.status_msg);
                            return
                        }
                        // <div>${videoInfo.videoLink}</div>
                        li.innerHTML = `
                            <div>${videoInfo.videoName}</div>
                            <div>${videoInfo.videoTags}</div>
                            <div>${data.videoFavoriteNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/点赞.png);width:10px;heigth:10px;"></div>
                            <div>${data.videoCommitNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/评论.png);width:10px;heigth:10px;"></div>
                            <div class="deleteVideoDIV"></div>
                        `;
                        // 检查用户是否登录且为视频发布人
                        if (await UserIsLogin()) {
                            if (userData && userData.userID === videoInfo.userID) {
                                const deleteButton = document.createElement('button');
                                deleteButton.className = 'deleteVideo';
                                deleteButton.textContent = '删除视频';
                                deleteButton.onclick = function() {
                                    // 添加删除视频的功能，例如通过API请求
                                    POST_Req("/video/delete", DeleteVideoParam(userData.token,videoInfo.videoID))
                                    .then(async response => {
                                        if(response.status_code != 0){
                                            showMessage(response.status_msg);
                                            return;
                                        }
                                        li.remove(); // 或重新加载视频
                                        showMessage("删除成功");
                                    })
                                    .catch(error => {
                                        console.error('Delete error:', error);
                                    });
                                };
                                li.querySelector('.deleteVideoDIV').appendChild(deleteButton);
                            }
                        }
                    })
                    .catch(error => {
                        console.error('Error:', error);
                    });
                   
                });
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }else{
        GET_Req(UrlStr,"token",userData.token)
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            userVideoInfos = userVideoInfos.concat(data.videoInfos)
            if(null == userVideoInfos[0]){
                //没有作品
                const videoContainer = document.querySelector('.user-videoInfo');
                videoContainer.innerHTML = ''; // 清空现有视频
                return;
            }else{
                //显示作品
                const videoContainer = document.querySelector('.user-videoInfo');
                videoContainer.innerHTML = ''; // 清空现有视频
                userVideoInfos.forEach(async videoInfo => {
                    // 创建视频元素
                    const li = document.createElement('li');
                    li.className = 'user-videoInfo-data';
                    li.onclick = function(){
                        videoORUserVideo = true;
                        userVideoIndex = userVideoInfos.indexOf(videoInfo);
                        checkBody(0);
                        VideoLoadOperate();
                    }
                    videoContainer.appendChild(li);
                    //获取单个视频信息
                    POST_Req("/video/info",VideoOperateInfoParam(userData.token,videoInfo.videoID))
                    .then(async data => {
                        if(data.status_code != 0){
                            showMessage(data.status_msg);
                            return
                        }
                        // <div>${videoInfo.videoLink}</div>
                        li.innerHTML = `
                            <div>${videoInfo.videoName}</div>
                            <div>${videoInfo.videoTags}</div>
                            <div>${data.videoFavoriteNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/点赞.png);width:10px;heigth:10px;"></div>
                            <div>${data.videoCommitNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/评论.png);width:10px;heigth:10px;"></div>
                            <div class="deleteVideoDIV"></div>
                        `;
                        // 检查用户是否登录且为视频发布人
                        if (await UserIsLogin()) {
                            if (userData && userData.userID === videoInfo.userID) {
                                const deleteButton = document.createElement('button');
                                deleteButton.className = 'deleteVideo';
                                deleteButton.textContent = '删除视频';
                                deleteButton.onclick = function() {
                                    // 添加删除视频的功能，例如通过API请求
                                    POST_Req("/video/delete", DeleteVideoParam(userData.token,videoInfo.videoID))
                                    .then(async response => {
                                        if(response.status_code != 0){
                                            showMessage(response.status_msg);
                                            return;
                                        }
                                        li.remove(); // 或重新加载视频
                                        showMessage("删除成功");
                                    })
                                    .catch(error => {
                                        console.error('Delete error:', error);
                                    });
                                };
                                li.querySelector('.deleteVideoDIV').appendChild(deleteButton);
                            }
                        }
                    })
                    .catch(error => {
                        console.error('Error:', error);
                    });
                });
            }

        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
}

async function UpdateSearchVideo(searchText){
    var tk;
    if(UserIsLogin()){
        // 获取登录用户信息
        var userData = JSON.parse(localStorage.getItem("userData"));
        tk = userData.token
    }else{
        tk = 0
    }
    //发送get请求获取查找的视频信息数组
    GET_Req("/video/search","searchText",searchText)
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        //初始化视频
        initUserVideo();
        userVideoInfos = userVideoInfos.concat(data.videoInfos)
        if(userVideoInfos[0] == null){
            showMessage("没有找到与之相关的视频")
            document.getElementById("rebackBarVideo").click();
            return;
        }
        //显示作品
        const videoContainer = document.querySelector('.search-body');
        videoContainer.innerHTML = ''; // 清空现有视频
        userVideoInfos.forEach(async videoInfo => {
            // 创建视频元素
            const li = document.createElement('li');
            li.className = 'search-videoInfo-data';
            li.onclick = function(){
                videoORUserVideo = true;
                userVideoIndex = userVideoInfos.indexOf(videoInfo);
                checkBody(0);
                VideoLoadOperate();
            }
            videoContainer.appendChild(li);
            //获取单个视频信息
            POST_Req("/video/info",VideoOperateInfoParam(tk,videoInfo.videoID))
            .then(async data => {
                if(data.status_code != 0){
                    showMessage(data.status_msg);
                    return
                }
                // <div>${videoInfo.videoLink}</div>
                li.innerHTML = `
                            <div>${videoInfo.videoName}</div>
                            <div>${videoInfo.videoTags}</div>
                            <div>${data.videoFavoriteNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/点赞.png);width:10px;heigth:10px;"></div>
                            <div>${data.videoCommitNum}</div>
                            <div class="Img" style="float:'none'; background-image: url(./Icon/评论.png);width:10px;heigth:10px;"></div>
                            <div class="deleteVideoDIV"></div>
                        `;
            })
            .catch(error => {
                console.error('Error:', error);
            });
        });

    })
    .catch(error => {
        console.error('Error:', error);
    });
}

//视频初始化
function initVideo(){
    index = 0;
    videoInfos = [];
    videoOperateInfo = {};
    videoComments = [];
    videoORUserVideo = false;
}

//用户中心视频初始化
function initUserVideo(){
    userVideoIndex = 0;
    userVideoInfos = [];
    videoORUserVideo = true;
}


//用户登录页面跳转
function UserLogin(){
    window.location.href = "login.html";
}

//修改密码
async function UpdatePassword(){
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
    <h1 id="floatWindow-title">修改密码界面</h1>

    <div class="Img" style="float:'left'; background-image: url(./Icon/用户信息.png)"></div>
    <span id="float-userName" style="height: 50px;font-size: 45px">${userData.userName}</span></br>
    <span id="float-userID" style="height: 25px;font-size: 20px">ID:${userData.userID}</span></br></br>

    <span class="Img" style="background-image: url(./Icon/密码.png);width:20px;height:20px"></span>
    <label for="password">原密码:</label>
    <input type="password" id="password" minlength="5" maxlength="18"  title="请输入原密码" required><br>

    <span class="Img" style="background-image: url(./Icon/密码.png);width:20px;height:20px"></span>
    <label for="newPassword">新密码:</label>
    <input type="password" id="newPassword" minlength="5" maxlength="18" title="请输入新密码" required><br>

    <span class="Img" style="background-image: url(./Icon/密码.png);width:20px;height:20px"></span>
    <label for="newPassword">再次输入新密码:</label>
    <input type="password" id="newPassword2" minlength="5" maxlength="18"  title="密码必须与上一次输入密码相同" required><br>

    <button id="submitUpdatePwd">确认修改</button>
    `
    // 点击关闭按钮关闭模态浮窗
    document.getElementById("closeFloatWindows").onclick = function() {
        document.getElementById("floatWindow").style.display = "none";
        addScrollEventListener();
    }


    //点击确认修改密码
    document.getElementById("submitUpdatePwd").onclick = function(){
        var pwd = document.getElementById("password").value;
        var newpwd = document.getElementById("newPassword").value;
        var newpwd2 = document.getElementById("newPassword2").value;
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

    };

    
}

//退出登录
async function UserExit(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    if(confirm("请问是否退出登录")){
        var nowData = {};
        localStorage.setItem("userData", JSON.stringify(nowData));
        window.location.reload();
    }
}

//注销用户
async function UserDelete(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    
    if(confirm("注销用户会导致您的用户所有信息全部删除,\n请问您真的要注销您的账户吗?")){
        //获取登录用户信息
        var userData = JSON.parse(localStorage.getItem("userData"));
        // 弹出可输入的对话框
        const userInput = window.prompt('请输入你的密码:', '');

        // 检查用户是否输入了内容
        if (userInput === null) {
            return
        }
        POST_Req("/user/delete",DeleteUserParam(userData.token,userInput))
        .then(data => {
            if(data.status_code != 1100){
                showMessage(data.status_msg);
                return
            }
            showMessage(data.status_msg);
            nowData = {};
            localStorage.setItem("userData", JSON.stringify(nowData));
            window.location.reload();
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
}

async function UserIsLogin() {
    // 获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    
    // 判断是否登录
    if (JSON.stringify(userData) !== "{}") {
        try {
            // 判断用户登录信息是否过期
            const data = await GET_Req("/user/updatetoken", "token", userData.token);
            if (data.status_code != 0) {
                showMessage(data.status_msg);
                return false;
            }
            userData.token = data.token;
            return true;
        } catch (error) {
            console.error('Error:', error);
            return false;
        }
    }
    return false; // 添加这一行以处理未登录情况
}

//上传视频
async function UpLoadVideo(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    window.location.href = "upload.html";
}

//更改评论区
async function UpdateComment(){
    //初始化评论区信息
    videoComments = [];
    //获取视频信息
    var videoInfo
    if(!videoORUserVideo){
        videoInfo = videoInfos[index];
    }else{
        videoInfo = userVideoInfos[userVideoIndex];
    }
    //获取评论
    GET_Req("/video/getcomment","videoID",videoInfo.videoID)
    .then(data => {
        if (data.status_code != 0) {
            showMessage(data.status_msg);
            return;
        }
        videoComments = data.videoComments; // 假设返回的数据中包含一个评论数组
        const commentsContainer = document.querySelector('.comments');
        commentsContainer.innerHTML = ''; // 清空现有评论
        if(!videoComments)return;
        videoComments.forEach(async comment => {
            // 创建评论元素
            const li = document.createElement('li');
            var str = ''
            coms = Math.trunc((Date.now() - comment.commentTime/1e6)/1000)
            if(coms < 60){
                str = Math.trunc(coms) + '秒前';
            }else if(coms < 60*60){
                str = Math.trunc(coms/60) + '分钟前';
            }else if(coms < 60*60*24){
                str = Math.trunc(coms/(60*60)) + '小时前';
            }else{
                str = Math.trunc(coms/(60*60*24)) + '天前'; 
            }
            li.className = 'comment';
            li.innerHTML = `
                <span class="user">${comment.userName}:</span>
                <p class="comment-text">${comment.commentText}</p>
                <div class="comment-footer">
                    <span class="timestamp">${str}</span>
                </div>
            `;
            // 检查用户是否登录且为评论者
            if (await UserIsLogin()) {
                var userData = JSON.parse(localStorage.getItem("userData"));
                if (userData && userData.userID === comment.userID) {
                    const deleteButton = document.createElement('button');
                    deleteButton.className = 'comment-delete';
                    deleteButton.textContent = '删除评论';
                    deleteButton.onclick = function() {
                        // 添加删除评论的功能，例如通过API请求
                        POST_Req("/video/deletecomment", DeleteCommentParam(userData.token,videoInfo.videoID,comment.commentID))
                        .then(response => {
                            if(response.status_code != 0){
                                showMessage(response.status_msg);
                                return;
                            }
                            li.remove(); // 或重新加载评论
                            showMessage("删除成功");
                        })
                        .catch(error => {
                            console.error('Delete error:', error);
                        });
                    };
                    li.querySelector('.comment-footer').appendChild(deleteButton);
                }
            }
            commentsContainer.appendChild(li);
        });

    })
    .catch(error => {
        console.error('Error:', error);
    });
}

//更新频道栏
function updateBar(){
    //设置按钮颜色
    var buttons = document.querySelectorAll('.sidebar button');
    var cnt = 0;
    buttons.forEach(function(button){
        cnt++;
        if(cnt == listIndex){
            button.style.backgroundColor = '#5a8dd9';
        }else{
            button.style.backgroundColor = '#666';
        }
    });
}

//视频嵌入之后的操作
async function VideoLoadOperate(){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if(await UserIsLogin()){
        
    }
    //获取视频信息
    var videoInfo
    if(!videoORUserVideo){
        videoInfo = videoInfos[index];
    }else{
        videoInfo = userVideoInfos[userVideoIndex];
    }
    //获取视频发布人与当前用户关系信息
    POST_Req("/user/info",UserInfoParam(userData.token,videoInfo.userID))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        //更改发布人1
        document.getElementById("publicUser1").innerHTML = data.name;
        //是否关注用户
        if (data.isCare){
            document.getElementById("careUser").style.backgroundImage = "url('./Icon/已关注.png')";
            document.getElementById("careUser").setAttribute("isCare",true);
        }else{
            document.getElementById("careUser").style.backgroundImage = "url('./Icon/未关注.png')";
            document.getElementById("careUser").setAttribute("isCare",false);
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
    POST_Req("/video/info",VideoOperateInfoParam(userData.token,videoInfo.videoID))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
        videoOperateInfo = data;
        //是否点赞视频
        if (data.isFavorite){
            document.getElementById("favorite").style.backgroundImage = "url('./Icon/点赞.png')";
            document.getElementById("favorite").setAttribute("isFavorite",true);
        }else{
            document.getElementById("favorite").style.backgroundImage = "url('./Icon/未点赞.png')";
            document.getElementById("favorite").setAttribute("isFavorite",false);
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

    //更新评论区
    UpdateComment();

    //重置操作
    favIsClick = 0;
}


//视频划走之后的操作
function VideoCloseOperate(vID){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    var isfav = parseInt(document.getElementById("favorite").getAttribute("operate"));
    POST_Req("/video/operate",OperateVideoParam(userData.token,vID,videoWatchTime,isfav))
    .then(data => {
        if(data.status_code != 0){
            showMessage(data.status_msg);
            return
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

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
async function scrollEventHandler(event) {
    // 你的滚轮事件处理逻辑
    // 阻止默认的滚动行为
    // event.preventDefault();
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    if(listIndex != 1 && !videoORUserVideo){
        if(await UserIsLogin()==false){
            return
        }
    }
    // 获取滚轮滚动的距离
    const deltaY = event.deltaY;

    // 输出滚动方向
    if (deltaY > 0) {
        //如果以及登录以及链入视频
        var videoInfo;
        var videoIndex;
        var videolen;
        if(!videoORUserVideo){
            videoInfo = videoInfos[index];
            videoIndex = index;
            videolen = videoInfos.length;
        }else{
            videoInfo = userVideoInfos[userVideoIndex];
            videoIndex = userVideoIndex;
            videolen = userVideoInfos.length;
        }

        if (null != videoInfo && null != userData.token){
            VideoCloseOperate(videoInfo.videoID)
        }
        //判断下个index合不合法
        if (videoIndex + 1 < 0 || videoIndex + 1 >= videolen){
            //不合法
            if(!videoORUserVideo){
                //如果为频道视频
                if(videoIndex + 1 >= videolen){
                    //再次调用获取视频添加到videoInfos后面
                    if (listIndex < 4 && listIndex > 0){
                        GET_Req("/video/"+listValue[listIndex] , "token" , userData.token)
                        .then(data => {
                            if(data.status_code != 0){
                                showMessage(data.status_msg);
                                return
                            }
                            videoInfos = videoInfos.concat(data.videoInfos)
                        })
                        .catch(error => {
                            console.error('Error:', error);
                        });
                    }else{
                        showMessage("已经是最后一个视频了")
                        return 
                    }
                }
            }else{
                //如果不是频道视频
                showMessage("已经是最后一个视频了")
                return
            }
        }else{
            //合法
            //切换下个视频
            if(!videoORUserVideo){
                index += 1;
            }else{
                userVideoIndex += 1;
            }
            VideoLoadOperate();
        }
    } else if (deltaY < 0) {
        var videoInfo;
        var videoIndex;
        var videolen;
        if(!videoORUserVideo){
            videoInfo = videoInfos[index];
            videoIndex = index;
            videolen = videoInfos.length;
        }else{
            videoInfo = userVideoInfos[userVideoIndex];
            videoIndex = userVideoIndex;
            videolen = userVideoInfos.length;
        }

        if (null != videoInfo && null != userData.token){
            VideoCloseOperate(videoInfo.videoID)
        }
        if (videoIndex - 1 < 0 || videoIndex - 1 >= videolen){
            //如果不合法
            if(!videoORUserVideo){
                //如果为频道视频
                if(videoIndex - 1 < 0){
                    if (listIndex < 4 && listIndex > 0){
                        showMessage("前面已经没有视频了,\n正在为您刷新页面");
                        //刷新videoInfos
                        initVideo();
                        GET_Req("/video/"+listValue[listIndex] , "token" , userData.token)
                        .then(data => {
                            if(data.status_code != 0){
                                showMessage(data.status_msg);
                                return
                            }
                            videoInfos = videoInfos.concat(data.videoInfos)
                            VideoLoadOperate();
                        })
                        .catch(error => {
                            console.error('Error:', error);
                        });
                    }else{
                        showMessage("已经是第一个视频了!")
                        return
                    }
                }
            }else{
                //如果不是频道视频
                showMessage("已经是第一个视频了!")
                return
            }
        }else{
            //合法
            if(!videoORUserVideo){
                index -= 1
            }else{
                userVideoIndex -= 1
            }
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
