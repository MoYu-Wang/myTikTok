
//浮窗显示函数
function showMessage(message) {
    const messageBox = document.getElementById('messageBox');
    messageBox.textContent = message;
    messageBox.classList.remove('hidden');
    messageBox.classList.add('show');
    setTimeout(() => {
        messageBox.classList.remove('show');
        messageBox.classList.add('hidden');
    }, 3000); // 3秒后隐藏消息框
}
