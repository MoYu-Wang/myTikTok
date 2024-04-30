showMessage("请根据鼠标滚轮上下滑动切换视频!")
var index = 0;
var videoInfos = [];
var videoOperateInfo = {};
let current = 0;
var userData = {};
var videoWatchTime = 0;//视频播放时间(ms)
var favIsClick = 0;//点赞操作缓存
var listIndex = 1;
var listValue = ["search","top","care","referee","works","favorite","history",""];
var videoComments = [];//视频评论


//初始加载事件
//频道显示
//视频切换源
window.onload = async function(){
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    console.log(userData)
    //判断是否登录
    if (await UserIsLogin()) {
        //获取基本信息
        GET_Req("/user/base","token" , userData.token)
        .then(data => {
            if(data.status_code != 0){
                showMessage(data.status_msg);
                return
            }
            console.log(data);
            userData.userID = data.userID
            userData.userName = data.userName
            userData.iphoneID = data.iphoneID
            document.getElementById("user_name").innerText = data.userName;
            document.getElementById("user_id").innerText = "抖音(低仿版)ID号:\n" + data.userID;
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


//个人中心
async function UserInfo(){
    if(await UserIsLogin()==false){
        showMessage("用户未登录")
        return 
    }
    showMessage("暂未实现");
}

//视频初始化
function initVideo(){
    index = 0;
    videoInfos = [];
    videoOperateInfo = {};
    videoComments = [];
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
    //设置浮窗可见
    document.getElementById("floatWindow").style.display = "block";
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    document.getElementById("float-userName").innerText = userData.userName
    document.getElementById("float-userID").innerText = "ID:" + userData.userID
    
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
        if (userInput !== null) {
            // 用户输入了内容，你可以在这里处理用户的输入
            console.log('你输入的密码是:', userInput);
        } else {
            // 用户取消了对话框
            console.log('用户取消了输入');
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
    console.log(userData);
    
    // 判断是否登录
    if (JSON.stringify(userData) !== "{}") {
        try {
            // 判断用户登录信息是否过期
            const data = await GET_Req("/user/updatetoken", "token", userData.token);
            if (data.status_code != 0) {
                showMessage(data.status_msg);
                return false;
            }
            console.log(data);
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
    var videoInfo = videoInfos[index];
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
            li.className = 'comment';
            li.innerHTML = `
                <span class="user">${comment.userName}:</span>
                <p class="comment-text">${comment.commentText}</p>
                <div class="comment-footer">
                    <span class="timestamp">${Math.trunc((Date.now() - comment.commentTime/1e6)/(1000*60))}分钟前</span>
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
            showMessage(data.status_msg);
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
            showMessage(data.status_msg);
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

    //更新评论区
    UpdateComment();

    //重置操作
    favIsClick = 0;
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
    if(listIndex != 1){
        if(await UserIsLogin()==false){
            return
        }
    }
    // 获取滚轮滚动的距离
    const deltaY = event.deltaY;

    // 输出滚动方向
    if (deltaY > 0) {
        //如果以及登录以及链入视频
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
                            showMessage(data.status_msg);
                            return
                        }
                        console.log(data);
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
                    showMessage("前面已经没有视频了,\n正在为您刷新页面");
                    //刷新videoInfos
                    initVideo();
                    GET_Req("/video/"+listValue[listIndex] , "token" , userData.token)
                    .then(data => {
                        if(data.status_code != 0){
                            showMessage(data.status_msg);
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
                    showMessage("已经是第一个视频了!")
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
