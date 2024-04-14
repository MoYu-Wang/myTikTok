//点击发布评论
document.getElementById("publishComment").addEventListener("click",function(){
    if(!UserIsLogin()){
        showMessage("用户未登录\n或登录信息已过期")
        return
    }
    //获取登录用户信息
    var userData = JSON.parse(localStorage.getItem("userData"));
    //获取视频信息
    var videoInfo = videoInfos[index];
    //获取评论内容
    var textareaContent = document.getElementById('commentText').value;
    console.log(textareaContent); // 打印<textarea>中的内容
    POST_Req("/video/comment",CommentVideoParam(userData.token,videoInfo.videoID,textareaContent))
    .then(data => {
        if (data.status_code != 0) {
            showMessage(data.status_msg);
            return;
        }
        UpdateComment()
        showMessage("评论成功")
        document.getElementById('commentText').value = '';
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

