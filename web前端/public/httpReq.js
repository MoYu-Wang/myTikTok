//项目服务器IP和端口

var myURL = "http://150.158.115.140:11316/myTikTok";

/** 
 * 用法:
 * POST_Req(route, jsonData)
    .then(data => {
        // 在这里处理返回的数据
        console.log("Received data:", data);
    })
    .catch(error => {
        // 错误处理
        console.error("Error:", error);
    });
 * 
*/
function POST_Req(route,jsondata){
    return fetch(myURL+ route, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(jsondata),// 确保发送的是字符串化的 JSON 数据
    })
    .then(response => {
        if (!response.ok) {
            alert("服务器未响应")
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
}


//strdata  =  "/key?value"
function GET_Req(route,key,value){
    var url;
    if (null != key && null != value){
      // 构建完整的 URL
      url = `${myURL}${route}?${key}=${value}`;
    }else{
      url = `${myURL}${route}`;
    }
    return fetch(url)
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok ' + response.statusText);
      }
      return response.json(); // 解析JSON格式的响应体
    })
}

function RegisterParam(uName,pwd,iphID){
  return {
    username: uName,
    password: pwd,
    iphoneID: iphID
  }
}

function LoginParam(uID,pwd,iphID){
  return {
    userID: uID,
    password: pwd,
    iphoneID: iphID
  }
}

function UpdataUserParam(Token,uName,pwd,iphID){
  return {
    userName:uName,
    password:pwd,
    iphoneID:iphID,
    token:Token
  }
}

function ForgetPwdParam(uID,iphID){
  return {
    userID: uID,
    iphoneID: iphID
  }
}

function DeleteUserParam(Token,pwd){
  return {
    token:Token,
    password: pwd
  }
}

function UserInfoParam(Token,uID){
  return {
    userID:uID,
    token:Token
  }
}

function VideoOperateInfoParam(Token,vID){
  return {
    videoID:vID,
    token:Token
  }
}

function UpLoadVideoParam(Token,vName,vTags,vLink){
  return {
    token:Token,
    videoName:vName,
    videoTags:vTags,
    videoLink:vLink
  }
}

function FavoriteVideoParam(Token,vID,isF){
  return {
    token:Token,
    videoID:vID,
    isFavorite:isF
  }
}

function CommentVideoParam(Token,vID,cText){
  return {
    token:Token,
    videoID:vID,
    commentText:cText
  }
}

function DeleteCommentParam(Token,vID,cID){
  return {
    token:Token,
    videoID:vID,
    commentID:cID
  }
}

function OperateVideoParam(Token,vID,wTime,isF){
  return {
    token:Token,
    videoID:vID,
    watchTime:wTime,
    isFavorite:isF,
  }
}

function UserWorksParam(Token,uID){
  return {
    token:Token,
    userID:uID
  }
}

function CareUserParam(Token,uID,opt){
  return {
    token:Token,
    userID:uID,
    operate:opt
  }
}

function SearchVideoParam(Token,sText){
  return {
    token:Token,
    searchText:sText
  }
}

function UpdatePasswordParam(Token,pwd,newpwd){
  return {
    token:Token,
    password:pwd,
    newPassword:newpwd
  }
}